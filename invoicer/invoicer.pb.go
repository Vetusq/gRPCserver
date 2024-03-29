// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: invoicer.proto

package invoicer

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

type RequestWalletTokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Tokenaddress string `protobuf:"bytes,2,opt,name=tokenaddress,proto3" json:"tokenaddress,omitempty"`
}

func (x *RequestWalletTokenInfo) Reset() {
	*x = RequestWalletTokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestWalletTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestWalletTokenInfo) ProtoMessage() {}

func (x *RequestWalletTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_invoicer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestWalletTokenInfo.ProtoReflect.Descriptor instead.
func (*RequestWalletTokenInfo) Descriptor() ([]byte, []int) {
	return file_invoicer_proto_rawDescGZIP(), []int{0}
}

func (x *RequestWalletTokenInfo) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *RequestWalletTokenInfo) GetTokenaddress() string {
	if x != nil {
		return x.Tokenaddress
	}
	return ""
}

type ResponceBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance string `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *ResponceBalance) Reset() {
	*x = ResponceBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceBalance) ProtoMessage() {}

func (x *ResponceBalance) ProtoReflect() protoreflect.Message {
	mi := &file_invoicer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceBalance.ProtoReflect.Descriptor instead.
func (*ResponceBalance) Descriptor() ([]byte, []int) {
	return file_invoicer_proto_rawDescGZIP(), []int{1}
}

func (x *ResponceBalance) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

type RequestWalletInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Signature      []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Publickeybytes []byte `protobuf:"bytes,3,opt,name=publickeybytes,proto3" json:"publickeybytes,omitempty"`
}

func (x *RequestWalletInfo) Reset() {
	*x = RequestWalletInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestWalletInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestWalletInfo) ProtoMessage() {}

func (x *RequestWalletInfo) ProtoReflect() protoreflect.Message {
	mi := &file_invoicer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestWalletInfo.ProtoReflect.Descriptor instead.
func (*RequestWalletInfo) Descriptor() ([]byte, []int) {
	return file_invoicer_proto_rawDescGZIP(), []int{2}
}

func (x *RequestWalletInfo) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *RequestWalletInfo) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *RequestWalletInfo) GetPublickeybytes() []byte {
	if x != nil {
		return x.Publickeybytes
	}
	return nil
}

type ResponceBalanceNonce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance      string `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce        uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	VerifiedSign bool   `protobuf:"varint,3,opt,name=verifiedSign,proto3" json:"verifiedSign,omitempty"`
}

func (x *ResponceBalanceNonce) Reset() {
	*x = ResponceBalanceNonce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponceBalanceNonce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponceBalanceNonce) ProtoMessage() {}

func (x *ResponceBalanceNonce) ProtoReflect() protoreflect.Message {
	mi := &file_invoicer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponceBalanceNonce.ProtoReflect.Descriptor instead.
func (*ResponceBalanceNonce) Descriptor() ([]byte, []int) {
	return file_invoicer_proto_rawDescGZIP(), []int{3}
}

func (x *ResponceBalanceNonce) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *ResponceBalanceNonce) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *ResponceBalanceNonce) GetVerifiedSign() bool {
	if x != nil {
		return x.VerifiedSign
	}
	return false
}

var File_invoicer_proto protoreflect.FileDescriptor

var file_invoicer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x56, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x73, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x79, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x6b, 0x65, 0x79, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x6a, 0x0a, 0x14, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x53, 0x69,
	0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x32, 0x89, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x12, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x17, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x56, 0x65, 0x74, 0x75, 0x73, 0x71, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_invoicer_proto_rawDescOnce sync.Once
	file_invoicer_proto_rawDescData = file_invoicer_proto_rawDesc
)

func file_invoicer_proto_rawDescGZIP() []byte {
	file_invoicer_proto_rawDescOnce.Do(func() {
		file_invoicer_proto_rawDescData = protoimpl.X.CompressGZIP(file_invoicer_proto_rawDescData)
	})
	return file_invoicer_proto_rawDescData
}

var file_invoicer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_invoicer_proto_goTypes = []interface{}{
	(*RequestWalletTokenInfo)(nil), // 0: RequestWalletTokenInfo
	(*ResponceBalance)(nil),        // 1: ResponceBalance
	(*RequestWalletInfo)(nil),      // 2: RequestWalletInfo
	(*ResponceBalanceNonce)(nil),   // 3: ResponceBalanceNonce
}
var file_invoicer_proto_depIdxs = []int32{
	2, // 0: Greeter.GetBalance:input_type -> RequestWalletInfo
	0, // 1: Greeter.StreamGetBalance:input_type -> RequestWalletTokenInfo
	3, // 2: Greeter.GetBalance:output_type -> ResponceBalanceNonce
	1, // 3: Greeter.StreamGetBalance:output_type -> ResponceBalance
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_invoicer_proto_init() }
func file_invoicer_proto_init() {
	if File_invoicer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_invoicer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestWalletTokenInfo); i {
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
		file_invoicer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceBalance); i {
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
		file_invoicer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestWalletInfo); i {
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
		file_invoicer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponceBalanceNonce); i {
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
			RawDescriptor: file_invoicer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_invoicer_proto_goTypes,
		DependencyIndexes: file_invoicer_proto_depIdxs,
		MessageInfos:      file_invoicer_proto_msgTypes,
	}.Build()
	File_invoicer_proto = out.File
	file_invoicer_proto_rawDesc = nil
	file_invoicer_proto_goTypes = nil
	file_invoicer_proto_depIdxs = nil
}
