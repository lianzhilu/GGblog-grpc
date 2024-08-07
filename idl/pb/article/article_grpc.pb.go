// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: article.proto

package articlePb

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

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	// 添加文章
	ArticleCreate(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*ArticleCommonResponse, error)
	// 查询单个文章的详细信息
	ArticleDetail(ctx context.Context, in *ArticleTargetRequest, opts ...grpc.CallOption) (*ArticleDetailResponse, error)
	// 根据关键字搜索文章
	ArticleSearchBykeyword(ctx context.Context, in *ArticleSearchRequest, opts ...grpc.CallOption) (*ArticleSearchResponse, error)
	// 查询某个用户发布的所有文章
	ArticleSearchByUser(ctx context.Context, in *ArticleSearchRequest, opts ...grpc.CallOption) (*ArticleSearchResponse, error)
	// 查询文章列表
	ArticleList(ctx context.Context, in *ArticleSearchRequest, opts ...grpc.CallOption) (*ArticleSearchResponse, error)
	// 编辑文章
	ArticleEdit(ctx context.Context, in *ArticleTargetRequest, opts ...grpc.CallOption) (*ArticleDetailResponse, error)
	// 删除文章
	ArticleDelete(ctx context.Context, in *ArticleTargetRequest, opts ...grpc.CallOption) (*ArticleCommonResponse, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) ArticleCreate(ctx context.Context, in *ArticleRequest, opts ...grpc.CallOption) (*ArticleCommonResponse, error) {
	out := new(ArticleCommonResponse)
	err := c.cc.Invoke(ctx, "/articlePb.ArticleService/ArticleCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ArticleDetail(ctx context.Context, in *ArticleTargetRequest, opts ...grpc.CallOption) (*ArticleDetailResponse, error) {
	out := new(ArticleDetailResponse)
	err := c.cc.Invoke(ctx, "/articlePb.ArticleService/ArticleDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ArticleSearchBykeyword(ctx context.Context, in *ArticleSearchRequest, opts ...grpc.CallOption) (*ArticleSearchResponse, error) {
	out := new(ArticleSearchResponse)
	err := c.cc.Invoke(ctx, "/articlePb.ArticleService/ArticleSearchBykeyword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ArticleSearchByUser(ctx context.Context, in *ArticleSearchRequest, opts ...grpc.CallOption) (*ArticleSearchResponse, error) {
	out := new(ArticleSearchResponse)
	err := c.cc.Invoke(ctx, "/articlePb.ArticleService/ArticleSearchByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ArticleList(ctx context.Context, in *ArticleSearchRequest, opts ...grpc.CallOption) (*ArticleSearchResponse, error) {
	out := new(ArticleSearchResponse)
	err := c.cc.Invoke(ctx, "/articlePb.ArticleService/ArticleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ArticleEdit(ctx context.Context, in *ArticleTargetRequest, opts ...grpc.CallOption) (*ArticleDetailResponse, error) {
	out := new(ArticleDetailResponse)
	err := c.cc.Invoke(ctx, "/articlePb.ArticleService/ArticleEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) ArticleDelete(ctx context.Context, in *ArticleTargetRequest, opts ...grpc.CallOption) (*ArticleCommonResponse, error) {
	out := new(ArticleCommonResponse)
	err := c.cc.Invoke(ctx, "/articlePb.ArticleService/ArticleDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	// 添加文章
	ArticleCreate(context.Context, *ArticleRequest) (*ArticleCommonResponse, error)
	// 查询单个文章的详细信息
	ArticleDetail(context.Context, *ArticleTargetRequest) (*ArticleDetailResponse, error)
	// 根据关键字搜索文章
	ArticleSearchBykeyword(context.Context, *ArticleSearchRequest) (*ArticleSearchResponse, error)
	// 查询某个用户发布的所有文章
	ArticleSearchByUser(context.Context, *ArticleSearchRequest) (*ArticleSearchResponse, error)
	// 查询文章列表
	ArticleList(context.Context, *ArticleSearchRequest) (*ArticleSearchResponse, error)
	// 编辑文章
	ArticleEdit(context.Context, *ArticleTargetRequest) (*ArticleDetailResponse, error)
	// 删除文章
	ArticleDelete(context.Context, *ArticleTargetRequest) (*ArticleCommonResponse, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) ArticleCreate(context.Context, *ArticleRequest) (*ArticleCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleCreate not implemented")
}
func (UnimplementedArticleServiceServer) ArticleDetail(context.Context, *ArticleTargetRequest) (*ArticleDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleDetail not implemented")
}
func (UnimplementedArticleServiceServer) ArticleSearchBykeyword(context.Context, *ArticleSearchRequest) (*ArticleSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleSearchBykeyword not implemented")
}
func (UnimplementedArticleServiceServer) ArticleSearchByUser(context.Context, *ArticleSearchRequest) (*ArticleSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleSearchByUser not implemented")
}
func (UnimplementedArticleServiceServer) ArticleList(context.Context, *ArticleSearchRequest) (*ArticleSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleList not implemented")
}
func (UnimplementedArticleServiceServer) ArticleEdit(context.Context, *ArticleTargetRequest) (*ArticleDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleEdit not implemented")
}
func (UnimplementedArticleServiceServer) ArticleDelete(context.Context, *ArticleTargetRequest) (*ArticleCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleDelete not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_ArticleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ArticleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlePb.ArticleService/ArticleCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ArticleCreate(ctx, req.(*ArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ArticleDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ArticleDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlePb.ArticleService/ArticleDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ArticleDetail(ctx, req.(*ArticleTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ArticleSearchBykeyword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ArticleSearchBykeyword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlePb.ArticleService/ArticleSearchBykeyword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ArticleSearchBykeyword(ctx, req.(*ArticleSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ArticleSearchByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ArticleSearchByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlePb.ArticleService/ArticleSearchByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ArticleSearchByUser(ctx, req.(*ArticleSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlePb.ArticleService/ArticleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ArticleList(ctx, req.(*ArticleSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ArticleEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ArticleEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlePb.ArticleService/ArticleEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ArticleEdit(ctx, req.(*ArticleTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_ArticleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).ArticleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articlePb.ArticleService/ArticleDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).ArticleDelete(ctx, req.(*ArticleTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "articlePb.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ArticleCreate",
			Handler:    _ArticleService_ArticleCreate_Handler,
		},
		{
			MethodName: "ArticleDetail",
			Handler:    _ArticleService_ArticleDetail_Handler,
		},
		{
			MethodName: "ArticleSearchBykeyword",
			Handler:    _ArticleService_ArticleSearchBykeyword_Handler,
		},
		{
			MethodName: "ArticleSearchByUser",
			Handler:    _ArticleService_ArticleSearchByUser_Handler,
		},
		{
			MethodName: "ArticleList",
			Handler:    _ArticleService_ArticleList_Handler,
		},
		{
			MethodName: "ArticleEdit",
			Handler:    _ArticleService_ArticleEdit_Handler,
		},
		{
			MethodName: "ArticleDelete",
			Handler:    _ArticleService_ArticleDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
