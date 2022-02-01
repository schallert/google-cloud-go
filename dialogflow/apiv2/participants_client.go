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

package dialogflow

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
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newParticipantsClientHook clientHook

// ParticipantsCallOptions contains the retry settings for each method of ParticipantsClient.
type ParticipantsCallOptions struct {
	CreateParticipant   []gax.CallOption
	GetParticipant      []gax.CallOption
	ListParticipants    []gax.CallOption
	UpdateParticipant   []gax.CallOption
	AnalyzeContent      []gax.CallOption
	SuggestArticles     []gax.CallOption
	SuggestFaqAnswers   []gax.CallOption
	SuggestSmartReplies []gax.CallOption
}

func defaultParticipantsGRPCClientOptions() []option.ClientOption {
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

func defaultParticipantsCallOptions() *ParticipantsCallOptions {
	return &ParticipantsCallOptions{
		CreateParticipant: []gax.CallOption{
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
		GetParticipant: []gax.CallOption{
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
		ListParticipants: []gax.CallOption{
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
		UpdateParticipant: []gax.CallOption{
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
		AnalyzeContent: []gax.CallOption{
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
		SuggestArticles: []gax.CallOption{
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
		SuggestFaqAnswers: []gax.CallOption{
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
		SuggestSmartReplies: []gax.CallOption{
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
	}
}

// internalParticipantsClient is an interface that defines the methods availaible from Dialogflow API.
type internalParticipantsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateParticipant(context.Context, *dialogflowpb.CreateParticipantRequest, ...gax.CallOption) (*dialogflowpb.Participant, error)
	GetParticipant(context.Context, *dialogflowpb.GetParticipantRequest, ...gax.CallOption) (*dialogflowpb.Participant, error)
	ListParticipants(context.Context, *dialogflowpb.ListParticipantsRequest, ...gax.CallOption) *ParticipantIterator
	UpdateParticipant(context.Context, *dialogflowpb.UpdateParticipantRequest, ...gax.CallOption) (*dialogflowpb.Participant, error)
	AnalyzeContent(context.Context, *dialogflowpb.AnalyzeContentRequest, ...gax.CallOption) (*dialogflowpb.AnalyzeContentResponse, error)
	SuggestArticles(context.Context, *dialogflowpb.SuggestArticlesRequest, ...gax.CallOption) (*dialogflowpb.SuggestArticlesResponse, error)
	SuggestFaqAnswers(context.Context, *dialogflowpb.SuggestFaqAnswersRequest, ...gax.CallOption) (*dialogflowpb.SuggestFaqAnswersResponse, error)
	SuggestSmartReplies(context.Context, *dialogflowpb.SuggestSmartRepliesRequest, ...gax.CallOption) (*dialogflowpb.SuggestSmartRepliesResponse, error)
}

// ParticipantsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing Participants.
type ParticipantsClient struct {
	// The internal transport-dependent client.
	internalClient internalParticipantsClient

	// The call options for this service.
	CallOptions *ParticipantsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ParticipantsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ParticipantsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ParticipantsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateParticipant creates a new participant in a conversation.
func (c *ParticipantsClient) CreateParticipant(ctx context.Context, req *dialogflowpb.CreateParticipantRequest, opts ...gax.CallOption) (*dialogflowpb.Participant, error) {
	return c.internalClient.CreateParticipant(ctx, req, opts...)
}

// GetParticipant retrieves a conversation participant.
func (c *ParticipantsClient) GetParticipant(ctx context.Context, req *dialogflowpb.GetParticipantRequest, opts ...gax.CallOption) (*dialogflowpb.Participant, error) {
	return c.internalClient.GetParticipant(ctx, req, opts...)
}

// ListParticipants returns the list of all participants in the specified conversation.
func (c *ParticipantsClient) ListParticipants(ctx context.Context, req *dialogflowpb.ListParticipantsRequest, opts ...gax.CallOption) *ParticipantIterator {
	return c.internalClient.ListParticipants(ctx, req, opts...)
}

// UpdateParticipant updates the specified participant.
func (c *ParticipantsClient) UpdateParticipant(ctx context.Context, req *dialogflowpb.UpdateParticipantRequest, opts ...gax.CallOption) (*dialogflowpb.Participant, error) {
	return c.internalClient.UpdateParticipant(ctx, req, opts...)
}

// AnalyzeContent adds a text (chat, for example), or audio (phone recording, for example)
// message from a participant into the conversation.
//
// Note: Always use agent versions for production traffic
// sent to virtual agents. See Versions and
// environments (at https://cloud.google.com/dialogflow/es/docs/agents-versions).
func (c *ParticipantsClient) AnalyzeContent(ctx context.Context, req *dialogflowpb.AnalyzeContentRequest, opts ...gax.CallOption) (*dialogflowpb.AnalyzeContentResponse, error) {
	return c.internalClient.AnalyzeContent(ctx, req, opts...)
}

// SuggestArticles gets suggested articles for a participant based on specific historical
// messages.
func (c *ParticipantsClient) SuggestArticles(ctx context.Context, req *dialogflowpb.SuggestArticlesRequest, opts ...gax.CallOption) (*dialogflowpb.SuggestArticlesResponse, error) {
	return c.internalClient.SuggestArticles(ctx, req, opts...)
}

// SuggestFaqAnswers gets suggested faq answers for a participant based on specific historical
// messages.
func (c *ParticipantsClient) SuggestFaqAnswers(ctx context.Context, req *dialogflowpb.SuggestFaqAnswersRequest, opts ...gax.CallOption) (*dialogflowpb.SuggestFaqAnswersResponse, error) {
	return c.internalClient.SuggestFaqAnswers(ctx, req, opts...)
}

// SuggestSmartReplies gets smart replies for a participant based on specific historical
// messages.
func (c *ParticipantsClient) SuggestSmartReplies(ctx context.Context, req *dialogflowpb.SuggestSmartRepliesRequest, opts ...gax.CallOption) (*dialogflowpb.SuggestSmartRepliesResponse, error) {
	return c.internalClient.SuggestSmartReplies(ctx, req, opts...)
}

// participantsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type participantsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ParticipantsClient
	CallOptions **ParticipantsCallOptions

	// The gRPC API client.
	participantsClient dialogflowpb.ParticipantsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewParticipantsClient creates a new participants client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing Participants.
func NewParticipantsClient(ctx context.Context, opts ...option.ClientOption) (*ParticipantsClient, error) {
	clientOpts := defaultParticipantsGRPCClientOptions()
	if newParticipantsClientHook != nil {
		hookOpts, err := newParticipantsClientHook(ctx, clientHookParams{})
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
	client := ParticipantsClient{CallOptions: defaultParticipantsCallOptions()}

	c := &participantsGRPCClient{
		connPool:           connPool,
		disableDeadlines:   disableDeadlines,
		participantsClient: dialogflowpb.NewParticipantsClient(connPool),
		CallOptions:        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *participantsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *participantsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *participantsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *participantsGRPCClient) CreateParticipant(ctx context.Context, req *dialogflowpb.CreateParticipantRequest, opts ...gax.CallOption) (*dialogflowpb.Participant, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateParticipant[0:len((*c.CallOptions).CreateParticipant):len((*c.CallOptions).CreateParticipant)], opts...)
	var resp *dialogflowpb.Participant
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.participantsClient.CreateParticipant(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *participantsGRPCClient) GetParticipant(ctx context.Context, req *dialogflowpb.GetParticipantRequest, opts ...gax.CallOption) (*dialogflowpb.Participant, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetParticipant[0:len((*c.CallOptions).GetParticipant):len((*c.CallOptions).GetParticipant)], opts...)
	var resp *dialogflowpb.Participant
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.participantsClient.GetParticipant(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *participantsGRPCClient) ListParticipants(ctx context.Context, req *dialogflowpb.ListParticipantsRequest, opts ...gax.CallOption) *ParticipantIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListParticipants[0:len((*c.CallOptions).ListParticipants):len((*c.CallOptions).ListParticipants)], opts...)
	it := &ParticipantIterator{}
	req = proto.Clone(req).(*dialogflowpb.ListParticipantsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dialogflowpb.Participant, string, error) {
		resp := &dialogflowpb.ListParticipantsResponse{}
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
			resp, err = c.participantsClient.ListParticipants(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetParticipants(), resp.GetNextPageToken(), nil
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

func (c *participantsGRPCClient) UpdateParticipant(ctx context.Context, req *dialogflowpb.UpdateParticipantRequest, opts ...gax.CallOption) (*dialogflowpb.Participant, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "participant.name", url.QueryEscape(req.GetParticipant().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateParticipant[0:len((*c.CallOptions).UpdateParticipant):len((*c.CallOptions).UpdateParticipant)], opts...)
	var resp *dialogflowpb.Participant
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.participantsClient.UpdateParticipant(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *participantsGRPCClient) AnalyzeContent(ctx context.Context, req *dialogflowpb.AnalyzeContentRequest, opts ...gax.CallOption) (*dialogflowpb.AnalyzeContentResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 220000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "participant", url.QueryEscape(req.GetParticipant())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AnalyzeContent[0:len((*c.CallOptions).AnalyzeContent):len((*c.CallOptions).AnalyzeContent)], opts...)
	var resp *dialogflowpb.AnalyzeContentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.participantsClient.AnalyzeContent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *participantsGRPCClient) SuggestArticles(ctx context.Context, req *dialogflowpb.SuggestArticlesRequest, opts ...gax.CallOption) (*dialogflowpb.SuggestArticlesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SuggestArticles[0:len((*c.CallOptions).SuggestArticles):len((*c.CallOptions).SuggestArticles)], opts...)
	var resp *dialogflowpb.SuggestArticlesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.participantsClient.SuggestArticles(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *participantsGRPCClient) SuggestFaqAnswers(ctx context.Context, req *dialogflowpb.SuggestFaqAnswersRequest, opts ...gax.CallOption) (*dialogflowpb.SuggestFaqAnswersResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SuggestFaqAnswers[0:len((*c.CallOptions).SuggestFaqAnswers):len((*c.CallOptions).SuggestFaqAnswers)], opts...)
	var resp *dialogflowpb.SuggestFaqAnswersResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.participantsClient.SuggestFaqAnswers(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *participantsGRPCClient) SuggestSmartReplies(ctx context.Context, req *dialogflowpb.SuggestSmartRepliesRequest, opts ...gax.CallOption) (*dialogflowpb.SuggestSmartRepliesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SuggestSmartReplies[0:len((*c.CallOptions).SuggestSmartReplies):len((*c.CallOptions).SuggestSmartReplies)], opts...)
	var resp *dialogflowpb.SuggestSmartRepliesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.participantsClient.SuggestSmartReplies(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ParticipantIterator manages a stream of *dialogflowpb.Participant.
type ParticipantIterator struct {
	items    []*dialogflowpb.Participant
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
	InternalFetch func(pageSize int, pageToken string) (results []*dialogflowpb.Participant, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ParticipantIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ParticipantIterator) Next() (*dialogflowpb.Participant, error) {
	var item *dialogflowpb.Participant
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ParticipantIterator) bufLen() int {
	return len(it.items)
}

func (it *ParticipantIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
