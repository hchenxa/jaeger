// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: jaeger.proto

/*
Package protomodel is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package protomodel

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_CollectorServiceV1_PostSpans_0(ctx context.Context, marshaler runtime.Marshaler, client CollectorServiceV1Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Batch
	var metadata runtime.ServerMetadata

	if req.ContentLength > 0 {
		if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	msg, err := client.PostSpans(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_QueryServiceV1_GetTrace_0(ctx context.Context, marshaler runtime.Marshaler, client QueryServiceV1Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetTraceID
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.GetTrace(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterCollectorServiceV1HandlerFromEndpoint is same as RegisterCollectorServiceV1Handler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCollectorServiceV1HandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCollectorServiceV1Handler(ctx, mux, conn)
}

// RegisterCollectorServiceV1Handler registers the http handlers for service CollectorServiceV1 to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCollectorServiceV1Handler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCollectorServiceV1HandlerClient(ctx, mux, NewCollectorServiceV1Client(conn))
}

// RegisterCollectorServiceV1Handler registers the http handlers for service CollectorServiceV1 to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "CollectorServiceV1Client".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CollectorServiceV1Client"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CollectorServiceV1Client" to call the correct interceptors.
func RegisterCollectorServiceV1HandlerClient(ctx context.Context, mux *runtime.ServeMux, client CollectorServiceV1Client) error {

	mux.Handle("POST", pattern_CollectorServiceV1_PostSpans_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CollectorServiceV1_PostSpans_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CollectorServiceV1_PostSpans_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CollectorServiceV1_PostSpans_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "spans"}, ""))
)

var (
	forward_CollectorServiceV1_PostSpans_0 = runtime.ForwardResponseMessage
)

// RegisterQueryServiceV1HandlerFromEndpoint is same as RegisterQueryServiceV1Handler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterQueryServiceV1HandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterQueryServiceV1Handler(ctx, mux, conn)
}

// RegisterQueryServiceV1Handler registers the http handlers for service QueryServiceV1 to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterQueryServiceV1Handler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterQueryServiceV1HandlerClient(ctx, mux, NewQueryServiceV1Client(conn))
}

// RegisterQueryServiceV1Handler registers the http handlers for service QueryServiceV1 to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "QueryServiceV1Client".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "QueryServiceV1Client"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "QueryServiceV1Client" to call the correct interceptors.
func RegisterQueryServiceV1HandlerClient(ctx context.Context, mux *runtime.ServeMux, client QueryServiceV1Client) error {

	mux.Handle("GET", pattern_QueryServiceV1_GetTrace_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_QueryServiceV1_GetTrace_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_QueryServiceV1_GetTrace_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_QueryServiceV1_GetTrace_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"api", "v1", "traces", "id"}, ""))
)

var (
	forward_QueryServiceV1_GetTrace_0 = runtime.ForwardResponseMessage
)