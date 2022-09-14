// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package grpc

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Users struct {
	User                 []*User  `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

type GetUserResponse struct {
	Users                *Users   `protobuf:"bytes,1,opt,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUsers() *Users {
	if m != nil {
		return m.Users
	}
	return nil
}

type FindUserByIdRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserByIdRequest) Reset()         { *m = FindUserByIdRequest{} }
func (m *FindUserByIdRequest) String() string { return proto.CompactTextString(m) }
func (*FindUserByIdRequest) ProtoMessage()    {}
func (*FindUserByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *FindUserByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserByIdRequest.Unmarshal(m, b)
}
func (m *FindUserByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserByIdRequest.Marshal(b, m, deterministic)
}
func (m *FindUserByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserByIdRequest.Merge(m, src)
}
func (m *FindUserByIdRequest) XXX_Size() int {
	return xxx_messageInfo_FindUserByIdRequest.Size(m)
}
func (m *FindUserByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserByIdRequest proto.InternalMessageInfo

func (m *FindUserByIdRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type FindUserByIdResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindUserByIdResponse) Reset()         { *m = FindUserByIdResponse{} }
func (m *FindUserByIdResponse) String() string { return proto.CompactTextString(m) }
func (*FindUserByIdResponse) ProtoMessage()    {}
func (*FindUserByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *FindUserByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindUserByIdResponse.Unmarshal(m, b)
}
func (m *FindUserByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindUserByIdResponse.Marshal(b, m, deterministic)
}
func (m *FindUserByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindUserByIdResponse.Merge(m, src)
}
func (m *FindUserByIdResponse) XXX_Size() int {
	return xxx_messageInfo_FindUserByIdResponse.Size(m)
}
func (m *FindUserByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindUserByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindUserByIdResponse proto.InternalMessageInfo

func (m *FindUserByIdResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type SignUpRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpRequest) Reset()         { *m = SignUpRequest{} }
func (m *SignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignUpRequest) ProtoMessage()    {}
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *SignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpRequest.Unmarshal(m, b)
}
func (m *SignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpRequest.Merge(m, src)
}
func (m *SignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignUpRequest.Size(m)
}
func (m *SignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpRequest proto.InternalMessageInfo

func (m *SignUpRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignUpRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignUpResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResponse) Reset()         { *m = SignUpResponse{} }
func (m *SignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignUpResponse) ProtoMessage()    {}
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *SignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResponse.Unmarshal(m, b)
}
func (m *SignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResponse.Marshal(b, m, deterministic)
}
func (m *SignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResponse.Merge(m, src)
}
func (m *SignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignUpResponse.Size(m)
}
func (m *SignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResponse proto.InternalMessageInfo

func (m *SignUpResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "grpc.User")
	proto.RegisterType((*Users)(nil), "grpc.Users")
	proto.RegisterType((*GetUserRequest)(nil), "grpc.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "grpc.GetUserResponse")
	proto.RegisterType((*FindUserByIdRequest)(nil), "grpc.FindUserByIdRequest")
	proto.RegisterType((*FindUserByIdResponse)(nil), "grpc.FindUserByIdResponse")
	proto.RegisterType((*SignUpRequest)(nil), "grpc.SignUpRequest")
	proto.RegisterType((*SignUpResponse)(nil), "grpc.SignUpResponse")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf)
}

var fileDescriptor_116e343673f7ffaf = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x75, 0x29, 0x25, 0x75, 0x50, 0x34, 0x53, 0x34, 0xc8, 0xc1, 0xd4, 0xbd, 0xd8, 0x8b, 0xc4,
	0xe0, 0x47, 0xe2, 0xd5, 0x83, 0xa4, 0x17, 0x0f, 0x34, 0xfd, 0x01, 0xb5, 0x6c, 0x1a, 0x0e, 0x02,
	0xee, 0x50, 0x8d, 0xff, 0xd0, 0x9f, 0x65, 0x96, 0x05, 0x5a, 0x1a, 0x62, 0x3c, 0xce, 0x7b, 0x6f,
	0xdf, 0x9b, 0x79, 0x59, 0x80, 0x0d, 0x09, 0x19, 0x14, 0x32, 0x2f, 0x73, 0x34, 0xd7, 0xb2, 0x58,
	0xf1, 0x57, 0x30, 0x17, 0x24, 0x24, 0x3a, 0x60, 0xa4, 0x89, 0xc7, 0x26, 0x6c, 0x7a, 0x18, 0x1b,
	0x69, 0x82, 0x3e, 0x8c, 0x94, 0x36, 0x5b, 0xbe, 0x0b, 0xcf, 0xa8, 0xd0, 0x76, 0x56, 0x5c, 0xb1,
	0x24, 0xfa, 0xca, 0x65, 0xe2, 0x0d, 0x34, 0xd7, 0xcc, 0xfc, 0x1a, 0x86, 0xca, 0x8f, 0xf0, 0x12,
	0x4c, 0xf5, 0xc0, 0x63, 0x93, 0xc1, 0xd4, 0x0e, 0x21, 0x50, 0x69, 0x81, 0xa2, 0xe2, 0x0a, 0xe7,
	0xa7, 0xe0, 0x44, 0xa2, 0xac, 0x00, 0xf1, 0xb1, 0x11, 0x54, 0xf2, 0x7b, 0x38, 0x69, 0x11, 0x2a,
	0xf2, 0x8c, 0x04, 0x5e, 0xc1, 0x50, 0x89, 0xa9, 0x5a, 0xcc, 0x0e, 0xed, 0xad, 0x0b, 0xc5, 0x9a,
	0xe1, 0x37, 0x30, 0x7e, 0x49, 0xb3, 0x44, 0x61, 0xcf, 0xdf, 0xb3, 0xa4, 0x36, 0xc3, 0x73, 0xb0,
	0x14, 0x3f, 0x6b, 0x6e, 0xaa, 0x27, 0xfe, 0x08, 0x6e, 0x57, 0x5e, 0x27, 0x6d, 0xd7, 0x65, 0xbd,
	0xeb, 0x46, 0x70, 0x3c, 0x4f, 0xd7, 0xd9, 0xa2, 0x68, 0x02, 0x76, 0x0b, 0x62, 0x7f, 0x14, 0x64,
	0xec, 0x15, 0x74, 0x0b, 0x4e, 0x63, 0xf4, 0xbf, 0xe8, 0xf0, 0x87, 0x81, 0xad, 0xc6, 0xb9, 0x90,
	0x9f, 0xe9, 0x4a, 0xe0, 0x13, 0x8c, 0xea, 0x9e, 0x08, 0x5d, 0xad, 0xee, 0x36, 0xe9, 0x9f, 0xed,
	0xa1, 0x3a, 0x88, 0x1f, 0x60, 0x04, 0x47, 0xbb, 0xd7, 0xe3, 0x85, 0x16, 0xf6, 0x14, 0xe8, 0xfb,
	0x7d, 0x54, 0x6b, 0xf4, 0x00, 0x96, 0xbe, 0x02, 0xc7, 0x5a, 0xd7, 0x29, 0xc7, 0x77, 0xbb, 0x60,
	0xf3, 0xec, 0xcd, 0xaa, 0xbe, 0xde, 0xdd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0x71, 0x99,
	0x54, 0x88, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUsers(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	FindUserById(ctx context.Context, in *FindUserByIdRequest, opts ...grpc.CallOption) (*FindUserByIdResponse, error)
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/grpc.UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindUserById(ctx context.Context, in *FindUserByIdRequest, opts ...grpc.CallOption) (*FindUserByIdResponse, error) {
	out := new(FindUserByIdResponse)
	err := c.cc.Invoke(ctx, "/grpc.UserService/FindUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/grpc.UserService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUsers(context.Context, *GetUserRequest) (*GetUserResponse, error)
	FindUserById(context.Context, *FindUserByIdRequest) (*FindUserByIdResponse, error)
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUsers(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (*UnimplementedUserServiceServer) FindUserById(ctx context.Context, req *FindUserByIdRequest) (*FindUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserById not implemented")
}
func (*UnimplementedUserServiceServer) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/FindUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindUserById(ctx, req.(*FindUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.UserService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "FindUserById",
			Handler:    _UserService_FindUserById_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _UserService_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}