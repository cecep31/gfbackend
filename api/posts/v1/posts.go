package v1

import "github.com/gogf/gf/v2/frame/g"

type GetPostsReq struct {
	g.Meta `path:"/posts" tags:"Posts" method:"get" summary:"Get all posts"`
	Page   int `json:"page" d:"1" description:"Page number"`
	Limit  int `json:"limit" d:"10" description:"Number of items per page"`
}

type GetPostsRes struct {
	Posts []Post `json:"posts"`
	Total int    `json:"total"`
}

type GetPostReq struct {
	g.Meta `path:"/posts/:id" tags:"Posts" method:"get" summary:"Get post by ID"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type GetPostRes struct {
	Post *Post `json:"post"`
}

type CreatePostReq struct {
	g.Meta    `path:"/posts" tags:"Posts" method:"post" summary:"Create new post"`
	Title     string `json:"title" v:"required" description:"Post title"`
	Body      string `json:"body" v:"required" description:"Post content"`
	Slug      string `json:"slug" v:"required" description:"Post slug"`
	PhotoUrl  string `json:"photoUrl" description:"Post photo URL"`
	Published bool   `json:"published" d:"false" description:"Post published status"`
}

type CreatePostRes struct {
	Post *Post `json:"post"`
}

type UpdatePostReq struct {
	g.Meta    `path:"/posts/:id" tags:"Posts" method:"put" summary:"Update post"`
	ID        string `json:"id" v:"required" description:"Post ID"`
	Title     string `json:"title" description:"Post title"`
	Body      string `json:"body" description:"Post content"`
	Slug      string `json:"slug" description:"Post slug"`
	PhotoUrl  string `json:"photoUrl" description:"Post photo URL"`
	Published bool   `json:"published" description:"Post published status"`
}

type UpdatePostRes struct {
	Post *Post `json:"post"`
}

type DeletePostReq struct {
	g.Meta `path:"/posts/:id" tags:"Posts" method:"delete" summary:"Delete post"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type DeletePostRes struct {
	Success bool `json:"success"`
}

type Post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Slug      string `json:"slug"`
	PhotoUrl  string `json:"photoUrl"`
	Published bool   `json:"published"`
	CreatedBy string `json:"createdBy"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}