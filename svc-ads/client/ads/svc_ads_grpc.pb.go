// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ads

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdServiceClient is the client API for AdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdServiceClient interface {
	GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*GetAdResponse, error)
	ListAds(ctx context.Context, in *ListAdsRequest, opts ...grpc.CallOption) (*ListAdsResponse, error)
	SearchAds(ctx context.Context, in *SearchAdsRequest, opts ...grpc.CallOption) (*SearchAdsResponse, error)
	CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*Ad, error)
	UpdateAd(ctx context.Context, in *UpdateAdRequest, opts ...grpc.CallOption) (*Ad, error)
	DeleteAd(ctx context.Context, in *DeleteAdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListPhotos(ctx context.Context, in *ListPhotosRequest, opts ...grpc.CallOption) (*ListPhotosResponse, error)
	ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error)
}

type adServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdServiceClient(cc grpc.ClientConnInterface) AdServiceClient {
	return &adServiceClient{cc}
}

func (c *adServiceClient) GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*GetAdResponse, error) {
	out := new(GetAdResponse)
	err := c.cc.Invoke(ctx, "/svc_ads.AdService/GetAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) ListAds(ctx context.Context, in *ListAdsRequest, opts ...grpc.CallOption) (*ListAdsResponse, error) {
	out := new(ListAdsResponse)
	err := c.cc.Invoke(ctx, "/svc_ads.AdService/ListAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) SearchAds(ctx context.Context, in *SearchAdsRequest, opts ...grpc.CallOption) (*SearchAdsResponse, error) {
	out := new(SearchAdsResponse)
	err := c.cc.Invoke(ctx, "/svc_ads.AdService/SearchAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) CreateAd(ctx context.Context, in *CreateAdRequest, opts ...grpc.CallOption) (*Ad, error) {
	out := new(Ad)
	err := c.cc.Invoke(ctx, "/svc_ads.AdService/CreateAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) UpdateAd(ctx context.Context, in *UpdateAdRequest, opts ...grpc.CallOption) (*Ad, error) {
	out := new(Ad)
	err := c.cc.Invoke(ctx, "/svc_ads.AdService/UpdateAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) DeleteAd(ctx context.Context, in *DeleteAdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/svc_ads.AdService/DeleteAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) ListPhotos(ctx context.Context, in *ListPhotosRequest, opts ...grpc.CallOption) (*ListPhotosResponse, error) {
	out := new(ListPhotosResponse)
	err := c.cc.Invoke(ctx, "/svc_ads.AdService/ListPhotos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	err := c.cc.Invoke(ctx, "/svc_ads.AdService/ListTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdServiceServer is the server API for AdService service.
// All implementations must embed UnimplementedAdServiceServer
// for forward compatibility
type AdServiceServer interface {
	GetAd(context.Context, *GetAdRequest) (*GetAdResponse, error)
	ListAds(context.Context, *ListAdsRequest) (*ListAdsResponse, error)
	SearchAds(context.Context, *SearchAdsRequest) (*SearchAdsResponse, error)
	CreateAd(context.Context, *CreateAdRequest) (*Ad, error)
	UpdateAd(context.Context, *UpdateAdRequest) (*Ad, error)
	DeleteAd(context.Context, *DeleteAdRequest) (*empty.Empty, error)
	ListPhotos(context.Context, *ListPhotosRequest) (*ListPhotosResponse, error)
	ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error)
	mustEmbedUnimplementedAdServiceServer()
}

// UnimplementedAdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdServiceServer struct {
}

func (UnimplementedAdServiceServer) GetAd(context.Context, *GetAdRequest) (*GetAdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAd not implemented")
}
func (UnimplementedAdServiceServer) ListAds(context.Context, *ListAdsRequest) (*ListAdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAds not implemented")
}
func (UnimplementedAdServiceServer) SearchAds(context.Context, *SearchAdsRequest) (*SearchAdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAds not implemented")
}
func (UnimplementedAdServiceServer) CreateAd(context.Context, *CreateAdRequest) (*Ad, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAd not implemented")
}
func (UnimplementedAdServiceServer) UpdateAd(context.Context, *UpdateAdRequest) (*Ad, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAd not implemented")
}
func (UnimplementedAdServiceServer) DeleteAd(context.Context, *DeleteAdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAd not implemented")
}
func (UnimplementedAdServiceServer) ListPhotos(context.Context, *ListPhotosRequest) (*ListPhotosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPhotos not implemented")
}
func (UnimplementedAdServiceServer) ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTags not implemented")
}
func (UnimplementedAdServiceServer) mustEmbedUnimplementedAdServiceServer() {}

// UnsafeAdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdServiceServer will
// result in compilation errors.
type UnsafeAdServiceServer interface {
	mustEmbedUnimplementedAdServiceServer()
}

func RegisterAdServiceServer(s grpc.ServiceRegistrar, srv AdServiceServer) {
	s.RegisterService(&AdService_ServiceDesc, srv)
}

func _AdService_GetAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).GetAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc_ads.AdService/GetAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).GetAd(ctx, req.(*GetAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_ListAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).ListAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc_ads.AdService/ListAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).ListAds(ctx, req.(*ListAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_SearchAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).SearchAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc_ads.AdService/SearchAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).SearchAds(ctx, req.(*SearchAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_CreateAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).CreateAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc_ads.AdService/CreateAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).CreateAd(ctx, req.(*CreateAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_UpdateAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).UpdateAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc_ads.AdService/UpdateAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).UpdateAd(ctx, req.(*UpdateAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_DeleteAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).DeleteAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc_ads.AdService/DeleteAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).DeleteAd(ctx, req.(*DeleteAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_ListPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPhotosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).ListPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc_ads.AdService/ListPhotos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).ListPhotos(ctx, req.(*ListPhotosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/svc_ads.AdService/ListTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).ListTags(ctx, req.(*ListTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdService_ServiceDesc is the grpc.ServiceDesc for AdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "svc_ads.AdService",
	HandlerType: (*AdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAd",
			Handler:    _AdService_GetAd_Handler,
		},
		{
			MethodName: "ListAds",
			Handler:    _AdService_ListAds_Handler,
		},
		{
			MethodName: "SearchAds",
			Handler:    _AdService_SearchAds_Handler,
		},
		{
			MethodName: "CreateAd",
			Handler:    _AdService_CreateAd_Handler,
		},
		{
			MethodName: "UpdateAd",
			Handler:    _AdService_UpdateAd_Handler,
		},
		{
			MethodName: "DeleteAd",
			Handler:    _AdService_DeleteAd_Handler,
		},
		{
			MethodName: "ListPhotos",
			Handler:    _AdService_ListPhotos_Handler,
		},
		{
			MethodName: "ListTags",
			Handler:    _AdService_ListTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/svc_ads.proto",
}
