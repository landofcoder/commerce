// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: order/api/order/v1/order_internal.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	price "github.com/go-saas/kit/pkg/price"
	_ "github.com/go-saas/kit/pkg/query"
	lbs "github.com/go-saas/lbs"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateInternalOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShippingAddr *lbs.Address               `protobuf:"bytes,16,opt,name=shipping_addr,json=shippingAddr,proto3" json:"shipping_addr,omitempty"`
	BillingAddr  *lbs.Address               `protobuf:"bytes,17,opt,name=billing_addr,json=billingAddr,proto3" json:"billing_addr,omitempty"`
	CustomerId   string                     `protobuf:"bytes,18,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	PayBefore    *durationpb.Duration       `protobuf:"bytes,19,opt,name=pay_before,json=payBefore,proto3,oneof" json:"pay_before,omitempty"`
	Extra        *structpb.Struct           `protobuf:"bytes,100,opt,name=extra,proto3" json:"extra,omitempty"`
	Items        []*CreateInternalOrderItem `protobuf:"bytes,200,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CreateInternalOrderRequest) Reset() {
	*x = CreateInternalOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_order_v1_order_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInternalOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInternalOrderRequest) ProtoMessage() {}

func (x *CreateInternalOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_order_v1_order_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInternalOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateInternalOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_api_order_v1_order_internal_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInternalOrderRequest) GetShippingAddr() *lbs.Address {
	if x != nil {
		return x.ShippingAddr
	}
	return nil
}

func (x *CreateInternalOrderRequest) GetBillingAddr() *lbs.Address {
	if x != nil {
		return x.BillingAddr
	}
	return nil
}

func (x *CreateInternalOrderRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateInternalOrderRequest) GetPayBefore() *durationpb.Duration {
	if x != nil {
		return x.PayBefore
	}
	return nil
}

func (x *CreateInternalOrderRequest) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *CreateInternalOrderRequest) GetItems() []*CreateInternalOrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateInternalOrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qty           int32          `protobuf:"varint,1,opt,name=qty,proto3" json:"qty,omitempty"`
	Price         *price.PricePb `protobuf:"bytes,30,opt,name=price,proto3" json:"price,omitempty"`
	OriginalPrice *price.PricePb `protobuf:"bytes,40,opt,name=original_price,json=originalPrice,proto3" json:"original_price,omitempty"`
	IsGiveaway    bool           `protobuf:"varint,50,opt,name=is_giveaway,json=isGiveaway,proto3" json:"is_giveaway,omitempty"`
	Product       *OrderProduct  `protobuf:"bytes,100,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *CreateInternalOrderItem) Reset() {
	*x = CreateInternalOrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_order_v1_order_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInternalOrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInternalOrderItem) ProtoMessage() {}

func (x *CreateInternalOrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_order_v1_order_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInternalOrderItem.ProtoReflect.Descriptor instead.
func (*CreateInternalOrderItem) Descriptor() ([]byte, []int) {
	return file_order_api_order_v1_order_internal_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInternalOrderItem) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

func (x *CreateInternalOrderItem) GetPrice() *price.PricePb {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *CreateInternalOrderItem) GetOriginalPrice() *price.PricePb {
	if x != nil {
		return x.OriginalPrice
	}
	return nil
}

func (x *CreateInternalOrderItem) GetIsGiveaway() bool {
	if x != nil {
		return x.IsGiveaway
	}
	return false
}

func (x *CreateInternalOrderItem) GetProduct() *OrderProduct {
	if x != nil {
		return x.Product
	}
	return nil
}

type GetInternalOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetInternalOrderRequest) Reset() {
	*x = GetInternalOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_order_v1_order_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInternalOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternalOrderRequest) ProtoMessage() {}

func (x *GetInternalOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_order_v1_order_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternalOrderRequest.ProtoReflect.Descriptor instead.
func (*GetInternalOrderRequest) Descriptor() ([]byte, []int) {
	return file_order_api_order_v1_order_internal_proto_rawDescGZIP(), []int{2}
}

func (x *GetInternalOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InternalOrderPaySuccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PayExtra  *structpb.Struct       `protobuf:"bytes,21,opt,name=pay_extra,json=payExtra,proto3" json:"pay_extra,omitempty"`
	PaidTime  *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=paid_time,json=paidTime,proto3" json:"paid_time,omitempty"`
	PaidPrice *price.PricePb         `protobuf:"bytes,30,opt,name=paid_price,json=paidPrice,proto3" json:"paid_price,omitempty"`
	PayWay    string                 `protobuf:"bytes,31,opt,name=pay_way,json=payWay,proto3" json:"pay_way,omitempty"`
}

func (x *InternalOrderPaySuccessRequest) Reset() {
	*x = InternalOrderPaySuccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_order_v1_order_internal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalOrderPaySuccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalOrderPaySuccessRequest) ProtoMessage() {}

func (x *InternalOrderPaySuccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_order_v1_order_internal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalOrderPaySuccessRequest.ProtoReflect.Descriptor instead.
func (*InternalOrderPaySuccessRequest) Descriptor() ([]byte, []int) {
	return file_order_api_order_v1_order_internal_proto_rawDescGZIP(), []int{3}
}

func (x *InternalOrderPaySuccessRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InternalOrderPaySuccessRequest) GetPayExtra() *structpb.Struct {
	if x != nil {
		return x.PayExtra
	}
	return nil
}

func (x *InternalOrderPaySuccessRequest) GetPaidTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PaidTime
	}
	return nil
}

func (x *InternalOrderPaySuccessRequest) GetPaidPrice() *price.PricePb {
	if x != nil {
		return x.PaidPrice
	}
	return nil
}

func (x *InternalOrderPaySuccessRequest) GetPayWay() string {
	if x != nil {
		return x.PayWay
	}
	return ""
}

type InternalOrderPaySuccessReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InternalOrderPaySuccessReply) Reset() {
	*x = InternalOrderPaySuccessReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_api_order_v1_order_internal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalOrderPaySuccessReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalOrderPaySuccessReply) ProtoMessage() {}

func (x *InternalOrderPaySuccessReply) ProtoReflect() protoreflect.Message {
	mi := &file_order_api_order_v1_order_internal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalOrderPaySuccessReply.ProtoReflect.Descriptor instead.
func (*InternalOrderPaySuccessReply) Descriptor() ([]byte, []int) {
	return file_order_api_order_v1_order_internal_proto_rawDescGZIP(), []int{4}
}

var File_order_api_order_v1_order_internal_proto protoreflect.FileDescriptor

var file_order_api_order_v1_order_internal_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x6c, 0x62, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x62, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x41, 0x64, 0x64, 0x72, 0x12, 0x2f, 0x0a, 0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x62,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x70, 0x61, 0x79, 0x42, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x4f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0xc8,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74,
	0x65, 0x6d, 0x42, 0x0b, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x79, 0x5f, 0x62,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0xf2, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x71, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x50, 0x62, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50,
	0x62, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x67, 0x69, 0x76, 0x65, 0x61, 0x77, 0x61, 0x79, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x47, 0x69, 0x76, 0x65, 0x61, 0x77, 0x61,
	0x79, 0x12, 0x47, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x0b, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf3, 0x01, 0x0a, 0x1e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x08, 0x70, 0x61, 0x79, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x61,
	0x69, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70, 0x61, 0x69, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x62, 0x52, 0x09, 0x70, 0x61, 0x69, 0x64, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x5f, 0x77, 0x61, 0x79, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x79, 0x57, 0x61, 0x79, 0x22, 0x1e, 0x0a, 0x1c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xdc, 0x02, 0x0a, 0x14,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x81, 0x01, 0x0a, 0x17, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x32, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x61, 0x61, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_api_order_v1_order_internal_proto_rawDescOnce sync.Once
	file_order_api_order_v1_order_internal_proto_rawDescData = file_order_api_order_v1_order_internal_proto_rawDesc
)

func file_order_api_order_v1_order_internal_proto_rawDescGZIP() []byte {
	file_order_api_order_v1_order_internal_proto_rawDescOnce.Do(func() {
		file_order_api_order_v1_order_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_api_order_v1_order_internal_proto_rawDescData)
	})
	return file_order_api_order_v1_order_internal_proto_rawDescData
}

var file_order_api_order_v1_order_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_order_api_order_v1_order_internal_proto_goTypes = []interface{}{
	(*CreateInternalOrderRequest)(nil),     // 0: order.api.order.v1.CreateInternalOrderRequest
	(*CreateInternalOrderItem)(nil),        // 1: order.api.order.v1.CreateInternalOrderItem
	(*GetInternalOrderRequest)(nil),        // 2: order.api.order.v1.GetInternalOrderRequest
	(*InternalOrderPaySuccessRequest)(nil), // 3: order.api.order.v1.InternalOrderPaySuccessRequest
	(*InternalOrderPaySuccessReply)(nil),   // 4: order.api.order.v1.InternalOrderPaySuccessReply
	(*lbs.Address)(nil),                    // 5: lbs.Address
	(*durationpb.Duration)(nil),            // 6: google.protobuf.Duration
	(*structpb.Struct)(nil),                // 7: google.protobuf.Struct
	(*price.PricePb)(nil),                  // 8: price.PricePb
	(*OrderProduct)(nil),                   // 9: order.api.order.v1.OrderProduct
	(*timestamppb.Timestamp)(nil),          // 10: google.protobuf.Timestamp
	(*Order)(nil),                          // 11: order.api.order.v1.Order
}
var file_order_api_order_v1_order_internal_proto_depIdxs = []int32{
	5,  // 0: order.api.order.v1.CreateInternalOrderRequest.shipping_addr:type_name -> lbs.Address
	5,  // 1: order.api.order.v1.CreateInternalOrderRequest.billing_addr:type_name -> lbs.Address
	6,  // 2: order.api.order.v1.CreateInternalOrderRequest.pay_before:type_name -> google.protobuf.Duration
	7,  // 3: order.api.order.v1.CreateInternalOrderRequest.extra:type_name -> google.protobuf.Struct
	1,  // 4: order.api.order.v1.CreateInternalOrderRequest.items:type_name -> order.api.order.v1.CreateInternalOrderItem
	8,  // 5: order.api.order.v1.CreateInternalOrderItem.price:type_name -> price.PricePb
	8,  // 6: order.api.order.v1.CreateInternalOrderItem.original_price:type_name -> price.PricePb
	9,  // 7: order.api.order.v1.CreateInternalOrderItem.product:type_name -> order.api.order.v1.OrderProduct
	7,  // 8: order.api.order.v1.InternalOrderPaySuccessRequest.pay_extra:type_name -> google.protobuf.Struct
	10, // 9: order.api.order.v1.InternalOrderPaySuccessRequest.paid_time:type_name -> google.protobuf.Timestamp
	8,  // 10: order.api.order.v1.InternalOrderPaySuccessRequest.paid_price:type_name -> price.PricePb
	0,  // 11: order.api.order.v1.OrderInternalService.CreateInternalOrder:input_type -> order.api.order.v1.CreateInternalOrderRequest
	2,  // 12: order.api.order.v1.OrderInternalService.GetInternalOrder:input_type -> order.api.order.v1.GetInternalOrderRequest
	3,  // 13: order.api.order.v1.OrderInternalService.InternalOrderPaySuccess:input_type -> order.api.order.v1.InternalOrderPaySuccessRequest
	11, // 14: order.api.order.v1.OrderInternalService.CreateInternalOrder:output_type -> order.api.order.v1.Order
	11, // 15: order.api.order.v1.OrderInternalService.GetInternalOrder:output_type -> order.api.order.v1.Order
	4,  // 16: order.api.order.v1.OrderInternalService.InternalOrderPaySuccess:output_type -> order.api.order.v1.InternalOrderPaySuccessReply
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_order_api_order_v1_order_internal_proto_init() }
func file_order_api_order_v1_order_internal_proto_init() {
	if File_order_api_order_v1_order_internal_proto != nil {
		return
	}
	file_order_api_order_v1_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_order_api_order_v1_order_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInternalOrderRequest); i {
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
		file_order_api_order_v1_order_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInternalOrderItem); i {
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
		file_order_api_order_v1_order_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInternalOrderRequest); i {
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
		file_order_api_order_v1_order_internal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalOrderPaySuccessRequest); i {
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
		file_order_api_order_v1_order_internal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalOrderPaySuccessReply); i {
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
	file_order_api_order_v1_order_internal_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_api_order_v1_order_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_api_order_v1_order_internal_proto_goTypes,
		DependencyIndexes: file_order_api_order_v1_order_internal_proto_depIdxs,
		MessageInfos:      file_order_api_order_v1_order_internal_proto_msgTypes,
	}.Build()
	File_order_api_order_v1_order_internal_proto = out.File
	file_order_api_order_v1_order_internal_proto_rawDesc = nil
	file_order_api_order_v1_order_internal_proto_goTypes = nil
	file_order_api_order_v1_order_internal_proto_depIdxs = nil
}
