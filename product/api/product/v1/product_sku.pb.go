// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: product/api/product/v1/product_sku.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/go-saas/kit/pkg/query"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateOrUpdateProductSku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	MainPic *Media   `protobuf:"bytes,13,opt,name=main_pic,json=mainPic,proto3" json:"main_pic,omitempty"`
	Medias  []*Media `protobuf:"bytes,14,rep,name=medias,proto3" json:"medias,omitempty"`
	Price   *Price   `protobuf:"bytes,15,opt,name=price,proto3" json:"price,omitempty"`
	Barcode string   `protobuf:"bytes,19,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Stock   []*Stock `protobuf:"bytes,62,rep,name=stock,proto3" json:"stock,omitempty"`
}

func (x *CreateOrUpdateProductSku) Reset() {
	*x = CreateOrUpdateProductSku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_product_sku_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateProductSku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateProductSku) ProtoMessage() {}

func (x *CreateOrUpdateProductSku) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_product_sku_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateProductSku.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateProductSku) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_product_sku_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrUpdateProductSku) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrUpdateProductSku) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateOrUpdateProductSku) GetMainPic() *Media {
	if x != nil {
		return x.MainPic
	}
	return nil
}

func (x *CreateOrUpdateProductSku) GetMedias() []*Media {
	if x != nil {
		return x.Medias
	}
	return nil
}

func (x *CreateOrUpdateProductSku) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *CreateOrUpdateProductSku) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *CreateOrUpdateProductSku) GetStock() []*Stock {
	if x != nil {
		return x.Stock
	}
	return nil
}

type ProductSku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	MainPic *Media   `protobuf:"bytes,13,opt,name=main_pic,json=mainPic,proto3" json:"main_pic,omitempty"`
	Medias  []*Media `protobuf:"bytes,14,rep,name=medias,proto3" json:"medias,omitempty"`
	Price   *Price   `protobuf:"bytes,15,opt,name=price,proto3" json:"price,omitempty"`
	Barcode string   `protobuf:"bytes,19,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Stock   []*Stock `protobuf:"bytes,62,rep,name=stock,proto3" json:"stock,omitempty"`
}

func (x *ProductSku) Reset() {
	*x = ProductSku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_api_product_v1_product_sku_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSku) ProtoMessage() {}

func (x *ProductSku) ProtoReflect() protoreflect.Message {
	mi := &file_product_api_product_v1_product_sku_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSku.ProtoReflect.Descriptor instead.
func (*ProductSku) Descriptor() ([]byte, []int) {
	return file_product_api_product_v1_product_sku_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSku) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductSku) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductSku) GetMainPic() *Media {
	if x != nil {
		return x.MainPic
	}
	return nil
}

func (x *ProductSku) GetMedias() []*Media {
	if x != nil {
		return x.Medias
	}
	return nil
}

func (x *ProductSku) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *ProductSku) GetBarcode() string {
	if x != nil {
		return x.Barcode
	}
	return ""
}

func (x *ProductSku) GetStock() []*Stock {
	if x != nil {
		return x.Stock
	}
	return nil
}

var File_product_api_product_v1_product_sku_proto protoreflect.FileDescriptor

var file_product_api_product_v1_product_sku_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x73, 0x6b, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x6b, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x70, 0x69, 0x63, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x07, 0x6d, 0x61, 0x69,
	0x6e, 0x50, 0x69, 0x63, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x3e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22,
	0xa7, 0x02, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x69, 0x63,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x07, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x69, 0x63, 0x12, 0x35,
	0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x06, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x72,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x3e, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_api_product_v1_product_sku_proto_rawDescOnce sync.Once
	file_product_api_product_v1_product_sku_proto_rawDescData = file_product_api_product_v1_product_sku_proto_rawDesc
)

func file_product_api_product_v1_product_sku_proto_rawDescGZIP() []byte {
	file_product_api_product_v1_product_sku_proto_rawDescOnce.Do(func() {
		file_product_api_product_v1_product_sku_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_api_product_v1_product_sku_proto_rawDescData)
	})
	return file_product_api_product_v1_product_sku_proto_rawDescData
}

var file_product_api_product_v1_product_sku_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_api_product_v1_product_sku_proto_goTypes = []interface{}{
	(*CreateOrUpdateProductSku)(nil), // 0: product.api.product.v1.CreateOrUpdateProductSku
	(*ProductSku)(nil),               // 1: product.api.product.v1.ProductSku
	(*Media)(nil),                    // 2: product.api.product.v1.Media
	(*Price)(nil),                    // 3: product.api.product.v1.Price
	(*Stock)(nil),                    // 4: product.api.product.v1.Stock
}
var file_product_api_product_v1_product_sku_proto_depIdxs = []int32{
	2, // 0: product.api.product.v1.CreateOrUpdateProductSku.main_pic:type_name -> product.api.product.v1.Media
	2, // 1: product.api.product.v1.CreateOrUpdateProductSku.medias:type_name -> product.api.product.v1.Media
	3, // 2: product.api.product.v1.CreateOrUpdateProductSku.price:type_name -> product.api.product.v1.Price
	4, // 3: product.api.product.v1.CreateOrUpdateProductSku.stock:type_name -> product.api.product.v1.Stock
	2, // 4: product.api.product.v1.ProductSku.main_pic:type_name -> product.api.product.v1.Media
	2, // 5: product.api.product.v1.ProductSku.medias:type_name -> product.api.product.v1.Media
	3, // 6: product.api.product.v1.ProductSku.price:type_name -> product.api.product.v1.Price
	4, // 7: product.api.product.v1.ProductSku.stock:type_name -> product.api.product.v1.Stock
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_product_api_product_v1_product_sku_proto_init() }
func file_product_api_product_v1_product_sku_proto_init() {
	if File_product_api_product_v1_product_sku_proto != nil {
		return
	}
	file_product_api_product_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_product_api_product_v1_product_sku_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateProductSku); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_api_product_v1_product_sku_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSku); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_api_product_v1_product_sku_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_api_product_v1_product_sku_proto_goTypes,
		DependencyIndexes: file_product_api_product_v1_product_sku_proto_depIdxs,
		MessageInfos:      file_product_api_product_v1_product_sku_proto_msgTypes,
	}.Build()
	File_product_api_product_v1_product_sku_proto = out.File
	file_product_api_product_v1_product_sku_proto_rawDesc = nil
	file_product_api_product_v1_product_sku_proto_goTypes = nil
	file_product_api_product_v1_product_sku_proto_depIdxs = nil
}
