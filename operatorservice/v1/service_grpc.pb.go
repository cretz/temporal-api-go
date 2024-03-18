// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// plugins:
// - protoc-gen-go-grpc
// - protoc
// source: temporal/api/operatorservice/v1/service.proto

package operatorservice

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
	OperatorService_AddSearchAttributes_FullMethodName                = "/temporal.api.operatorservice.v1.OperatorService/AddSearchAttributes"
	OperatorService_RemoveSearchAttributes_FullMethodName             = "/temporal.api.operatorservice.v1.OperatorService/RemoveSearchAttributes"
	OperatorService_ListSearchAttributes_FullMethodName               = "/temporal.api.operatorservice.v1.OperatorService/ListSearchAttributes"
	OperatorService_DeleteNamespace_FullMethodName                    = "/temporal.api.operatorservice.v1.OperatorService/DeleteNamespace"
	OperatorService_AddOrUpdateRemoteCluster_FullMethodName           = "/temporal.api.operatorservice.v1.OperatorService/AddOrUpdateRemoteCluster"
	OperatorService_RemoveRemoteCluster_FullMethodName                = "/temporal.api.operatorservice.v1.OperatorService/RemoveRemoteCluster"
	OperatorService_ListClusters_FullMethodName                       = "/temporal.api.operatorservice.v1.OperatorService/ListClusters"
	OperatorService_GetNexusIncomingService_FullMethodName            = "/temporal.api.operatorservice.v1.OperatorService/GetNexusIncomingService"
	OperatorService_CreateOrUpdateNexusIncomingService_FullMethodName = "/temporal.api.operatorservice.v1.OperatorService/CreateOrUpdateNexusIncomingService"
	OperatorService_DeleteNexusIncomingService_FullMethodName         = "/temporal.api.operatorservice.v1.OperatorService/DeleteNexusIncomingService"
	OperatorService_ListNexusIncomingServices_FullMethodName          = "/temporal.api.operatorservice.v1.OperatorService/ListNexusIncomingServices"
	OperatorService_GetNexusOutgoingService_FullMethodName            = "/temporal.api.operatorservice.v1.OperatorService/GetNexusOutgoingService"
	OperatorService_CreateNexusOutgoingService_FullMethodName         = "/temporal.api.operatorservice.v1.OperatorService/CreateNexusOutgoingService"
	OperatorService_UpdateNexusOutgoingService_FullMethodName         = "/temporal.api.operatorservice.v1.OperatorService/UpdateNexusOutgoingService"
	OperatorService_DeleteNexusOutgoingService_FullMethodName         = "/temporal.api.operatorservice.v1.OperatorService/DeleteNexusOutgoingService"
	OperatorService_ListNexusOutgoingServices_FullMethodName          = "/temporal.api.operatorservice.v1.OperatorService/ListNexusOutgoingServices"
)

// OperatorServiceClient is the client API for OperatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatorServiceClient interface {
	// AddSearchAttributes add custom search attributes.
	//
	// Returns ALREADY_EXISTS status code if a Search Attribute with any of the specified names already exists
	// Returns INTERNAL status code with temporal.api.errordetails.v1.SystemWorkflowFailure in Error Details if registration process fails,
	AddSearchAttributes(ctx context.Context, in *AddSearchAttributesRequest, opts ...grpc.CallOption) (*AddSearchAttributesResponse, error)
	// RemoveSearchAttributes removes custom search attributes.
	//
	// Returns NOT_FOUND status code if a Search Attribute with any of the specified names is not registered
	RemoveSearchAttributes(ctx context.Context, in *RemoveSearchAttributesRequest, opts ...grpc.CallOption) (*RemoveSearchAttributesResponse, error)
	// ListSearchAttributes returns comprehensive information about search attributes.
	ListSearchAttributes(ctx context.Context, in *ListSearchAttributesRequest, opts ...grpc.CallOption) (*ListSearchAttributesResponse, error)
	// DeleteNamespace synchronously deletes a namespace and asynchronously reclaims all namespace resources.
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error)
	// AddOrUpdateRemoteCluster adds or updates remote cluster.
	AddOrUpdateRemoteCluster(ctx context.Context, in *AddOrUpdateRemoteClusterRequest, opts ...grpc.CallOption) (*AddOrUpdateRemoteClusterResponse, error)
	// RemoveRemoteCluster removes remote cluster.
	RemoveRemoteCluster(ctx context.Context, in *RemoveRemoteClusterRequest, opts ...grpc.CallOption) (*RemoveRemoteClusterResponse, error)
	// ListClusters returns information about Temporal clusters.
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Get a registered incoming Nexus service by name. The returned version can be used for optimistic updates.
	GetNexusIncomingService(ctx context.Context, in *GetNexusIncomingServiceRequest, opts ...grpc.CallOption) (*GetNexusIncomingServiceResponse, error)
	// Optimistically create or update a Nexus service based on provided version.
	// To update an existing service, get the current service record via the `GetNexusIncomingService` API, modify it
	// and submit to this API.
	// Set version to 0 to create a new service.
	// Returns the updated service with the updated version, which can be used for subsequent updates.
	CreateOrUpdateNexusIncomingService(ctx context.Context, in *CreateOrUpdateNexusIncomingServiceRequest, opts ...grpc.CallOption) (*CreateOrUpdateNexusIncomingServiceResponse, error)
	// Delete an incoming Nexus service by name.
	DeleteNexusIncomingService(ctx context.Context, in *DeleteNexusIncomingServiceRequest, opts ...grpc.CallOption) (*DeleteNexusIncomingServiceResponse, error)
	// List all Nexus incoming services in the cluster. Use next_page_token in the response for pagination.
	ListNexusIncomingServices(ctx context.Context, in *ListNexusIncomingServicesRequest, opts ...grpc.CallOption) (*ListNexusIncomingServicesResponse, error)
	// Get a registered outgoing Nexus service by namespace and service name. The returned version can be used for
	// optimistic updates.
	GetNexusOutgoingService(ctx context.Context, in *GetNexusOutgoingServiceRequest, opts ...grpc.CallOption) (*GetNexusOutgoingServiceResponse, error)
	// Create a Nexus service. This will fail if a service with the same name already exists in the namespace with a
	// status of ALREADY_EXISTS.
	// Returns the updated service with its initial version. You may use this version for subsequent updates. You don't
	// need to increment the version yourself. The server will increment the version for you after each update.
	CreateNexusOutgoingService(ctx context.Context, in *CreateNexusOutgoingServiceRequest, opts ...grpc.CallOption) (*CreateNexusOutgoingServiceResponse, error)
	// Update an outgoing Nexus service by namespace and service name. The version in the request should match the
	// current version of the service. This will fail with a status of FAILED_PRECONDITION if the version does not match.
	// Returns the updated service with the updated version, which can be used for subsequent updates. You don't need
	// to increment the version yourself. The server will increment the version for you.
	UpdateNexusOutgoingService(ctx context.Context, in *UpdateNexusOutgoingServiceRequest, opts ...grpc.CallOption) (*UpdateNexusOutgoingServiceResponse, error)
	// Delete an outgoing Nexus service by namespace and service name.
	DeleteNexusOutgoingService(ctx context.Context, in *DeleteNexusOutgoingServiceRequest, opts ...grpc.CallOption) (*DeleteNexusOutgoingServiceResponse, error)
	// List all Nexus outgoing services for a namespace, sorted by service name in ascending order. Set page_token in
	// the request to the next_page_token field of the previous response to get the next page of results. An empty
	// next_page_token indicates that there are no more results. During pagination, a newly added service with a name
	// lexicographically earlier than the previous page's last service name may be missed.
	ListNexusOutgoingServices(ctx context.Context, in *ListNexusOutgoingServicesRequest, opts ...grpc.CallOption) (*ListNexusOutgoingServicesResponse, error)
}

type operatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatorServiceClient(cc grpc.ClientConnInterface) OperatorServiceClient {
	return &operatorServiceClient{cc}
}

func (c *operatorServiceClient) AddSearchAttributes(ctx context.Context, in *AddSearchAttributesRequest, opts ...grpc.CallOption) (*AddSearchAttributesResponse, error) {
	out := new(AddSearchAttributesResponse)
	err := c.cc.Invoke(ctx, OperatorService_AddSearchAttributes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) RemoveSearchAttributes(ctx context.Context, in *RemoveSearchAttributesRequest, opts ...grpc.CallOption) (*RemoveSearchAttributesResponse, error) {
	out := new(RemoveSearchAttributesResponse)
	err := c.cc.Invoke(ctx, OperatorService_RemoveSearchAttributes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListSearchAttributes(ctx context.Context, in *ListSearchAttributesRequest, opts ...grpc.CallOption) (*ListSearchAttributesResponse, error) {
	out := new(ListSearchAttributesResponse)
	err := c.cc.Invoke(ctx, OperatorService_ListSearchAttributes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error) {
	out := new(DeleteNamespaceResponse)
	err := c.cc.Invoke(ctx, OperatorService_DeleteNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) AddOrUpdateRemoteCluster(ctx context.Context, in *AddOrUpdateRemoteClusterRequest, opts ...grpc.CallOption) (*AddOrUpdateRemoteClusterResponse, error) {
	out := new(AddOrUpdateRemoteClusterResponse)
	err := c.cc.Invoke(ctx, OperatorService_AddOrUpdateRemoteCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) RemoveRemoteCluster(ctx context.Context, in *RemoveRemoteClusterRequest, opts ...grpc.CallOption) (*RemoveRemoteClusterResponse, error) {
	out := new(RemoveRemoteClusterResponse)
	err := c.cc.Invoke(ctx, OperatorService_RemoveRemoteCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, OperatorService_ListClusters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) GetNexusIncomingService(ctx context.Context, in *GetNexusIncomingServiceRequest, opts ...grpc.CallOption) (*GetNexusIncomingServiceResponse, error) {
	out := new(GetNexusIncomingServiceResponse)
	err := c.cc.Invoke(ctx, OperatorService_GetNexusIncomingService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) CreateOrUpdateNexusIncomingService(ctx context.Context, in *CreateOrUpdateNexusIncomingServiceRequest, opts ...grpc.CallOption) (*CreateOrUpdateNexusIncomingServiceResponse, error) {
	out := new(CreateOrUpdateNexusIncomingServiceResponse)
	err := c.cc.Invoke(ctx, OperatorService_CreateOrUpdateNexusIncomingService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) DeleteNexusIncomingService(ctx context.Context, in *DeleteNexusIncomingServiceRequest, opts ...grpc.CallOption) (*DeleteNexusIncomingServiceResponse, error) {
	out := new(DeleteNexusIncomingServiceResponse)
	err := c.cc.Invoke(ctx, OperatorService_DeleteNexusIncomingService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListNexusIncomingServices(ctx context.Context, in *ListNexusIncomingServicesRequest, opts ...grpc.CallOption) (*ListNexusIncomingServicesResponse, error) {
	out := new(ListNexusIncomingServicesResponse)
	err := c.cc.Invoke(ctx, OperatorService_ListNexusIncomingServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) GetNexusOutgoingService(ctx context.Context, in *GetNexusOutgoingServiceRequest, opts ...grpc.CallOption) (*GetNexusOutgoingServiceResponse, error) {
	out := new(GetNexusOutgoingServiceResponse)
	err := c.cc.Invoke(ctx, OperatorService_GetNexusOutgoingService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) CreateNexusOutgoingService(ctx context.Context, in *CreateNexusOutgoingServiceRequest, opts ...grpc.CallOption) (*CreateNexusOutgoingServiceResponse, error) {
	out := new(CreateNexusOutgoingServiceResponse)
	err := c.cc.Invoke(ctx, OperatorService_CreateNexusOutgoingService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) UpdateNexusOutgoingService(ctx context.Context, in *UpdateNexusOutgoingServiceRequest, opts ...grpc.CallOption) (*UpdateNexusOutgoingServiceResponse, error) {
	out := new(UpdateNexusOutgoingServiceResponse)
	err := c.cc.Invoke(ctx, OperatorService_UpdateNexusOutgoingService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) DeleteNexusOutgoingService(ctx context.Context, in *DeleteNexusOutgoingServiceRequest, opts ...grpc.CallOption) (*DeleteNexusOutgoingServiceResponse, error) {
	out := new(DeleteNexusOutgoingServiceResponse)
	err := c.cc.Invoke(ctx, OperatorService_DeleteNexusOutgoingService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListNexusOutgoingServices(ctx context.Context, in *ListNexusOutgoingServicesRequest, opts ...grpc.CallOption) (*ListNexusOutgoingServicesResponse, error) {
	out := new(ListNexusOutgoingServicesResponse)
	err := c.cc.Invoke(ctx, OperatorService_ListNexusOutgoingServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServiceServer is the server API for OperatorService service.
// All implementations must embed UnimplementedOperatorServiceServer
// for forward compatibility
type OperatorServiceServer interface {
	// AddSearchAttributes add custom search attributes.
	//
	// Returns ALREADY_EXISTS status code if a Search Attribute with any of the specified names already exists
	// Returns INTERNAL status code with temporal.api.errordetails.v1.SystemWorkflowFailure in Error Details if registration process fails,
	AddSearchAttributes(context.Context, *AddSearchAttributesRequest) (*AddSearchAttributesResponse, error)
	// RemoveSearchAttributes removes custom search attributes.
	//
	// Returns NOT_FOUND status code if a Search Attribute with any of the specified names is not registered
	RemoveSearchAttributes(context.Context, *RemoveSearchAttributesRequest) (*RemoveSearchAttributesResponse, error)
	// ListSearchAttributes returns comprehensive information about search attributes.
	ListSearchAttributes(context.Context, *ListSearchAttributesRequest) (*ListSearchAttributesResponse, error)
	// DeleteNamespace synchronously deletes a namespace and asynchronously reclaims all namespace resources.
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	// AddOrUpdateRemoteCluster adds or updates remote cluster.
	AddOrUpdateRemoteCluster(context.Context, *AddOrUpdateRemoteClusterRequest) (*AddOrUpdateRemoteClusterResponse, error)
	// RemoveRemoteCluster removes remote cluster.
	RemoveRemoteCluster(context.Context, *RemoveRemoteClusterRequest) (*RemoveRemoteClusterResponse, error)
	// ListClusters returns information about Temporal clusters.
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Get a registered incoming Nexus service by name. The returned version can be used for optimistic updates.
	GetNexusIncomingService(context.Context, *GetNexusIncomingServiceRequest) (*GetNexusIncomingServiceResponse, error)
	// Optimistically create or update a Nexus service based on provided version.
	// To update an existing service, get the current service record via the `GetNexusIncomingService` API, modify it
	// and submit to this API.
	// Set version to 0 to create a new service.
	// Returns the updated service with the updated version, which can be used for subsequent updates.
	CreateOrUpdateNexusIncomingService(context.Context, *CreateOrUpdateNexusIncomingServiceRequest) (*CreateOrUpdateNexusIncomingServiceResponse, error)
	// Delete an incoming Nexus service by name.
	DeleteNexusIncomingService(context.Context, *DeleteNexusIncomingServiceRequest) (*DeleteNexusIncomingServiceResponse, error)
	// List all Nexus incoming services in the cluster. Use next_page_token in the response for pagination.
	ListNexusIncomingServices(context.Context, *ListNexusIncomingServicesRequest) (*ListNexusIncomingServicesResponse, error)
	// Get a registered outgoing Nexus service by namespace and service name. The returned version can be used for
	// optimistic updates.
	GetNexusOutgoingService(context.Context, *GetNexusOutgoingServiceRequest) (*GetNexusOutgoingServiceResponse, error)
	// Create a Nexus service. This will fail if a service with the same name already exists in the namespace with a
	// status of ALREADY_EXISTS.
	// Returns the updated service with its initial version. You may use this version for subsequent updates. You don't
	// need to increment the version yourself. The server will increment the version for you after each update.
	CreateNexusOutgoingService(context.Context, *CreateNexusOutgoingServiceRequest) (*CreateNexusOutgoingServiceResponse, error)
	// Update an outgoing Nexus service by namespace and service name. The version in the request should match the
	// current version of the service. This will fail with a status of FAILED_PRECONDITION if the version does not match.
	// Returns the updated service with the updated version, which can be used for subsequent updates. You don't need
	// to increment the version yourself. The server will increment the version for you.
	UpdateNexusOutgoingService(context.Context, *UpdateNexusOutgoingServiceRequest) (*UpdateNexusOutgoingServiceResponse, error)
	// Delete an outgoing Nexus service by namespace and service name.
	DeleteNexusOutgoingService(context.Context, *DeleteNexusOutgoingServiceRequest) (*DeleteNexusOutgoingServiceResponse, error)
	// List all Nexus outgoing services for a namespace, sorted by service name in ascending order. Set page_token in
	// the request to the next_page_token field of the previous response to get the next page of results. An empty
	// next_page_token indicates that there are no more results. During pagination, a newly added service with a name
	// lexicographically earlier than the previous page's last service name may be missed.
	ListNexusOutgoingServices(context.Context, *ListNexusOutgoingServicesRequest) (*ListNexusOutgoingServicesResponse, error)
	mustEmbedUnimplementedOperatorServiceServer()
}

// UnimplementedOperatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperatorServiceServer struct {
}

func (UnimplementedOperatorServiceServer) AddSearchAttributes(context.Context, *AddSearchAttributesRequest) (*AddSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSearchAttributes not implemented")
}
func (UnimplementedOperatorServiceServer) RemoveSearchAttributes(context.Context, *RemoveSearchAttributesRequest) (*RemoveSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSearchAttributes not implemented")
}
func (UnimplementedOperatorServiceServer) ListSearchAttributes(context.Context, *ListSearchAttributesRequest) (*ListSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSearchAttributes not implemented")
}
func (UnimplementedOperatorServiceServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedOperatorServiceServer) AddOrUpdateRemoteCluster(context.Context, *AddOrUpdateRemoteClusterRequest) (*AddOrUpdateRemoteClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateRemoteCluster not implemented")
}
func (UnimplementedOperatorServiceServer) RemoveRemoteCluster(context.Context, *RemoveRemoteClusterRequest) (*RemoveRemoteClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRemoteCluster not implemented")
}
func (UnimplementedOperatorServiceServer) ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (UnimplementedOperatorServiceServer) GetNexusIncomingService(context.Context, *GetNexusIncomingServiceRequest) (*GetNexusIncomingServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNexusIncomingService not implemented")
}
func (UnimplementedOperatorServiceServer) CreateOrUpdateNexusIncomingService(context.Context, *CreateOrUpdateNexusIncomingServiceRequest) (*CreateOrUpdateNexusIncomingServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateNexusIncomingService not implemented")
}
func (UnimplementedOperatorServiceServer) DeleteNexusIncomingService(context.Context, *DeleteNexusIncomingServiceRequest) (*DeleteNexusIncomingServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNexusIncomingService not implemented")
}
func (UnimplementedOperatorServiceServer) ListNexusIncomingServices(context.Context, *ListNexusIncomingServicesRequest) (*ListNexusIncomingServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNexusIncomingServices not implemented")
}
func (UnimplementedOperatorServiceServer) GetNexusOutgoingService(context.Context, *GetNexusOutgoingServiceRequest) (*GetNexusOutgoingServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNexusOutgoingService not implemented")
}
func (UnimplementedOperatorServiceServer) CreateNexusOutgoingService(context.Context, *CreateNexusOutgoingServiceRequest) (*CreateNexusOutgoingServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNexusOutgoingService not implemented")
}
func (UnimplementedOperatorServiceServer) UpdateNexusOutgoingService(context.Context, *UpdateNexusOutgoingServiceRequest) (*UpdateNexusOutgoingServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNexusOutgoingService not implemented")
}
func (UnimplementedOperatorServiceServer) DeleteNexusOutgoingService(context.Context, *DeleteNexusOutgoingServiceRequest) (*DeleteNexusOutgoingServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNexusOutgoingService not implemented")
}
func (UnimplementedOperatorServiceServer) ListNexusOutgoingServices(context.Context, *ListNexusOutgoingServicesRequest) (*ListNexusOutgoingServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNexusOutgoingServices not implemented")
}
func (UnimplementedOperatorServiceServer) mustEmbedUnimplementedOperatorServiceServer() {}

// UnsafeOperatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatorServiceServer will
// result in compilation errors.
type UnsafeOperatorServiceServer interface {
	mustEmbedUnimplementedOperatorServiceServer()
}

func RegisterOperatorServiceServer(s grpc.ServiceRegistrar, srv OperatorServiceServer) {
	s.RegisterService(&OperatorService_ServiceDesc, srv)
}

func _OperatorService_AddSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).AddSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_AddSearchAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).AddSearchAttributes(ctx, req.(*AddSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_RemoveSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).RemoveSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_RemoveSearchAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).RemoveSearchAttributes(ctx, req.(*RemoveSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_ListSearchAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListSearchAttributes(ctx, req.(*ListSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_DeleteNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_AddOrUpdateRemoteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateRemoteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).AddOrUpdateRemoteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_AddOrUpdateRemoteCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).AddOrUpdateRemoteCluster(ctx, req.(*AddOrUpdateRemoteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_RemoveRemoteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRemoteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).RemoveRemoteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_RemoveRemoteCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).RemoveRemoteCluster(ctx, req.(*RemoveRemoteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_ListClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_GetNexusIncomingService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNexusIncomingServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).GetNexusIncomingService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_GetNexusIncomingService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).GetNexusIncomingService(ctx, req.(*GetNexusIncomingServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_CreateOrUpdateNexusIncomingService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateNexusIncomingServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).CreateOrUpdateNexusIncomingService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_CreateOrUpdateNexusIncomingService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).CreateOrUpdateNexusIncomingService(ctx, req.(*CreateOrUpdateNexusIncomingServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_DeleteNexusIncomingService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNexusIncomingServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).DeleteNexusIncomingService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_DeleteNexusIncomingService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).DeleteNexusIncomingService(ctx, req.(*DeleteNexusIncomingServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListNexusIncomingServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNexusIncomingServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListNexusIncomingServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_ListNexusIncomingServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListNexusIncomingServices(ctx, req.(*ListNexusIncomingServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_GetNexusOutgoingService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNexusOutgoingServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).GetNexusOutgoingService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_GetNexusOutgoingService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).GetNexusOutgoingService(ctx, req.(*GetNexusOutgoingServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_CreateNexusOutgoingService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNexusOutgoingServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).CreateNexusOutgoingService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_CreateNexusOutgoingService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).CreateNexusOutgoingService(ctx, req.(*CreateNexusOutgoingServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_UpdateNexusOutgoingService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNexusOutgoingServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).UpdateNexusOutgoingService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_UpdateNexusOutgoingService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).UpdateNexusOutgoingService(ctx, req.(*UpdateNexusOutgoingServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_DeleteNexusOutgoingService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNexusOutgoingServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).DeleteNexusOutgoingService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_DeleteNexusOutgoingService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).DeleteNexusOutgoingService(ctx, req.(*DeleteNexusOutgoingServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListNexusOutgoingServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNexusOutgoingServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListNexusOutgoingServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_ListNexusOutgoingServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListNexusOutgoingServices(ctx, req.(*ListNexusOutgoingServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperatorService_ServiceDesc is the grpc.ServiceDesc for OperatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.api.operatorservice.v1.OperatorService",
	HandlerType: (*OperatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSearchAttributes",
			Handler:    _OperatorService_AddSearchAttributes_Handler,
		},
		{
			MethodName: "RemoveSearchAttributes",
			Handler:    _OperatorService_RemoveSearchAttributes_Handler,
		},
		{
			MethodName: "ListSearchAttributes",
			Handler:    _OperatorService_ListSearchAttributes_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _OperatorService_DeleteNamespace_Handler,
		},
		{
			MethodName: "AddOrUpdateRemoteCluster",
			Handler:    _OperatorService_AddOrUpdateRemoteCluster_Handler,
		},
		{
			MethodName: "RemoveRemoteCluster",
			Handler:    _OperatorService_RemoveRemoteCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _OperatorService_ListClusters_Handler,
		},
		{
			MethodName: "GetNexusIncomingService",
			Handler:    _OperatorService_GetNexusIncomingService_Handler,
		},
		{
			MethodName: "CreateOrUpdateNexusIncomingService",
			Handler:    _OperatorService_CreateOrUpdateNexusIncomingService_Handler,
		},
		{
			MethodName: "DeleteNexusIncomingService",
			Handler:    _OperatorService_DeleteNexusIncomingService_Handler,
		},
		{
			MethodName: "ListNexusIncomingServices",
			Handler:    _OperatorService_ListNexusIncomingServices_Handler,
		},
		{
			MethodName: "GetNexusOutgoingService",
			Handler:    _OperatorService_GetNexusOutgoingService_Handler,
		},
		{
			MethodName: "CreateNexusOutgoingService",
			Handler:    _OperatorService_CreateNexusOutgoingService_Handler,
		},
		{
			MethodName: "UpdateNexusOutgoingService",
			Handler:    _OperatorService_UpdateNexusOutgoingService_Handler,
		},
		{
			MethodName: "DeleteNexusOutgoingService",
			Handler:    _OperatorService_DeleteNexusOutgoingService_Handler,
		},
		{
			MethodName: "ListNexusOutgoingServices",
			Handler:    _OperatorService_ListNexusOutgoingServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temporal/api/operatorservice/v1/service.proto",
}
