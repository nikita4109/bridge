// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/bridge_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BridgeStatusUpdate_Status int32

const (
	BridgeStatusUpdate_PENDING  BridgeStatusUpdate_Status = 0
	BridgeStatusUpdate_SENT     BridgeStatusUpdate_Status = 1
	BridgeStatusUpdate_FAILED   BridgeStatusUpdate_Status = 2
	BridgeStatusUpdate_RECEIVED BridgeStatusUpdate_Status = 3
)

// Enum value maps for BridgeStatusUpdate_Status.
var (
	BridgeStatusUpdate_Status_name = map[int32]string{
		0: "PENDING",
		1: "SENT",
		2: "FAILED",
		3: "RECEIVED",
	}
	BridgeStatusUpdate_Status_value = map[string]int32{
		"PENDING":  0,
		"SENT":     1,
		"FAILED":   2,
		"RECEIVED": 3,
	}
)

func (x BridgeStatusUpdate_Status) Enum() *BridgeStatusUpdate_Status {
	p := new(BridgeStatusUpdate_Status)
	*p = x
	return p
}

func (x BridgeStatusUpdate_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BridgeStatusUpdate_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bridge_service_proto_enumTypes[0].Descriptor()
}

func (BridgeStatusUpdate_Status) Type() protoreflect.EnumType {
	return &file_proto_bridge_service_proto_enumTypes[0]
}

func (x BridgeStatusUpdate_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BridgeStatusUpdate_Status.Descriptor instead.
func (BridgeStatusUpdate_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_bridge_service_proto_rawDescGZIP(), []int{3, 0}
}

type SendBridgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromChainId uint64 `protobuf:"varint,1,opt,name=from_chain_id,json=fromChainId,proto3" json:"from_chain_id,omitempty"`
	ToChainId   uint64 `protobuf:"varint,2,opt,name=to_chain_id,json=toChainId,proto3" json:"to_chain_id,omitempty"`
	TokenName   string `protobuf:"bytes,3,opt,name=token_name,json=tokenName,proto3" json:"token_name,omitempty"`
	Amount      string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"` // string representation of big.Int
}

func (x *SendBridgeRequest) Reset() {
	*x = SendBridgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBridgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBridgeRequest) ProtoMessage() {}

func (x *SendBridgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBridgeRequest.ProtoReflect.Descriptor instead.
func (*SendBridgeRequest) Descriptor() ([]byte, []int) {
	return file_proto_bridge_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendBridgeRequest) GetFromChainId() uint64 {
	if x != nil {
		return x.FromChainId
	}
	return 0
}

func (x *SendBridgeRequest) GetToChainId() uint64 {
	if x != nil {
		return x.ToChainId
	}
	return 0
}

func (x *SendBridgeRequest) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *SendBridgeRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type SendBridgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BridgeId string `protobuf:"bytes,1,opt,name=bridge_id,json=bridgeId,proto3" json:"bridge_id,omitempty"`
}

func (x *SendBridgeResponse) Reset() {
	*x = SendBridgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBridgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBridgeResponse) ProtoMessage() {}

func (x *SendBridgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBridgeResponse.ProtoReflect.Descriptor instead.
func (*SendBridgeResponse) Descriptor() ([]byte, []int) {
	return file_proto_bridge_service_proto_rawDescGZIP(), []int{1}
}

func (x *SendBridgeResponse) GetBridgeId() string {
	if x != nil {
		return x.BridgeId
	}
	return ""
}

type GetBridgeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BridgeId string `protobuf:"bytes,1,opt,name=bridge_id,json=bridgeId,proto3" json:"bridge_id,omitempty"`
}

func (x *GetBridgeStatusRequest) Reset() {
	*x = GetBridgeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBridgeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBridgeStatusRequest) ProtoMessage() {}

func (x *GetBridgeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBridgeStatusRequest.ProtoReflect.Descriptor instead.
func (*GetBridgeStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_bridge_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetBridgeStatusRequest) GetBridgeId() string {
	if x != nil {
		return x.BridgeId
	}
	return ""
}

type BridgeStatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          BridgeStatusUpdate_Status `protobuf:"varint,1,opt,name=status,proto3,enum=bridge.BridgeStatusUpdate_Status" json:"status,omitempty"`
	TransactionHash string                    `protobuf:"bytes,2,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	Error           string                    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BridgeStatusUpdate) Reset() {
	*x = BridgeStatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BridgeStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BridgeStatusUpdate) ProtoMessage() {}

func (x *BridgeStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BridgeStatusUpdate.ProtoReflect.Descriptor instead.
func (*BridgeStatusUpdate) Descriptor() ([]byte, []int) {
	return file_proto_bridge_service_proto_rawDescGZIP(), []int{3}
}

func (x *BridgeStatusUpdate) GetStatus() BridgeStatusUpdate_Status {
	if x != nil {
		return x.Status
	}
	return BridgeStatusUpdate_PENDING
}

func (x *BridgeStatusUpdate) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *BridgeStatusUpdate) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_bridge_service_proto protoreflect.FileDescriptor

var file_proto_bridge_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x6f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49, 0x64, 0x22,
	0xcb, 0x01, 0x0a, 0x12, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e,
	0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x39, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x44, 0x10, 0x03, 0x32, 0xa9, 0x01,
	0x0a, 0x0d, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x19, 0x2e,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x2e, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x62, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_bridge_service_proto_rawDescOnce sync.Once
	file_proto_bridge_service_proto_rawDescData = file_proto_bridge_service_proto_rawDesc
)

func file_proto_bridge_service_proto_rawDescGZIP() []byte {
	file_proto_bridge_service_proto_rawDescOnce.Do(func() {
		file_proto_bridge_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bridge_service_proto_rawDescData)
	})
	return file_proto_bridge_service_proto_rawDescData
}

var file_proto_bridge_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_bridge_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_bridge_service_proto_goTypes = []interface{}{
	(BridgeStatusUpdate_Status)(0), // 0: bridge.BridgeStatusUpdate.Status
	(*SendBridgeRequest)(nil),      // 1: bridge.SendBridgeRequest
	(*SendBridgeResponse)(nil),     // 2: bridge.SendBridgeResponse
	(*GetBridgeStatusRequest)(nil), // 3: bridge.GetBridgeStatusRequest
	(*BridgeStatusUpdate)(nil),     // 4: bridge.BridgeStatusUpdate
}
var file_proto_bridge_service_proto_depIdxs = []int32{
	0, // 0: bridge.BridgeStatusUpdate.status:type_name -> bridge.BridgeStatusUpdate.Status
	1, // 1: bridge.BridgeService.SendBridge:input_type -> bridge.SendBridgeRequest
	3, // 2: bridge.BridgeService.GetBridgeStatus:input_type -> bridge.GetBridgeStatusRequest
	2, // 3: bridge.BridgeService.SendBridge:output_type -> bridge.SendBridgeResponse
	4, // 4: bridge.BridgeService.GetBridgeStatus:output_type -> bridge.BridgeStatusUpdate
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_bridge_service_proto_init() }
func file_proto_bridge_service_proto_init() {
	if File_proto_bridge_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bridge_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBridgeRequest); i {
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
		file_proto_bridge_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBridgeResponse); i {
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
		file_proto_bridge_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBridgeStatusRequest); i {
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
		file_proto_bridge_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BridgeStatusUpdate); i {
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
			RawDescriptor: file_proto_bridge_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bridge_service_proto_goTypes,
		DependencyIndexes: file_proto_bridge_service_proto_depIdxs,
		EnumInfos:         file_proto_bridge_service_proto_enumTypes,
		MessageInfos:      file_proto_bridge_service_proto_msgTypes,
	}.Build()
	File_proto_bridge_service_proto = out.File
	file_proto_bridge_service_proto_rawDesc = nil
	file_proto_bridge_service_proto_goTypes = nil
	file_proto_bridge_service_proto_depIdxs = nil
}
