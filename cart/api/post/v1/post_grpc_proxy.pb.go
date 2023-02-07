// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc-proxy v1.2.0
// - protoc             (unknown)
// source: cart/api/post/v1/post.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var _ PostServiceServer = (*postServiceClientProxy)(nil)

// postServiceClientProxy is the proxy to turn PostService client to server interface.
type postServiceClientProxy struct {
	cc PostServiceClient
}

func NewPostServiceClientProxy(cc PostServiceClient) PostServiceServer {
	return &postServiceClientProxy{cc}
}

func (c *postServiceClientProxy) ListPost(ctx context.Context, in *ListPostRequest) (*ListPostReply, error) {
	return c.cc.ListPost(ctx, in)
}
func (c *postServiceClientProxy) GetPost(ctx context.Context, in *GetPostRequest) (*Post, error) {
	return c.cc.GetPost(ctx, in)
}
func (c *postServiceClientProxy) CreatePost(ctx context.Context, in *CreatePostRequest) (*Post, error) {
	return c.cc.CreatePost(ctx, in)
}
func (c *postServiceClientProxy) UpdatePost(ctx context.Context, in *UpdatePostRequest) (*Post, error) {
	return c.cc.UpdatePost(ctx, in)
}
func (c *postServiceClientProxy) DeletePost(ctx context.Context, in *DeletePostRequest) (*DeletePostReply, error) {
	return c.cc.DeletePost(ctx, in)
}