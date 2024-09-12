// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: server.proto

package pb

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// SchedulerYARPCClient is the YARPC client-side interface for the Scheduler service.
type SchedulerYARPCClient interface {
	Add(context.Context, *SchedulerEntryRequest, ...yarpc.CallOption) (*SchedulerResponse, error)
	Remove(context.Context, *SchedulerEntryRequest, ...yarpc.CallOption) (*SchedulerResponse, error)
}

func newSchedulerYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) SchedulerYARPCClient {
	return &_SchedulerYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "server.Scheduler",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewSchedulerYARPCClient builds a new YARPC client for the Scheduler service.
func NewSchedulerYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) SchedulerYARPCClient {
	return newSchedulerYARPCClient(clientConfig, nil, options...)
}

// SchedulerYARPCServer is the YARPC server-side interface for the Scheduler service.
type SchedulerYARPCServer interface {
	Add(context.Context, *SchedulerEntryRequest) (*SchedulerResponse, error)
	Remove(context.Context, *SchedulerEntryRequest) (*SchedulerResponse, error)
}

type buildSchedulerYARPCProceduresParams struct {
	Server      SchedulerYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildSchedulerYARPCProcedures(params buildSchedulerYARPCProceduresParams) []transport.Procedure {
	handler := &_SchedulerYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "server.Scheduler",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Add",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Add,
							NewRequest:  newSchedulerServiceAddYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
				{
					MethodName: "Remove",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Remove,
							NewRequest:  newSchedulerServiceRemoveYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildSchedulerYARPCProcedures prepares an implementation of the Scheduler service for YARPC registration.
func BuildSchedulerYARPCProcedures(server SchedulerYARPCServer) []transport.Procedure {
	return buildSchedulerYARPCProcedures(buildSchedulerYARPCProceduresParams{Server: server})
}

// FxSchedulerYARPCClientParams defines the input
// for NewFxSchedulerYARPCClient. It provides the
// paramaters to get a SchedulerYARPCClient in an
// Fx application.
type FxSchedulerYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxSchedulerYARPCClientResult defines the output
// of NewFxSchedulerYARPCClient. It provides a
// SchedulerYARPCClient to an Fx application.
type FxSchedulerYARPCClientResult struct {
	fx.Out

	Client SchedulerYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxSchedulerYARPCClient provides a SchedulerYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    pb.NewFxSchedulerYARPCClient("service-name"),
//    ...
//  )
func NewFxSchedulerYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxSchedulerYARPCClientParams) FxSchedulerYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxSchedulerYARPCClientResult{
			Client: newSchedulerYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxSchedulerYARPCProceduresParams defines the input
// for NewFxSchedulerYARPCProcedures. It provides the
// paramaters to get SchedulerYARPCServer procedures in an
// Fx application.
type FxSchedulerYARPCProceduresParams struct {
	fx.In

	Server      SchedulerYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxSchedulerYARPCProceduresResult defines the output
// of NewFxSchedulerYARPCProcedures. It provides
// SchedulerYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxSchedulerYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxSchedulerYARPCProcedures provides SchedulerYARPCServer procedures to an Fx application.
// It expects a SchedulerYARPCServer to be present in the container.
//
//  fx.Provide(
//    pb.NewFxSchedulerYARPCProcedures(),
//    ...
//  )
func NewFxSchedulerYARPCProcedures() interface{} {
	return func(params FxSchedulerYARPCProceduresParams) FxSchedulerYARPCProceduresResult {
		return FxSchedulerYARPCProceduresResult{
			Procedures: buildSchedulerYARPCProcedures(buildSchedulerYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "server.Scheduler",
				FileDescriptors: yarpcFileDescriptorClosuread098daeda4239f7,
			},
		}
	}
}

type _SchedulerYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_SchedulerYARPCCaller) Add(ctx context.Context, request *SchedulerEntryRequest, options ...yarpc.CallOption) (*SchedulerResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Add", request, newSchedulerServiceAddYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*SchedulerResponse)
	if !ok {
		return nil, protobuf.CastError(emptySchedulerServiceAddYARPCResponse, responseMessage)
	}
	return response, err
}

func (c *_SchedulerYARPCCaller) Remove(ctx context.Context, request *SchedulerEntryRequest, options ...yarpc.CallOption) (*SchedulerResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Remove", request, newSchedulerServiceRemoveYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*SchedulerResponse)
	if !ok {
		return nil, protobuf.CastError(emptySchedulerServiceRemoveYARPCResponse, responseMessage)
	}
	return response, err
}

type _SchedulerYARPCHandler struct {
	server SchedulerYARPCServer
}

func (h *_SchedulerYARPCHandler) Add(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *SchedulerEntryRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*SchedulerEntryRequest)
		if !ok {
			return nil, protobuf.CastError(emptySchedulerServiceAddYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Add(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func (h *_SchedulerYARPCHandler) Remove(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *SchedulerEntryRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*SchedulerEntryRequest)
		if !ok {
			return nil, protobuf.CastError(emptySchedulerServiceRemoveYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Remove(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newSchedulerServiceAddYARPCRequest() proto.Message {
	return &SchedulerEntryRequest{}
}

func newSchedulerServiceAddYARPCResponse() proto.Message {
	return &SchedulerResponse{}
}

func newSchedulerServiceRemoveYARPCRequest() proto.Message {
	return &SchedulerEntryRequest{}
}

func newSchedulerServiceRemoveYARPCResponse() proto.Message {
	return &SchedulerResponse{}
}

var (
	emptySchedulerServiceAddYARPCRequest     = &SchedulerEntryRequest{}
	emptySchedulerServiceAddYARPCResponse    = &SchedulerResponse{}
	emptySchedulerServiceRemoveYARPCRequest  = &SchedulerEntryRequest{}
	emptySchedulerServiceRemoveYARPCResponse = &SchedulerResponse{}
)

var yarpcFileDescriptorClosuread098daeda4239f7 = [][]byte{
	// server.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
		0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x8a, 0xb9, 0x44,
		0x83, 0x93, 0x33, 0x52, 0x53, 0x4a, 0x73, 0x52, 0x8b, 0x5c, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0x52,
		0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18,
		0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x09, 0x2e, 0xf6, 0x82, 0xc4, 0xca, 0x9c, 0xfc, 0xc4, 0x14,
		0x09, 0x26, 0xb0, 0x28, 0x8c, 0x2b, 0xa4, 0xc7, 0x25, 0x54, 0x0c, 0x35, 0x24, 0x25, 0x24, 0x33,
		0x37, 0xd5, 0xb5, 0x20, 0x3f, 0x39, 0x43, 0x82, 0x59, 0x81, 0x51, 0x83, 0x39, 0x08, 0x8b, 0x8c,
		0x92, 0x2e, 0x97, 0x20, 0xdc, 0xd2, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x90, 0xf1,
		0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x50, 0x4b, 0x61, 0x5c, 0xa3, 0x29, 0x8c, 0x5c, 0x9c,
		0x70, 0xf5, 0x42, 0x8e, 0x5c, 0xcc, 0x8e, 0x29, 0x29, 0x42, 0xb2, 0x7a, 0x50, 0xff, 0x60, 0x75,
		0xbe, 0x94, 0x24, 0x86, 0x34, 0xcc, 0x22, 0x25, 0x06, 0x21, 0x17, 0x2e, 0xb6, 0xa0, 0xd4, 0xdc,
		0xfc, 0xb2, 0x54, 0x4a, 0x4c, 0x71, 0xe2, 0x8e, 0xe2, 0x84, 0xc8, 0xea, 0x17, 0x24, 0x25, 0xb1,
		0x81, 0x83, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x17, 0xfb, 0x0e, 0xf7, 0x66, 0x01, 0x00,
		0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) SchedulerYARPCClient {
			return NewSchedulerYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
