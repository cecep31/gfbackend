package v1

import (
	"context"

	"gfbackend/api/posts/v1"
	"gfbackend/internal/service"
)

func (c *ControllerV1) GetPosts(ctx context.Context, req *v1.GetPostsReq) (res *v1.GetPostsRes, err error) {
	return service.Posts().GetPosts(ctx, req)
}

func (c *ControllerV1) GetPost(ctx context.Context, req *v1.GetPostReq) (res *v1.GetPostRes, err error) {
	return service.Posts().GetPost(ctx, req)
}

func (c *ControllerV1) CreatePost(ctx context.Context, req *v1.CreatePostReq) (res *v1.CreatePostRes, err error) {
	return service.Posts().CreatePost(ctx, req)
}

func (c *ControllerV1) UpdatePost(ctx context.Context, req *v1.UpdatePostReq) (res *v1.UpdatePostRes, err error) {
	return service.Posts().UpdatePost(ctx, req)
}

func (c *ControllerV1) DeletePost(ctx context.Context, req *v1.DeletePostReq) (res *v1.DeletePostRes, err error) {
	return service.Posts().DeletePost(ctx, req)
}