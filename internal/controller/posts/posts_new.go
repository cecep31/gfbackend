// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package posts

import (
	"gfbackend/api/posts"
	"gfbackend/internal/controller/posts/v1"
)

func NewV1() posts.IPostsV1 {
	return &v1.ControllerV1{}
}