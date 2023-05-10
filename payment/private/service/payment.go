package service

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/go-saas/commerce/order/api/order/v1"
	pb "github.com/go-saas/commerce/payment/api/gateway/v1"
	sapi "github.com/go-saas/kit/pkg/api"
	"github.com/go-saas/kit/pkg/authn"
	"github.com/go-saas/kit/pkg/authz/authz"
	kithttp "github.com/go-saas/kit/pkg/server/http"
	"github.com/go-saas/kit/pkg/utils"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
	"github.com/stripe/stripe-go/v74/webhook"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
)

type PaymentService struct {
	trust            sapi.TrustedContextValidator
	auth             authz.Service
	orderInternalSrv v1.OrderInternalServiceServer
	stripeClient     *client.API
	l                *log.Helper
}

var _ pb.PaymentGatewayServiceServer = (*PaymentService)(nil)
var _ pb.StripePaymentGatewayServiceServer = (*PaymentService)(nil)

func NewPaymentService(
	trust sapi.TrustedContextValidator,
	auth authz.Service,
	orderInternalSrv v1.OrderInternalServiceServer,
	stripeClient *client.API,
	logger log.Logger,
) *PaymentService {
	return &PaymentService{
		trust:            trust,
		auth:             auth,
		orderInternalSrv: orderInternalSrv,
		stripeClient:     stripeClient,
		l:                log.NewHelper(logger),
	}
}

func (s *PaymentService) GetPaymentMethod(ctx context.Context, req *pb.GetPaymentMethodRequest) (*pb.GetPaymentMethodReply, error) {
	return &pb.GetPaymentMethodReply{}, nil
}

func (s *PaymentService) CreateStripePaymentIntent(ctx context.Context, req *pb.CreateStripePaymentIntentRequest) (*pb.CreateStripePaymentIntentReply, error) {
	userInfo, err := authn.ErrIfUnauthenticated(ctx)
	if err != nil {
		return nil, err
	}
	order, err := s.orderInternalSrv.GetInternalOrder(ctx, &v1.GetInternalOrderRequest{Id: req.OrderId})
	if err != nil {
		return nil, err
	}
	if order.CustomerId != userInfo.GetId() {
		return nil, errors.NotFound("", "")
	}
	params := &stripe.CustomerParams{}
	params.Metadata = map[string]string{
		"user_id": userInfo.GetId(),
	}

	customer, err := s.stripeClient.Customers.New(params)
	if err != nil {
		return nil, handleStripeError(err)
	}
	paymentIntentParams := &stripe.PaymentIntentParams{
		Amount:   &order.TotalPrice.Amount,
		Currency: &order.TotalPrice.CurrencyCode,
		Customer: &customer.ID,
	}
	params.Metadata = map[string]string{
		"user_id":  userInfo.GetId(),
		"order_id": req.OrderId,
	}
	intent, err := s.stripeClient.PaymentIntents.New(paymentIntentParams)
	if err != nil {
		return nil, handleStripeError(err)
	}
	return &pb.CreateStripePaymentIntentReply{
		PaymentIntent: intent.ClientSecret,
		CustomerId:    customer.ID,
	}, nil
}

func (s *PaymentService) StripeWebhook(ctx context.Context, req *emptypb.Empty) (*pb.StripeWebhookReply, error) {
	if req, ok := kithttp.ResolveHttpRequest(ctx); ok {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		event, err := webhook.ConstructEvent(body, req.Header.Get("stripe-signature"), s.stripeClient.Customers.Key)
		if err != nil {
			return nil, handleStripeError(err)
		}
		data := event.Data
		eventType := event.Type
		log.Infof("receive event type %s with data %v", eventType, data)
		switch eventType {
		case "payment_intent.succeeded":
			intent := stripe.PaymentIntent{}
			json.Unmarshal(data.Raw, &intent)
			_, err = s.orderInternalSrv.InternalOrderPaySuccess(ctx, &v1.InternalOrderPaySuccessRequest{
				Id:        intent.Metadata["order_id"],
				PayExtra:  utils.Map2Structpb(data.Object),
				PaidPrice: nil,
				PayWay:    "stripe",
			})
			if err != nil {
				return nil, err
			}
		case "payment_intent.payment_failed":
		case "setup_intent.setup_failed":
		case "setup_intent.succeeded":
		case "setup_intent.created":
		default:

		}
		return &pb.StripeWebhookReply{}, nil
	} else {
		return nil, errors.BadRequest("", "")
	}
}

func handleStripeError(err error) error {
	//TODO handle stripe
	return err
}
