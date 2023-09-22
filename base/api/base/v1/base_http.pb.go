// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.19.4
// source: base/v1/base.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBaseCreateBase = "/api.base.v1.Base/CreateBase"

type BaseHTTPServer interface {
	CreateBase(context.Context, *CreateBaseRequest) (*CreateBaseReply, error)
}

func RegisterBaseHTTPServer(s *http.Server, srv BaseHTTPServer) {
	r := s.Route("/")
	r.POST("hi/{id}/hero/{account}", _Base_CreateBase0_HTTP_Handler(srv))
	r.GET("hi/hello/{name}", _Base_CreateBase1_HTTP_Handler(srv))
}

func _Base_CreateBase0_HTTP_Handler(srv BaseHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateBaseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBaseCreateBase)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBase(ctx, req.(*CreateBaseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateBaseReply)
		return ctx.Result(200, reply)
	}
}

func _Base_CreateBase1_HTTP_Handler(srv BaseHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateBaseRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBaseCreateBase)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBase(ctx, req.(*CreateBaseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateBaseReply)
		return ctx.Result(200, reply)
	}
}

type BaseHTTPClient interface {
	CreateBase(ctx context.Context, req *CreateBaseRequest, opts ...http.CallOption) (rsp *CreateBaseReply, err error)
}

type BaseHTTPClientImpl struct {
	cc *http.Client
}

func NewBaseHTTPClient(client *http.Client) BaseHTTPClient {
	return &BaseHTTPClientImpl{client}
}

func (c *BaseHTTPClientImpl) CreateBase(ctx context.Context, in *CreateBaseRequest, opts ...http.CallOption) (*CreateBaseReply, error) {
	var out CreateBaseReply
	pattern := "hi/hello/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBaseCreateBase))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}