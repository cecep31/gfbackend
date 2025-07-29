// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package posts

import (
	"context"

	v1 "gfbackend/api/posts/v1"
)

type IPostsV1 interface {
	GetPosts(ctx context.Context, req *v1.GetPostsReq) (res *v1.GetPostsRes, err error)
	GetPost(ctx context.Context, req *v1.GetPostReq) (res *v1.GetPostRes, err error)
	CreatePost(ctx context.Context, req *v1.CreatePostReq) (res *v1.CreatePostRes, err error)
	UpdatePost(ctx context.Context, req *v1.UpdatePostReq) (res *v1.UpdatePostRes, err error)
	DeletePost(ctx context.Context, req *v1.DeletePostReq) (res *v1.DeletePostRes, err error)
}