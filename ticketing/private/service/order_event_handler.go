package service

import (
	"context"
	v1 "github.com/go-saas/commerce/order/event/v1"
	"github.com/go-saas/commerce/ticketing/private/biz"
	"github.com/go-saas/kit/event"
)

func NewOrderSuccessNotification(showRepo biz.ShowRepo, ticketRepo biz.TicketRepo) event.ConsumerHandler {
	msg := &v1.OrderPaySuccessEvent{}
	return event.ProtoHandler[*v1.OrderPaySuccessEvent](msg, event.HandlerFuncOf[*v1.OrderPaySuccessEvent](func(ctx context.Context, msg *v1.OrderPaySuccessEvent) error {
		var tickets []*biz.Ticket
		for _, item := range msg.Order.Items {
			if item.Product.Type != biz.OrderTypeShow {
				continue
			}
			//find show
			show, err := showRepo.Get(ctx, item.Product.Id)
			if err != nil {
				return err
			}
			//generate ticket
			for i := 0; i < int(item.Qty); i++ {
				tic := biz.NewTicket(msg.Order.CustomerId)
				tic.LocationID = show.LocationID
				tic.HallID = show.HallID
				//TODO seat
				tic.ActivityID = show.ActivityID
				tic.ShowID = show.ID.String()
				tic.ShowSalesTypeID = item.Product.SkuId
				tic.Status = biz.TicketStatusValid
				tic.OrderID = msg.Order.Id
				tic.PaidTime = msg.Order.PaidTime.AsTime()
				tickets = append(tickets, tic)
			}
		}
		if len(tickets) == 0 {
			return nil
		}
		err := ticketRepo.BatchCreate(ctx, tickets, 10)
		if err != nil {
			return err
		}
		return nil
	}))
}
