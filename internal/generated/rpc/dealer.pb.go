// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: api/rpc/dealer.proto

package rpc

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workflow string               `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"` // Тип сделки заказа
	Items    []*ItemCreateRequest `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`       // Список позиций в заказе
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_dealer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_dealer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_rpc_dealer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetWorkflow() string {
	if x != nil {
		return x.Workflow
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*ItemCreateRequest {
	if x != nil {
		return x.Items
	}
	return nil
}

type ItemCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                             // Уникальный идентификатор объявления
	Name     string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                          // Наименование товара
	SellerId int64                  `protobuf:"varint,3,opt,name=seller_id,json=sellerId,proto3" json:"seller_id,omitempty"` // Идентификатор продавца
	Price    int64                  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`                       // Стоимость товара
	Discount *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=discount,proto3,oneof" json:"discount,omitempty"`            // Скидка на товар
}

func (x *ItemCreateRequest) Reset() {
	*x = ItemCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_dealer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCreateRequest) ProtoMessage() {}

func (x *ItemCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_dealer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCreateRequest.ProtoReflect.Descriptor instead.
func (*ItemCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_rpc_dealer_proto_rawDescGZIP(), []int{1}
}

func (x *ItemCreateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemCreateRequest) GetSellerId() int64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *ItemCreateRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ItemCreateRequest) GetDiscount() *wrapperspb.Int64Value {
	if x != nil {
		return x.Discount
	}
	return nil
}

type CreateOrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`           // Уникальный идентификатор покупки
	Created int64 `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"` // Метка времени создания
}

func (x *CreateOrderReply) Reset() {
	*x = CreateOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_dealer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderReply) ProtoMessage() {}

func (x *CreateOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_dealer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderReply.ProtoReflect.Descriptor instead.
func (*CreateOrderReply) Descriptor() ([]byte, []int) {
	return file_api_rpc_dealer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateOrderReply) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

type ApplyTransitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                // Уникальный идентификатор заказа по которому выполняется переход
	Transition string `protobuf:"bytes,2,opt,name=transition,proto3" json:"transition,omitempty"` // Название перехода для текущего сценария
}

func (x *ApplyTransitionRequest) Reset() {
	*x = ApplyTransitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_dealer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyTransitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyTransitionRequest) ProtoMessage() {}

func (x *ApplyTransitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_dealer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyTransitionRequest.ProtoReflect.Descriptor instead.
func (*ApplyTransitionRequest) Descriptor() ([]byte, []int) {
	return file_api_rpc_dealer_proto_rawDescGZIP(), []int{3}
}

func (x *ApplyTransitionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApplyTransitionRequest) GetTransition() string {
	if x != nil {
		return x.Transition
	}
	return ""
}

type ApplyTransitionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Флаг успешность выполнения перехода
}

func (x *ApplyTransitionReply) Reset() {
	*x = ApplyTransitionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_dealer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyTransitionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyTransitionReply) ProtoMessage() {}

func (x *ApplyTransitionReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_dealer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyTransitionReply.ProtoReflect.Descriptor instead.
func (*ApplyTransitionReply) Descriptor() ([]byte, []int) {
	return file_api_rpc_dealer_proto_rawDescGZIP(), []int{4}
}

func (x *ApplyTransitionReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_rpc_dealer_proto protoreflect.FileDescriptor

var file_api_rpc_dealer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x2c, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x11,
	0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x48, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x14, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x96, 0x01,
	0x0a, 0x06, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0f, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rpc_dealer_proto_rawDescOnce sync.Once
	file_api_rpc_dealer_proto_rawDescData = file_api_rpc_dealer_proto_rawDesc
)

func file_api_rpc_dealer_proto_rawDescGZIP() []byte {
	file_api_rpc_dealer_proto_rawDescOnce.Do(func() {
		file_api_rpc_dealer_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rpc_dealer_proto_rawDescData)
	})
	return file_api_rpc_dealer_proto_rawDescData
}

var file_api_rpc_dealer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_rpc_dealer_proto_goTypes = []interface{}{
	(*CreateOrderRequest)(nil),     // 0: api.CreateOrderRequest
	(*ItemCreateRequest)(nil),      // 1: api.ItemCreateRequest
	(*CreateOrderReply)(nil),       // 2: api.CreateOrderReply
	(*ApplyTransitionRequest)(nil), // 3: api.ApplyTransitionRequest
	(*ApplyTransitionReply)(nil),   // 4: api.ApplyTransitionReply
	(*wrapperspb.Int64Value)(nil),  // 5: google.protobuf.Int64Value
}
var file_api_rpc_dealer_proto_depIdxs = []int32{
	1, // 0: api.CreateOrderRequest.items:type_name -> api.ItemCreateRequest
	5, // 1: api.ItemCreateRequest.discount:type_name -> google.protobuf.Int64Value
	0, // 2: api.Dealer.CreateOrder:input_type -> api.CreateOrderRequest
	3, // 3: api.Dealer.ApplyTransition:input_type -> api.ApplyTransitionRequest
	2, // 4: api.Dealer.CreateOrder:output_type -> api.CreateOrderReply
	4, // 5: api.Dealer.ApplyTransition:output_type -> api.ApplyTransitionReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_rpc_dealer_proto_init() }
func file_api_rpc_dealer_proto_init() {
	if File_api_rpc_dealer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rpc_dealer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_api_rpc_dealer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCreateRequest); i {
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
		file_api_rpc_dealer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderReply); i {
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
		file_api_rpc_dealer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyTransitionRequest); i {
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
		file_api_rpc_dealer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyTransitionReply); i {
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
	file_api_rpc_dealer_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_rpc_dealer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_rpc_dealer_proto_goTypes,
		DependencyIndexes: file_api_rpc_dealer_proto_depIdxs,
		MessageInfos:      file_api_rpc_dealer_proto_msgTypes,
	}.Build()
	File_api_rpc_dealer_proto = out.File
	file_api_rpc_dealer_proto_rawDesc = nil
	file_api_rpc_dealer_proto_goTypes = nil
	file_api_rpc_dealer_proto_depIdxs = nil
}
