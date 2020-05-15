// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/blogs.proto

package blogs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Blogs service

type BlogsService interface {
	Fetch(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *Article, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Article, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *Article, opts ...client.CallOption) (*Response, error)
	FetchTopic(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	CreateTopic(ctx context.Context, in *Topic, opts ...client.CallOption) (*Response, error)
}

type blogsService struct {
	c    client.Client
	name string
}

func NewBlogsService(name string, c client.Client) BlogsService {
	return &blogsService{
		c:    c,
		name: name,
	}
}

func (c *blogsService) Fetch(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Blogs.Fetch", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogsService) Create(ctx context.Context, in *Article, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Blogs.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogsService) Delete(ctx context.Context, in *Article, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Blogs.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogsService) Update(ctx context.Context, in *Article, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Blogs.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogsService) FetchTopic(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Blogs.FetchTopic", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogsService) CreateTopic(ctx context.Context, in *Topic, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Blogs.CreateTopic", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Blogs service

type BlogsHandler interface {
	Fetch(context.Context, *ListRequest, *ListResponse) error
	Create(context.Context, *Article, *Response) error
	Delete(context.Context, *Article, *Response) error
	Update(context.Context, *Article, *Response) error
	FetchTopic(context.Context, *ListRequest, *ListResponse) error
	CreateTopic(context.Context, *Topic, *Response) error
}

func RegisterBlogsHandler(s server.Server, hdlr BlogsHandler, opts ...server.HandlerOption) error {
	type blogs interface {
		Fetch(ctx context.Context, in *ListRequest, out *ListResponse) error
		Create(ctx context.Context, in *Article, out *Response) error
		Delete(ctx context.Context, in *Article, out *Response) error
		Update(ctx context.Context, in *Article, out *Response) error
		FetchTopic(ctx context.Context, in *ListRequest, out *ListResponse) error
		CreateTopic(ctx context.Context, in *Topic, out *Response) error
	}
	type Blogs struct {
		blogs
	}
	h := &blogsHandler{hdlr}
	return s.Handle(s.NewHandler(&Blogs{h}, opts...))
}

type blogsHandler struct {
	BlogsHandler
}

func (h *blogsHandler) Fetch(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.BlogsHandler.Fetch(ctx, in, out)
}

func (h *blogsHandler) Create(ctx context.Context, in *Article, out *Response) error {
	return h.BlogsHandler.Create(ctx, in, out)
}

func (h *blogsHandler) Delete(ctx context.Context, in *Article, out *Response) error {
	return h.BlogsHandler.Delete(ctx, in, out)
}

func (h *blogsHandler) Update(ctx context.Context, in *Article, out *Response) error {
	return h.BlogsHandler.Update(ctx, in, out)
}

func (h *blogsHandler) FetchTopic(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.BlogsHandler.FetchTopic(ctx, in, out)
}

func (h *blogsHandler) CreateTopic(ctx context.Context, in *Topic, out *Response) error {
	return h.BlogsHandler.CreateTopic(ctx, in, out)
}
