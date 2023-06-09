// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: kvdb/v1/kvdb.proto

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

const (
	Kvdb_ListDB_FullMethodName       = "/api.kvdb.v1.Kvdb/ListDB"
	Kvdb_SearchPrefix_FullMethodName = "/api.kvdb.v1.Kvdb/SearchPrefix"
	Kvdb_Search_FullMethodName       = "/api.kvdb.v1.Kvdb/Search"
)

// KvdbClient is the client API for Kvdb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KvdbClient interface {
	// ListDB 列出所有的db
	ListDB(ctx context.Context, in *ListDBRequest, opts ...grpc.CallOption) (*ListDBReply, error)
	// SearchPrefix 列出某前缀所有的key value
	SearchPrefix(ctx context.Context, in *SearchPrefixRequest, opts ...grpc.CallOption) (*SearchPrefixReply, error)
	// Search 查询key 对应的 value
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
}

type kvdbClient struct {
	cc grpc.ClientConnInterface
}

func NewKvdbClient(cc grpc.ClientConnInterface) KvdbClient {
	return &kvdbClient{cc}
}

func (c *kvdbClient) ListDB(ctx context.Context, in *ListDBRequest, opts ...grpc.CallOption) (*ListDBReply, error) {
	out := new(ListDBReply)
	err := c.cc.Invoke(ctx, Kvdb_ListDB_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvdbClient) SearchPrefix(ctx context.Context, in *SearchPrefixRequest, opts ...grpc.CallOption) (*SearchPrefixReply, error) {
	out := new(SearchPrefixReply)
	err := c.cc.Invoke(ctx, Kvdb_SearchPrefix_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kvdbClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, Kvdb_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KvdbServer is the server API for Kvdb service.
// All implementations must embed UnimplementedKvdbServer
// for forward compatibility
type KvdbServer interface {
	// ListDB 列出所有的db
	ListDB(context.Context, *ListDBRequest) (*ListDBReply, error)
	// SearchPrefix 列出某前缀所有的key value
	SearchPrefix(context.Context, *SearchPrefixRequest) (*SearchPrefixReply, error)
	// Search 查询key 对应的 value
	Search(context.Context, *SearchRequest) (*SearchReply, error)
	mustEmbedUnimplementedKvdbServer()
}

// UnimplementedKvdbServer must be embedded to have forward compatible implementations.
type UnimplementedKvdbServer struct {
}

func (UnimplementedKvdbServer) ListDB(context.Context, *ListDBRequest) (*ListDBReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDB not implemented")
}
func (UnimplementedKvdbServer) SearchPrefix(context.Context, *SearchPrefixRequest) (*SearchPrefixReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPrefix not implemented")
}
func (UnimplementedKvdbServer) Search(context.Context, *SearchRequest) (*SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedKvdbServer) mustEmbedUnimplementedKvdbServer() {}

// UnsafeKvdbServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KvdbServer will
// result in compilation errors.
type UnsafeKvdbServer interface {
	mustEmbedUnimplementedKvdbServer()
}

func RegisterKvdbServer(s grpc.ServiceRegistrar, srv KvdbServer) {
	s.RegisterService(&Kvdb_ServiceDesc, srv)
}

func _Kvdb_ListDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvdbServer).ListDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kvdb_ListDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvdbServer).ListDB(ctx, req.(*ListDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kvdb_SearchPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvdbServer).SearchPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kvdb_SearchPrefix_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvdbServer).SearchPrefix(ctx, req.(*SearchPrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kvdb_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KvdbServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Kvdb_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KvdbServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Kvdb_ServiceDesc is the grpc.ServiceDesc for Kvdb service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kvdb_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.kvdb.v1.Kvdb",
	HandlerType: (*KvdbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDB",
			Handler:    _Kvdb_ListDB_Handler,
		},
		{
			MethodName: "SearchPrefix",
			Handler:    _Kvdb_SearchPrefix_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Kvdb_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kvdb/v1/kvdb.proto",
}
