// Code generated by protoc-gen-go. DO NOT EDIT.
// source: timekeeper_rpc.proto

package timekeeper_rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type TimeStartRecord struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Task                 string   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	TimeStart            int64    `protobuf:"varint,3,opt,name=timeStart,proto3" json:"timeStart,omitempty"`
	Comment              string   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeStartRecord) Reset()         { *m = TimeStartRecord{} }
func (m *TimeStartRecord) String() string { return proto.CompactTextString(m) }
func (*TimeStartRecord) ProtoMessage()    {}
func (*TimeStartRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5729d01c2bd5da, []int{0}
}

func (m *TimeStartRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeStartRecord.Unmarshal(m, b)
}
func (m *TimeStartRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeStartRecord.Marshal(b, m, deterministic)
}
func (m *TimeStartRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeStartRecord.Merge(m, src)
}
func (m *TimeStartRecord) XXX_Size() int {
	return xxx_messageInfo_TimeStartRecord.Size(m)
}
func (m *TimeStartRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeStartRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TimeStartRecord proto.InternalMessageInfo

func (m *TimeStartRecord) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *TimeStartRecord) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *TimeStartRecord) GetTimeStart() int64 {
	if m != nil {
		return m.TimeStart
	}
	return 0
}

func (m *TimeStartRecord) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type TimeRecord struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Task                 string   `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`
	TimeStart            int64    `protobuf:"varint,4,opt,name=timeStart,proto3" json:"timeStart,omitempty"`
	TimeEnd              int64    `protobuf:"varint,5,opt,name=timeEnd,proto3" json:"timeEnd,omitempty"`
	Comment              string   `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRecord) Reset()         { *m = TimeRecord{} }
func (m *TimeRecord) String() string { return proto.CompactTextString(m) }
func (*TimeRecord) ProtoMessage()    {}
func (*TimeRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5729d01c2bd5da, []int{1}
}

func (m *TimeRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRecord.Unmarshal(m, b)
}
func (m *TimeRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRecord.Marshal(b, m, deterministic)
}
func (m *TimeRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRecord.Merge(m, src)
}
func (m *TimeRecord) XXX_Size() int {
	return xxx_messageInfo_TimeRecord.Size(m)
}
func (m *TimeRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRecord proto.InternalMessageInfo

func (m *TimeRecord) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *TimeRecord) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *TimeRecord) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *TimeRecord) GetTimeStart() int64 {
	if m != nil {
		return m.TimeStart
	}
	return 0
}

func (m *TimeRecord) GetTimeEnd() int64 {
	if m != nil {
		return m.TimeEnd
	}
	return 0
}

func (m *TimeRecord) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type TimeGroupedRecord struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Task                 string   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	Date                 int64    `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	SpentMinutes         int64    `protobuf:"varint,4,opt,name=spentMinutes,proto3" json:"spentMinutes,omitempty"`
	Comments             string   `protobuf:"bytes,5,opt,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeGroupedRecord) Reset()         { *m = TimeGroupedRecord{} }
func (m *TimeGroupedRecord) String() string { return proto.CompactTextString(m) }
func (*TimeGroupedRecord) ProtoMessage()    {}
func (*TimeGroupedRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5729d01c2bd5da, []int{2}
}

func (m *TimeGroupedRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeGroupedRecord.Unmarshal(m, b)
}
func (m *TimeGroupedRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeGroupedRecord.Marshal(b, m, deterministic)
}
func (m *TimeGroupedRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeGroupedRecord.Merge(m, src)
}
func (m *TimeGroupedRecord) XXX_Size() int {
	return xxx_messageInfo_TimeGroupedRecord.Size(m)
}
func (m *TimeGroupedRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeGroupedRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TimeGroupedRecord proto.InternalMessageInfo

func (m *TimeGroupedRecord) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *TimeGroupedRecord) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *TimeGroupedRecord) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *TimeGroupedRecord) GetSpentMinutes() int64 {
	if m != nil {
		return m.SpentMinutes
	}
	return 0
}

func (m *TimeGroupedRecord) GetComments() string {
	if m != nil {
		return m.Comments
	}
	return ""
}

type GetForDate struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Date                 int64    `protobuf:"varint,2,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetForDate) Reset()         { *m = GetForDate{} }
func (m *GetForDate) String() string { return proto.CompactTextString(m) }
func (*GetForDate) ProtoMessage()    {}
func (*GetForDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5729d01c2bd5da, []int{3}
}

func (m *GetForDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetForDate.Unmarshal(m, b)
}
func (m *GetForDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetForDate.Marshal(b, m, deterministic)
}
func (m *GetForDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetForDate.Merge(m, src)
}
func (m *GetForDate) XXX_Size() int {
	return xxx_messageInfo_GetForDate.Size(m)
}
func (m *GetForDate) XXX_DiscardUnknown() {
	xxx_messageInfo_GetForDate.DiscardUnknown(m)
}

var xxx_messageInfo_GetForDate proto.InternalMessageInfo

func (m *GetForDate) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *GetForDate) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

type RecordID struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ID                   string   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordID) Reset()         { *m = RecordID{} }
func (m *RecordID) String() string { return proto.CompactTextString(m) }
func (*RecordID) ProtoMessage()    {}
func (*RecordID) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5729d01c2bd5da, []int{4}
}

func (m *RecordID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordID.Unmarshal(m, b)
}
func (m *RecordID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordID.Marshal(b, m, deterministic)
}
func (m *RecordID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordID.Merge(m, src)
}
func (m *RecordID) XXX_Size() int {
	return xxx_messageInfo_RecordID.Size(m)
}
func (m *RecordID) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordID.DiscardUnknown(m)
}

var xxx_messageInfo_RecordID proto.InternalMessageInfo

func (m *RecordID) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *RecordID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*TimeStartRecord)(nil), "timekeeper_rpc.TimeStartRecord")
	proto.RegisterType((*TimeRecord)(nil), "timekeeper_rpc.TimeRecord")
	proto.RegisterType((*TimeGroupedRecord)(nil), "timekeeper_rpc.TimeGroupedRecord")
	proto.RegisterType((*GetForDate)(nil), "timekeeper_rpc.GetForDate")
	proto.RegisterType((*RecordID)(nil), "timekeeper_rpc.RecordID")
}

func init() { proto.RegisterFile("timekeeper_rpc.proto", fileDescriptor_9e5729d01c2bd5da) }

var fileDescriptor_9e5729d01c2bd5da = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xbb, 0x49, 0xfe, 0xfd, 0xb7, 0x83, 0x56, 0x3a, 0x88, 0x2c, 0x41, 0xb0, 0xee, 0xa9,
	0xa7, 0x22, 0xf5, 0xe2, 0x49, 0x10, 0xa3, 0xa5, 0x88, 0x20, 0x51, 0xbc, 0x4a, 0x6d, 0xe6, 0x50,
	0x4a, 0xb2, 0x61, 0x33, 0x79, 0x12, 0x6f, 0xbe, 0xa3, 0xef, 0x20, 0xd9, 0x26, 0x6d, 0x53, 0xa3,
	0x82, 0xb7, 0x9d, 0xd9, 0xfd, 0xf6, 0xfb, 0xe5, 0x9b, 0x0d, 0x1c, 0xf2, 0x22, 0xa6, 0x25, 0x51,
	0x4a, 0xe6, 0xc5, 0xa4, 0xf3, 0x51, 0x6a, 0x34, 0x6b, 0xec, 0xd5, 0xbb, 0x2a, 0x87, 0x83, 0xa7,
	0x45, 0x4c, 0x8f, 0x3c, 0x33, 0x1c, 0xd2, 0x5c, 0x9b, 0x08, 0x8f, 0xa0, 0x9d, 0x67, 0x64, 0xa6,
	0x81, 0x14, 0x03, 0x31, 0xec, 0x86, 0x65, 0x85, 0x08, 0x1e, 0xcf, 0xb2, 0xa5, 0x74, 0x6c, 0xd7,
	0xae, 0xf1, 0x18, 0xba, 0x5c, 0xc9, 0xa5, 0x3b, 0x10, 0x43, 0x37, 0xdc, 0x34, 0x50, 0xc2, 0xff,
	0xb9, 0x8e, 0x63, 0x4a, 0x58, 0x7a, 0x56, 0x54, 0x95, 0xea, 0x5d, 0x00, 0x14, 0xbe, 0xa5, 0x65,
	0x0f, 0x9c, 0xb5, 0x9d, 0x33, 0x0d, 0xb6, 0x10, 0x9c, 0x46, 0x04, 0xf7, 0x3b, 0x04, 0xaf, 0x01,
	0xa1, 0x28, 0x6e, 0x92, 0x48, 0xfe, 0xb3, 0x7b, 0x55, 0xb9, 0x0d, 0xd7, 0xae, 0xc3, 0xbd, 0x09,
	0xe8, 0x17, 0x70, 0x13, 0xa3, 0xf3, 0x94, 0xa2, 0x3f, 0xc4, 0x82, 0xe0, 0x45, 0x33, 0xa6, 0x32,
	0x11, 0xbb, 0x46, 0x05, 0x7b, 0x59, 0x4a, 0x09, 0xdf, 0x2f, 0x92, 0x9c, 0x29, 0x2b, 0x51, 0x6b,
	0x3d, 0xf4, 0xa1, 0x53, 0x42, 0x64, 0x16, 0xb7, 0x1b, 0xae, 0x6b, 0x75, 0x01, 0x30, 0x21, 0xbe,
	0xd5, 0x26, 0x28, 0x6e, 0xfb, 0x81, 0xc6, 0x3a, 0x3b, 0x1b, 0x67, 0x35, 0x86, 0xce, 0xea, 0x1b,
	0x6a, 0xc9, 0xd6, 0x75, 0xab, 0x09, 0x38, 0xd5, 0x04, 0xc6, 0x1f, 0x02, 0xf6, 0x8b, 0x0c, 0xee,
	0xec, 0x53, 0x09, 0x1f, 0xae, 0x31, 0x00, 0xf7, 0x2a, 0x8a, 0xf0, 0x64, 0xb4, 0xf3, 0xae, 0x76,
	0x9e, 0x8f, 0xef, 0x37, 0x1d, 0x58, 0xed, 0xa9, 0x16, 0x5e, 0x82, 0x17, 0xe8, 0x84, 0x50, 0xee,
	0x9e, 0xaa, 0x08, 0x7f, 0xd1, 0x3f, 0x43, 0x7f, 0x42, 0x5c, 0x4e, 0xa6, 0x0a, 0xe3, 0x8b, 0x64,
	0x13, 0x94, 0x7f, 0xda, 0x74, 0x5d, 0x6d, 0xb2, 0xaa, 0x75, 0x26, 0x5e, 0xdb, 0xf6, 0xf7, 0x38,
	0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x88, 0x87, 0x2a, 0x06, 0x36, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimeKeeperRPCClient is the client API for TimeKeeperRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimeKeeperRPCClient interface {
	Add(ctx context.Context, in *TimeStartRecord, opts ...grpc.CallOption) (*TimeRecord, error)
	Done(ctx context.Context, in *RecordID, opts ...grpc.CallOption) (*TimeRecord, error)
	GetGroupedForDate(ctx context.Context, in *GetForDate, opts ...grpc.CallOption) (TimeKeeperRPC_GetGroupedForDateClient, error)
}

type timeKeeperRPCClient struct {
	cc *grpc.ClientConn
}

func NewTimeKeeperRPCClient(cc *grpc.ClientConn) TimeKeeperRPCClient {
	return &timeKeeperRPCClient{cc}
}

func (c *timeKeeperRPCClient) Add(ctx context.Context, in *TimeStartRecord, opts ...grpc.CallOption) (*TimeRecord, error) {
	out := new(TimeRecord)
	err := c.cc.Invoke(ctx, "/timekeeper_rpc.TimeKeeperRPC/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeKeeperRPCClient) Done(ctx context.Context, in *RecordID, opts ...grpc.CallOption) (*TimeRecord, error) {
	out := new(TimeRecord)
	err := c.cc.Invoke(ctx, "/timekeeper_rpc.TimeKeeperRPC/Done", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeKeeperRPCClient) GetGroupedForDate(ctx context.Context, in *GetForDate, opts ...grpc.CallOption) (TimeKeeperRPC_GetGroupedForDateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TimeKeeperRPC_serviceDesc.Streams[0], "/timekeeper_rpc.TimeKeeperRPC/GetGroupedForDate", opts...)
	if err != nil {
		return nil, err
	}
	x := &timeKeeperRPCGetGroupedForDateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TimeKeeperRPC_GetGroupedForDateClient interface {
	Recv() (*TimeGroupedRecord, error)
	grpc.ClientStream
}

type timeKeeperRPCGetGroupedForDateClient struct {
	grpc.ClientStream
}

func (x *timeKeeperRPCGetGroupedForDateClient) Recv() (*TimeGroupedRecord, error) {
	m := new(TimeGroupedRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TimeKeeperRPCServer is the server API for TimeKeeperRPC service.
type TimeKeeperRPCServer interface {
	Add(context.Context, *TimeStartRecord) (*TimeRecord, error)
	Done(context.Context, *RecordID) (*TimeRecord, error)
	GetGroupedForDate(*GetForDate, TimeKeeperRPC_GetGroupedForDateServer) error
}

// UnimplementedTimeKeeperRPCServer can be embedded to have forward compatible implementations.
type UnimplementedTimeKeeperRPCServer struct {
}

func (*UnimplementedTimeKeeperRPCServer) Add(ctx context.Context, req *TimeStartRecord) (*TimeRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedTimeKeeperRPCServer) Done(ctx context.Context, req *RecordID) (*TimeRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Done not implemented")
}
func (*UnimplementedTimeKeeperRPCServer) GetGroupedForDate(req *GetForDate, srv TimeKeeperRPC_GetGroupedForDateServer) error {
	return status.Errorf(codes.Unimplemented, "method GetGroupedForDate not implemented")
}

func RegisterTimeKeeperRPCServer(s *grpc.Server, srv TimeKeeperRPCServer) {
	s.RegisterService(&_TimeKeeperRPC_serviceDesc, srv)
}

func _TimeKeeperRPC_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeStartRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeKeeperRPCServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timekeeper_rpc.TimeKeeperRPC/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeKeeperRPCServer).Add(ctx, req.(*TimeStartRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeKeeperRPC_Done_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeKeeperRPCServer).Done(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timekeeper_rpc.TimeKeeperRPC/Done",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeKeeperRPCServer).Done(ctx, req.(*RecordID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeKeeperRPC_GetGroupedForDate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetForDate)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TimeKeeperRPCServer).GetGroupedForDate(m, &timeKeeperRPCGetGroupedForDateServer{stream})
}

type TimeKeeperRPC_GetGroupedForDateServer interface {
	Send(*TimeGroupedRecord) error
	grpc.ServerStream
}

type timeKeeperRPCGetGroupedForDateServer struct {
	grpc.ServerStream
}

func (x *timeKeeperRPCGetGroupedForDateServer) Send(m *TimeGroupedRecord) error {
	return x.ServerStream.SendMsg(m)
}

var _TimeKeeperRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "timekeeper_rpc.TimeKeeperRPC",
	HandlerType: (*TimeKeeperRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _TimeKeeperRPC_Add_Handler,
		},
		{
			MethodName: "Done",
			Handler:    _TimeKeeperRPC_Done_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetGroupedForDate",
			Handler:       _TimeKeeperRPC_GetGroupedForDate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "timekeeper_rpc.proto",
}
