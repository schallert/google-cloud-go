// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package cx

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newTransitionRouteGroupsClientHook clientHook

// TransitionRouteGroupsCallOptions contains the retry settings for each method of TransitionRouteGroupsClient.
type TransitionRouteGroupsCallOptions struct {
	ListTransitionRouteGroups  []gax.CallOption
	GetTransitionRouteGroup    []gax.CallOption
	CreateTransitionRouteGroup []gax.CallOption
	UpdateTransitionRouteGroup []gax.CallOption
	DeleteTransitionRouteGroup []gax.CallOption
	GetLocation                []gax.CallOption
	ListLocations              []gax.CallOption
	CancelOperation            []gax.CallOption
	GetOperation               []gax.CallOption
	ListOperations             []gax.CallOption
}

func defaultTransitionRouteGroupsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTransitionRouteGroupsCallOptions() *TransitionRouteGroupsCallOptions {
	return &TransitionRouteGroupsCallOptions{
		ListTransitionRouteGroups: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetTransitionRouteGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateTransitionRouteGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateTransitionRouteGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteTransitionRouteGroup: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetLocation:     []gax.CallOption{},
		ListLocations:   []gax.CallOption{},
		CancelOperation: []gax.CallOption{},
		GetOperation:    []gax.CallOption{},
		ListOperations:  []gax.CallOption{},
	}
}

// internalTransitionRouteGroupsClient is an interface that defines the methods availaible from Dialogflow API.
type internalTransitionRouteGroupsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListTransitionRouteGroups(context.Context, *cxpb.ListTransitionRouteGroupsRequest, ...gax.CallOption) *TransitionRouteGroupIterator
	GetTransitionRouteGroup(context.Context, *cxpb.GetTransitionRouteGroupRequest, ...gax.CallOption) (*cxpb.TransitionRouteGroup, error)
	CreateTransitionRouteGroup(context.Context, *cxpb.CreateTransitionRouteGroupRequest, ...gax.CallOption) (*cxpb.TransitionRouteGroup, error)
	UpdateTransitionRouteGroup(context.Context, *cxpb.UpdateTransitionRouteGroupRequest, ...gax.CallOption) (*cxpb.TransitionRouteGroup, error)
	DeleteTransitionRouteGroup(context.Context, *cxpb.DeleteTransitionRouteGroupRequest, ...gax.CallOption) error
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// TransitionRouteGroupsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing TransitionRouteGroups.
type TransitionRouteGroupsClient struct {
	// The internal transport-dependent client.
	internalClient internalTransitionRouteGroupsClient

	// The call options for this service.
	CallOptions *TransitionRouteGroupsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TransitionRouteGroupsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TransitionRouteGroupsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *TransitionRouteGroupsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListTransitionRouteGroups returns the list of all transition route groups in the specified flow.
func (c *TransitionRouteGroupsClient) ListTransitionRouteGroups(ctx context.Context, req *cxpb.ListTransitionRouteGroupsRequest, opts ...gax.CallOption) *TransitionRouteGroupIterator {
	return c.internalClient.ListTransitionRouteGroups(ctx, req, opts...)
}

// GetTransitionRouteGroup retrieves the specified TransitionRouteGroup.
func (c *TransitionRouteGroupsClient) GetTransitionRouteGroup(ctx context.Context, req *cxpb.GetTransitionRouteGroupRequest, opts ...gax.CallOption) (*cxpb.TransitionRouteGroup, error) {
	return c.internalClient.GetTransitionRouteGroup(ctx, req, opts...)
}

// CreateTransitionRouteGroup creates an TransitionRouteGroup in the specified flow.
//
// Note: You should always train a flow prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *TransitionRouteGroupsClient) CreateTransitionRouteGroup(ctx context.Context, req *cxpb.CreateTransitionRouteGroupRequest, opts ...gax.CallOption) (*cxpb.TransitionRouteGroup, error) {
	return c.internalClient.CreateTransitionRouteGroup(ctx, req, opts...)
}

// UpdateTransitionRouteGroup updates the specified TransitionRouteGroup.
//
// Note: You should always train a flow prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *TransitionRouteGroupsClient) UpdateTransitionRouteGroup(ctx context.Context, req *cxpb.UpdateTransitionRouteGroupRequest, opts ...gax.CallOption) (*cxpb.TransitionRouteGroup, error) {
	return c.internalClient.UpdateTransitionRouteGroup(ctx, req, opts...)
}

// DeleteTransitionRouteGroup deletes the specified TransitionRouteGroup.
//
// Note: You should always train a flow prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/cx/docs/concept/training).
func (c *TransitionRouteGroupsClient) DeleteTransitionRouteGroup(ctx context.Context, req *cxpb.DeleteTransitionRouteGroupRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteTransitionRouteGroup(ctx, req, opts...)
}

// GetLocation gets information about a location.
func (c *TransitionRouteGroupsClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *TransitionRouteGroupsClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *TransitionRouteGroupsClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *TransitionRouteGroupsClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *TransitionRouteGroupsClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// transitionRouteGroupsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type transitionRouteGroupsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing TransitionRouteGroupsClient
	CallOptions **TransitionRouteGroupsCallOptions

	// The gRPC API client.
	transitionRouteGroupsClient cxpb.TransitionRouteGroupsClient

	operationsClient longrunningpb.OperationsClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTransitionRouteGroupsClient creates a new transition route groups client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing TransitionRouteGroups.
func NewTransitionRouteGroupsClient(ctx context.Context, opts ...option.ClientOption) (*TransitionRouteGroupsClient, error) {
	clientOpts := defaultTransitionRouteGroupsGRPCClientOptions()
	if newTransitionRouteGroupsClientHook != nil {
		hookOpts, err := newTransitionRouteGroupsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TransitionRouteGroupsClient{CallOptions: defaultTransitionRouteGroupsCallOptions()}

	c := &transitionRouteGroupsGRPCClient{
		connPool:                    connPool,
		disableDeadlines:            disableDeadlines,
		transitionRouteGroupsClient: cxpb.NewTransitionRouteGroupsClient(connPool),
		CallOptions:                 &client.CallOptions,
		operationsClient:            longrunningpb.NewOperationsClient(connPool),
		locationsClient:             locationpb.NewLocationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *transitionRouteGroupsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *transitionRouteGroupsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *transitionRouteGroupsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *transitionRouteGroupsGRPCClient) ListTransitionRouteGroups(ctx context.Context, req *cxpb.ListTransitionRouteGroupsRequest, opts ...gax.CallOption) *TransitionRouteGroupIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListTransitionRouteGroups[0:len((*c.CallOptions).ListTransitionRouteGroups):len((*c.CallOptions).ListTransitionRouteGroups)], opts...)
	it := &TransitionRouteGroupIterator{}
	req = proto.Clone(req).(*cxpb.ListTransitionRouteGroupsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.TransitionRouteGroup, string, error) {
		resp := &cxpb.ListTransitionRouteGroupsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.transitionRouteGroupsClient.ListTransitionRouteGroups(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTransitionRouteGroups(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *transitionRouteGroupsGRPCClient) GetTransitionRouteGroup(ctx context.Context, req *cxpb.GetTransitionRouteGroupRequest, opts ...gax.CallOption) (*cxpb.TransitionRouteGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetTransitionRouteGroup[0:len((*c.CallOptions).GetTransitionRouteGroup):len((*c.CallOptions).GetTransitionRouteGroup)], opts...)
	var resp *cxpb.TransitionRouteGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.transitionRouteGroupsClient.GetTransitionRouteGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *transitionRouteGroupsGRPCClient) CreateTransitionRouteGroup(ctx context.Context, req *cxpb.CreateTransitionRouteGroupRequest, opts ...gax.CallOption) (*cxpb.TransitionRouteGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateTransitionRouteGroup[0:len((*c.CallOptions).CreateTransitionRouteGroup):len((*c.CallOptions).CreateTransitionRouteGroup)], opts...)
	var resp *cxpb.TransitionRouteGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.transitionRouteGroupsClient.CreateTransitionRouteGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *transitionRouteGroupsGRPCClient) UpdateTransitionRouteGroup(ctx context.Context, req *cxpb.UpdateTransitionRouteGroupRequest, opts ...gax.CallOption) (*cxpb.TransitionRouteGroup, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "transition_route_group.name", url.QueryEscape(req.GetTransitionRouteGroup().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateTransitionRouteGroup[0:len((*c.CallOptions).UpdateTransitionRouteGroup):len((*c.CallOptions).UpdateTransitionRouteGroup)], opts...)
	var resp *cxpb.TransitionRouteGroup
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.transitionRouteGroupsClient.UpdateTransitionRouteGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *transitionRouteGroupsGRPCClient) DeleteTransitionRouteGroup(ctx context.Context, req *cxpb.DeleteTransitionRouteGroupRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteTransitionRouteGroup[0:len((*c.CallOptions).DeleteTransitionRouteGroup):len((*c.CallOptions).DeleteTransitionRouteGroup)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.transitionRouteGroupsClient.DeleteTransitionRouteGroup(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *transitionRouteGroupsGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.locationsClient.GetLocation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *transitionRouteGroupsGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.locationsClient.ListLocations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *transitionRouteGroupsGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *transitionRouteGroupsGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *transitionRouteGroupsGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// TransitionRouteGroupIterator manages a stream of *cxpb.TransitionRouteGroup.
type TransitionRouteGroupIterator struct {
	items    []*cxpb.TransitionRouteGroup
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.TransitionRouteGroup, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TransitionRouteGroupIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TransitionRouteGroupIterator) Next() (*cxpb.TransitionRouteGroup, error) {
	var item *cxpb.TransitionRouteGroup
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TransitionRouteGroupIterator) bufLen() int {
	return len(it.items)
}

func (it *TransitionRouteGroupIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
