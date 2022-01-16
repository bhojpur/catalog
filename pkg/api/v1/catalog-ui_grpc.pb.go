// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// CatalogUIClient is the client API for CatalogUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogUIClient interface {
	// ListEngineSpecs returns a list of Catalog Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (CatalogUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type catalogUIClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogUIClient(cc grpc.ClientConnInterface) CatalogUIClient {
	return &catalogUIClient{cc}
}

func (c *catalogUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (CatalogUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CatalogUI_ServiceDesc.Streams[0], "/v1.CatalogUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &catalogUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CatalogUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type catalogUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *catalogUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *catalogUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.CatalogUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogUIServer is the server API for CatalogUI service.
// All implementations must embed UnimplementedCatalogUIServer
// for forward compatibility
type CatalogUIServer interface {
	// ListEngineSpecs returns a list of Catalog Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, CatalogUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedCatalogUIServer()
}

// UnimplementedCatalogUIServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogUIServer struct {
}

func (UnimplementedCatalogUIServer) ListEngineSpecs(*ListEngineSpecsRequest, CatalogUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedCatalogUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedCatalogUIServer) mustEmbedUnimplementedCatalogUIServer() {}

// UnsafeCatalogUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogUIServer will
// result in compilation errors.
type UnsafeCatalogUIServer interface {
	mustEmbedUnimplementedCatalogUIServer()
}

func RegisterCatalogUIServer(s grpc.ServiceRegistrar, srv CatalogUIServer) {
	s.RegisterService(&CatalogUI_ServiceDesc, srv)
}

func _CatalogUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatalogUIServer).ListEngineSpecs(m, &catalogUIListEngineSpecsServer{stream})
}

type CatalogUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type catalogUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *catalogUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CatalogUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CatalogUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CatalogUI_ServiceDesc is the grpc.ServiceDesc for CatalogUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CatalogUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CatalogUI",
	HandlerType: (*CatalogUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _CatalogUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _CatalogUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "catalog-ui.proto",
}
