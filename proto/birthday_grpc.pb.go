// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package birthday_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BirthdayFunctionsClient is the client API for BirthdayFunctions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BirthdayFunctionsClient interface {
	// CreateBirthday creates a new birthday if it doesn't exist.
	// and if it does, its overwriting it
	CreateBirthday(ctx context.Context, in *CreateBirthdayRequest, opts ...grpc.CallOption) (*BirthdayObject, error)
	// GetBirthday returns the selected birthday using personalNumber parameter.
	GetBirthday(ctx context.Context, in *GetBirthdayRequest, opts ...grpc.CallOption) (*BirthdayObject, error)
	// GetAllBirthday returns all birthdays.
	GetAllBirthdays(ctx context.Context, in *GetAllBirthdaysRequest, opts ...grpc.CallOption) (*GetAllBirthdaysResponse, error)
	// DeleteBirthday deletes the selected birthday.
	DeleteBirthday(ctx context.Context, in *DeleteBirthdayRequest, opts ...grpc.CallOption) (*DeleteBirthdayResponse, error)
}

type birthdayFunctionsClient struct {
	cc grpc.ClientConnInterface
}

func NewBirthdayFunctionsClient(cc grpc.ClientConnInterface) BirthdayFunctionsClient {
	return &birthdayFunctionsClient{cc}
}

func (c *birthdayFunctionsClient) CreateBirthday(ctx context.Context, in *CreateBirthdayRequest, opts ...grpc.CallOption) (*BirthdayObject, error) {
	out := new(BirthdayObject)
	err := c.cc.Invoke(ctx, "/birthday.BirthdayFunctions/CreateBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birthdayFunctionsClient) GetBirthday(ctx context.Context, in *GetBirthdayRequest, opts ...grpc.CallOption) (*BirthdayObject, error) {
	out := new(BirthdayObject)
	err := c.cc.Invoke(ctx, "/birthday.BirthdayFunctions/GetBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birthdayFunctionsClient) GetAllBirthdays(ctx context.Context, in *GetAllBirthdaysRequest, opts ...grpc.CallOption) (*GetAllBirthdaysResponse, error) {
	out := new(GetAllBirthdaysResponse)
	err := c.cc.Invoke(ctx, "/birthday.BirthdayFunctions/GetAllBirthdays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *birthdayFunctionsClient) DeleteBirthday(ctx context.Context, in *DeleteBirthdayRequest, opts ...grpc.CallOption) (*DeleteBirthdayResponse, error) {
	out := new(DeleteBirthdayResponse)
	err := c.cc.Invoke(ctx, "/birthday.BirthdayFunctions/DeleteBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BirthdayFunctionsServer is the server API for BirthdayFunctions service.
// All implementations must embed UnimplementedBirthdayFunctionsServer
// for forward compatibility
type BirthdayFunctionsServer interface {
	// CreateBirthday creates a new birthday if it doesn't exist.
	// and if it does, its overwriting it
	CreateBirthday(context.Context, *CreateBirthdayRequest) (*BirthdayObject, error)
	// GetBirthday returns the selected birthday using personalNumber parameter.
	GetBirthday(context.Context, *GetBirthdayRequest) (*BirthdayObject, error)
	// GetAllBirthday returns all birthdays.
	GetAllBirthdays(context.Context, *GetAllBirthdaysRequest) (*GetAllBirthdaysResponse, error)
	// DeleteBirthday deletes the selected birthday.
	DeleteBirthday(context.Context, *DeleteBirthdayRequest) (*DeleteBirthdayResponse, error)
	mustEmbedUnimplementedBirthdayFunctionsServer()
}

// UnimplementedBirthdayFunctionsServer must be embedded to have forward compatible implementations.
type UnimplementedBirthdayFunctionsServer struct {
}

func (UnimplementedBirthdayFunctionsServer) CreateBirthday(context.Context, *CreateBirthdayRequest) (*BirthdayObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBirthday not implemented")
}
func (UnimplementedBirthdayFunctionsServer) GetBirthday(context.Context, *GetBirthdayRequest) (*BirthdayObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBirthday not implemented")
}
func (UnimplementedBirthdayFunctionsServer) GetAllBirthdays(context.Context, *GetAllBirthdaysRequest) (*GetAllBirthdaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBirthdays not implemented")
}
func (UnimplementedBirthdayFunctionsServer) DeleteBirthday(context.Context, *DeleteBirthdayRequest) (*DeleteBirthdayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBirthday not implemented")
}
func (UnimplementedBirthdayFunctionsServer) mustEmbedUnimplementedBirthdayFunctionsServer() {}

// UnsafeBirthdayFunctionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BirthdayFunctionsServer will
// result in compilation errors.
type UnsafeBirthdayFunctionsServer interface {
	mustEmbedUnimplementedBirthdayFunctionsServer()
}

func RegisterBirthdayFunctionsServer(s grpc.ServiceRegistrar, srv BirthdayFunctionsServer) {
	s.RegisterService(&BirthdayFunctions_ServiceDesc, srv)
}

func _BirthdayFunctions_CreateBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBirthdayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirthdayFunctionsServer).CreateBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/birthday.BirthdayFunctions/CreateBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirthdayFunctionsServer).CreateBirthday(ctx, req.(*CreateBirthdayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirthdayFunctions_GetBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBirthdayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirthdayFunctionsServer).GetBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/birthday.BirthdayFunctions/GetBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirthdayFunctionsServer).GetBirthday(ctx, req.(*GetBirthdayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirthdayFunctions_GetAllBirthdays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBirthdaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirthdayFunctionsServer).GetAllBirthdays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/birthday.BirthdayFunctions/GetAllBirthdays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirthdayFunctionsServer).GetAllBirthdays(ctx, req.(*GetAllBirthdaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BirthdayFunctions_DeleteBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBirthdayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BirthdayFunctionsServer).DeleteBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/birthday.BirthdayFunctions/DeleteBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BirthdayFunctionsServer).DeleteBirthday(ctx, req.(*DeleteBirthdayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BirthdayFunctions_ServiceDesc is the grpc.ServiceDesc for BirthdayFunctions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BirthdayFunctions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "birthday.BirthdayFunctions",
	HandlerType: (*BirthdayFunctionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBirthday",
			Handler:    _BirthdayFunctions_CreateBirthday_Handler,
		},
		{
			MethodName: "GetBirthday",
			Handler:    _BirthdayFunctions_GetBirthday_Handler,
		},
		{
			MethodName: "GetAllBirthdays",
			Handler:    _BirthdayFunctions_GetAllBirthdays_Handler,
		},
		{
			MethodName: "DeleteBirthday",
			Handler:    _BirthdayFunctions_DeleteBirthday_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "birthday.proto",
}
