// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: transactions.proto

package transactions

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Op:
	//
	//	*Request_AuthRequest_
	Op isRequest_Op `protobuf_oneof:"op"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{0}
}

func (m *Request) GetOp() isRequest_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (x *Request) GetAuthRequest() *Request_AuthRequest {
	if x, ok := x.GetOp().(*Request_AuthRequest_); ok {
		return x.AuthRequest
	}
	return nil
}

type isRequest_Op interface {
	isRequest_Op()
}

type Request_AuthRequest_ struct {
	AuthRequest *Request_AuthRequest `protobuf:"bytes,1,opt,name=auth_request,json=authRequest,proto3,oneof"`
}

func (*Request_AuthRequest_) isRequest_Op() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Op:
	//
	//	*Response_TokenResponse_
	//	*Response_MessageCreatedResponse_
	Op isResponse_Op `protobuf_oneof:"op"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{1}
}

func (m *Response) GetOp() isResponse_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (x *Response) GetTokenResponse() *Response_TokenResponse {
	if x, ok := x.GetOp().(*Response_TokenResponse_); ok {
		return x.TokenResponse
	}
	return nil
}

func (x *Response) GetMessageCreatedResponse() *Response_MessageCreatedResponse {
	if x, ok := x.GetOp().(*Response_MessageCreatedResponse_); ok {
		return x.MessageCreatedResponse
	}
	return nil
}

type isResponse_Op interface {
	isResponse_Op()
}

type Response_TokenResponse_ struct {
	TokenResponse *Response_TokenResponse `protobuf:"bytes,1,opt,name=token_response,json=tokenResponse,proto3,oneof"`
}

type Response_MessageCreatedResponse_ struct {
	MessageCreatedResponse *Response_MessageCreatedResponse `protobuf:"bytes,2,opt,name=message_created_response,json=messageCreatedResponse,proto3,oneof"`
}

func (*Response_TokenResponse_) isResponse_Op() {}

func (*Response_MessageCreatedResponse_) isResponse_Op() {}

type Request_AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Say string `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
}

func (x *Request_AuthRequest) Reset() {
	*x = Request_AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request_AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request_AuthRequest) ProtoMessage() {}

func (x *Request_AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request_AuthRequest.ProtoReflect.Descriptor instead.
func (*Request_AuthRequest) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Request_AuthRequest) GetSay() string {
	if x != nil {
		return x.Say
	}
	return ""
}

type Response_TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Response_TokenResponse) Reset() {
	*x = Response_TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_TokenResponse) ProtoMessage() {}

func (x *Response_TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_TokenResponse.ProtoReflect.Descriptor instead.
func (*Response_TokenResponse) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Response_TokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Response_MessageCreatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Response_MessageCreatedResponse) Reset() {
	*x = Response_MessageCreatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_MessageCreatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_MessageCreatedResponse) ProtoMessage() {}

func (x *Response_MessageCreatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_MessageCreatedResponse.ProtoReflect.Descriptor instead.
func (*Response_MessageCreatedResponse) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Response_MessageCreatedResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_transactions_proto protoreflect.FileDescriptor

var file_transactions_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x78, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a,
	0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x61, 0x79, 0x42, 0x04, 0x0a, 0x02, 0x6f, 0x70, 0x22, 0x9b, 0x02, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x18, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x16, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x25, 0x0a, 0x0d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x28, 0x0a, 0x16, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x42, 0x04, 0x0a, 0x02, 0x6f, 0x70, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x72, 0x65, 0x64,
	0x43, 0x68, 0x61, 0x74, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x72, 0x65, 0x64, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transactions_proto_rawDescOnce sync.Once
	file_transactions_proto_rawDescData = file_transactions_proto_rawDesc
)

func file_transactions_proto_rawDescGZIP() []byte {
	file_transactions_proto_rawDescOnce.Do(func() {
		file_transactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_transactions_proto_rawDescData)
	})
	return file_transactions_proto_rawDescData
}

var file_transactions_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_transactions_proto_goTypes = []interface{}{
	(*Request)(nil),                         // 0: transactions.Request
	(*Response)(nil),                        // 1: transactions.Response
	(*Request_AuthRequest)(nil),             // 2: transactions.Request.AuthRequest
	(*Response_TokenResponse)(nil),          // 3: transactions.Response.TokenResponse
	(*Response_MessageCreatedResponse)(nil), // 4: transactions.Response.MessageCreatedResponse
}
var file_transactions_proto_depIdxs = []int32{
	2, // 0: transactions.Request.auth_request:type_name -> transactions.Request.AuthRequest
	3, // 1: transactions.Response.token_response:type_name -> transactions.Response.TokenResponse
	4, // 2: transactions.Response.message_created_response:type_name -> transactions.Response.MessageCreatedResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transactions_proto_init() }
func file_transactions_proto_init() {
	if File_transactions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_transactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_transactions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request_AuthRequest); i {
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
		file_transactions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_TokenResponse); i {
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
		file_transactions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_MessageCreatedResponse); i {
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
	file_transactions_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Request_AuthRequest_)(nil),
	}
	file_transactions_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Response_TokenResponse_)(nil),
		(*Response_MessageCreatedResponse_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transactions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transactions_proto_goTypes,
		DependencyIndexes: file_transactions_proto_depIdxs,
		MessageInfos:      file_transactions_proto_msgTypes,
	}.Build()
	File_transactions_proto = out.File
	file_transactions_proto_rawDesc = nil
	file_transactions_proto_goTypes = nil
	file_transactions_proto_depIdxs = nil
}
