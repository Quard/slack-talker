// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal_api.proto

package poindexter

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type URL struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URL) Reset()         { *m = URL{} }
func (m *URL) String() string { return proto.CompactTextString(m) }
func (*URL) ProtoMessage()    {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_927ea7049a2d0eaa, []int{0}
}

func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (m *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(m, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *URL) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ID struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_927ea7049a2d0eaa, []int{1}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type User struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_927ea7049a2d0eaa, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ReadingListRecord struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	ImageUrl             string   `protobuf:"bytes,5,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	Created              int64    `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	IsRead               bool     `protobuf:"varint,7,opt,name=isRead,proto3" json:"isRead,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadingListRecord) Reset()         { *m = ReadingListRecord{} }
func (m *ReadingListRecord) String() string { return proto.CompactTextString(m) }
func (*ReadingListRecord) ProtoMessage()    {}
func (*ReadingListRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_927ea7049a2d0eaa, []int{3}
}

func (m *ReadingListRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadingListRecord.Unmarshal(m, b)
}
func (m *ReadingListRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadingListRecord.Marshal(b, m, deterministic)
}
func (m *ReadingListRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadingListRecord.Merge(m, src)
}
func (m *ReadingListRecord) XXX_Size() int {
	return xxx_messageInfo_ReadingListRecord.Size(m)
}
func (m *ReadingListRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadingListRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ReadingListRecord proto.InternalMessageInfo

func (m *ReadingListRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadingListRecord) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ReadingListRecord) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ReadingListRecord) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ReadingListRecord) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *ReadingListRecord) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ReadingListRecord) GetIsRead() bool {
	if m != nil {
		return m.IsRead
	}
	return false
}

func init() {
	proto.RegisterType((*URL)(nil), "internal_api.URL")
	proto.RegisterType((*ID)(nil), "internal_api.ID")
	proto.RegisterType((*User)(nil), "internal_api.User")
	proto.RegisterType((*ReadingListRecord)(nil), "internal_api.ReadingListRecord")
}

func init() { proto.RegisterFile("internal_api.proto", fileDescriptor_927ea7049a2d0eaa) }

var fileDescriptor_927ea7049a2d0eaa = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x86, 0xbf, 0x24, 0xf3, 0xf7, 0x1d, 0x45, 0x66, 0x0e, 0x32, 0x84, 0xd9, 0x58, 0xba, 0xea,
	0x42, 0x46, 0xd1, 0x9d, 0x2e, 0xa4, 0xd0, 0x4d, 0xa0, 0x82, 0x04, 0xba, 0x96, 0x38, 0x09, 0x43,
	0xb0, 0xb6, 0x43, 0x9a, 0xb9, 0x09, 0xef, 0xc6, 0x3b, 0x94, 0xd6, 0xb6, 0x4c, 0x15, 0x61, 0x70,
	0xd7, 0xf7, 0xe1, 0x3c, 0x9c, 0xb7, 0x87, 0x00, 0xda, 0xc2, 0x1b, 0x57, 0xa8, 0xfc, 0x59, 0xed,
	0xec, 0x7a, 0xe7, 0x4a, 0x5f, 0xe2, 0xe9, 0x21, 0x0b, 0xaf, 0x80, 0x65, 0x32, 0xc5, 0x25, 0x4c,
	0xf6, 0x95, 0x71, 0x22, 0xe1, 0x24, 0x20, 0xd1, 0x7f, 0xd9, 0x26, 0x9c, 0x03, 0xdb, 0xbb, 0x9c,
	0xd3, 0x06, 0xd6, 0x9f, 0xe1, 0x25, 0x50, 0x91, 0xfc, 0x3a, 0x7f, 0x06, 0xd4, 0xea, 0x76, 0x9c,
	0x5a, 0x1d, 0x2e, 0x61, 0x94, 0x55, 0xc6, 0xd5, 0xbc, 0x9f, 0xa5, 0x22, 0x09, 0x3f, 0x08, 0x2c,
	0xa4, 0x51, 0xda, 0x16, 0xdb, 0xd4, 0x56, 0x5e, 0x9a, 0x4d, 0xe9, 0x74, 0x6b, 0x93, 0xce, 0x3e,
	0xd8, 0x42, 0x07, 0x5b, 0xce, 0x61, 0xec, 0xad, 0xcf, 0x0d, 0x67, 0x0d, 0xfe, 0x0a, 0x5d, 0xd7,
	0x51, 0xdf, 0x15, 0x57, 0x30, 0xb3, 0x6f, 0x6a, 0x6b, 0x32, 0x97, 0xf3, 0x71, 0x83, 0xfb, 0x8c,
	0x1c, 0xa6, 0x1b, 0x67, 0x94, 0x37, 0x9a, 0x4f, 0x02, 0x12, 0x31, 0xd9, 0xc5, 0x7a, 0xab, 0xad,
	0xea, 0x72, 0x7c, 0x1a, 0x90, 0x68, 0x26, 0xdb, 0x74, 0xf3, 0x4e, 0xe1, 0x44, 0xb4, 0xb7, 0x8b,
	0x9f, 0x04, 0xde, 0x03, 0x8b, 0xb5, 0xc6, 0xc5, 0x7a, 0x70, 0xe4, 0x4c, 0xa6, 0xab, 0x8b, 0x21,
	0xfa, 0xf1, 0xa3, 0xe1, 0x3f, 0x7c, 0x80, 0x51, 0x9d, 0x11, 0xbf, 0xd9, 0x95, 0x71, 0x47, 0xe8,
	0xd7, 0x04, 0x63, 0x80, 0x47, 0xe5, 0x5e, 0xe3, 0xa6, 0x1b, 0xce, 0x87, 0x8a, 0x48, 0x8e, 0xe9,
	0x70, 0x07, 0x2c, 0x31, 0xf9, 0x9f, 0xdc, 0x97, 0x49, 0xf3, 0x98, 0x6e, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xc8, 0xbf, 0xf2, 0x7c, 0x62, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InternalAPIClient is the client API for InternalAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalAPIClient interface {
	Add(ctx context.Context, in *URL, opts ...grpc.CallOption) (*ReadingListRecord, error)
	List(ctx context.Context, in *User, opts ...grpc.CallOption) (InternalAPI_ListClient, error)
	MarkAsRead(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ReadingListRecord, error)
	Del(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ReadingListRecord, error)
}

type internalAPIClient struct {
	cc *grpc.ClientConn
}

func NewInternalAPIClient(cc *grpc.ClientConn) InternalAPIClient {
	return &internalAPIClient{cc}
}

func (c *internalAPIClient) Add(ctx context.Context, in *URL, opts ...grpc.CallOption) (*ReadingListRecord, error) {
	out := new(ReadingListRecord)
	err := c.cc.Invoke(ctx, "/internal_api.InternalAPI/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAPIClient) List(ctx context.Context, in *User, opts ...grpc.CallOption) (InternalAPI_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_InternalAPI_serviceDesc.Streams[0], "/internal_api.InternalAPI/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &internalAPIListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InternalAPI_ListClient interface {
	Recv() (*ReadingListRecord, error)
	grpc.ClientStream
}

type internalAPIListClient struct {
	grpc.ClientStream
}

func (x *internalAPIListClient) Recv() (*ReadingListRecord, error) {
	m := new(ReadingListRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *internalAPIClient) MarkAsRead(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ReadingListRecord, error) {
	out := new(ReadingListRecord)
	err := c.cc.Invoke(ctx, "/internal_api.InternalAPI/MarkAsRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAPIClient) Del(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ReadingListRecord, error) {
	out := new(ReadingListRecord)
	err := c.cc.Invoke(ctx, "/internal_api.InternalAPI/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalAPIServer is the server API for InternalAPI service.
type InternalAPIServer interface {
	Add(context.Context, *URL) (*ReadingListRecord, error)
	List(*User, InternalAPI_ListServer) error
	MarkAsRead(context.Context, *ID) (*ReadingListRecord, error)
	Del(context.Context, *ID) (*ReadingListRecord, error)
}

// UnimplementedInternalAPIServer can be embedded to have forward compatible implementations.
type UnimplementedInternalAPIServer struct {
}

func (*UnimplementedInternalAPIServer) Add(ctx context.Context, req *URL) (*ReadingListRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedInternalAPIServer) List(req *User, srv InternalAPI_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedInternalAPIServer) MarkAsRead(ctx context.Context, req *ID) (*ReadingListRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsRead not implemented")
}
func (*UnimplementedInternalAPIServer) Del(ctx context.Context, req *ID) (*ReadingListRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}

func RegisterInternalAPIServer(s *grpc.Server, srv InternalAPIServer) {
	s.RegisterService(&_InternalAPI_serviceDesc, srv)
}

func _InternalAPI_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAPIServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal_api.InternalAPI/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAPIServer).Add(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAPI_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InternalAPIServer).List(m, &internalAPIListServer{stream})
}

type InternalAPI_ListServer interface {
	Send(*ReadingListRecord) error
	grpc.ServerStream
}

type internalAPIListServer struct {
	grpc.ServerStream
}

func (x *internalAPIListServer) Send(m *ReadingListRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _InternalAPI_MarkAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAPIServer).MarkAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal_api.InternalAPI/MarkAsRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAPIServer).MarkAsRead(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAPI_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAPIServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal_api.InternalAPI/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAPIServer).Del(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal_api.InternalAPI",
	HandlerType: (*InternalAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _InternalAPI_Add_Handler,
		},
		{
			MethodName: "MarkAsRead",
			Handler:    _InternalAPI_MarkAsRead_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _InternalAPI_Del_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _InternalAPI_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal_api.proto",
}
