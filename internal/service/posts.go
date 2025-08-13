package service

import (
	"context"
	"gfbackend/api/posts/v1"
	"gfbackend/internal/dao"
	"gfbackend/internal/model/do"
	"gfbackend/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type IPosts interface {
	GetPosts(ctx context.Context, req *v1.GetPostsReq) (*v1.GetPostsRes, error)
	GetPost(ctx context.Context, req *v1.GetPostReq) (*v1.GetPostRes, error)
	CreatePost(ctx context.Context, req *v1.CreatePostReq) (*v1.CreatePostRes, error)
	UpdatePost(ctx context.Context, req *v1.UpdatePostReq) (*v1.UpdatePostRes, error)
	DeletePost(ctx context.Context, req *v1.DeletePostReq) (*v1.DeletePostRes, error)
}

type sPosts struct{}

func NewPosts() *sPosts {
	return &sPosts{}
}

var (
	insPosts = sPosts{}
)

func Posts() *sPosts {
	return &insPosts
}

func (s *sPosts) init() {
	s = &insPosts
}

func (s *sPosts) GetPosts(ctx context.Context, req *v1.GetPostsReq) (*v1.GetPostsRes, error) {
	var posts []*entity.Posts
	
	query := dao.Posts.Ctx(ctx).Where("deleted_at IS NULL")
	
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Offset < 0 {
		req.Offset = 0
	}
	
	count, err := query.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to count posts")
	}
	
	err = query.Order("created_at DESC").Limit(req.Limit).Offset(req.Offset).Scan(&posts)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get posts")
	}
	
	var resultPosts []v1.Post
	err = gconv.Scan(posts, &resultPosts)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to convert posts")
	}
	
	totalPages := int(count) / req.Limit
	if int(count)%req.Limit > 0 {
		totalPages++
	}
	
	return &v1.GetPostsRes{
		Success: true,
		Message: "Posts retrieved successfully",
		Data:    resultPosts,
		Meta: &v1.MetaData{
			TotalItems: int(count),
			Offset:     req.Offset,
			Limit:      req.Limit,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *sPosts) GetPost(ctx context.Context, req *v1.GetPostReq) (*v1.GetPostRes, error) {
	var post *entity.Posts
	
	err := dao.Posts.Ctx(ctx).Where("id", req.ID).Where("deleted_at IS NULL").Scan(&post)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get post")
	}
	
	if post == nil {
		return nil, gerror.New("post not found")
	}
	
	var resultPost *v1.Post
	err = gconv.Scan(post, &resultPost)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to convert post")
	}
	
	return &v1.GetPostRes{
		Success: true,
		Message: "Post retrieved successfully",
		Data:    resultPost,
	}, nil
}

func (s *sPosts) CreatePost(ctx context.Context, req *v1.CreatePostReq) (*v1.CreatePostRes, error) {
	// Check if slug already exists
	exists, err := dao.Posts.Ctx(ctx).Where("slug", req.Slug).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to check slug existence")
	}
	if exists > 0 {
		return nil, gerror.New("slug already exists")
	}
	
	// Get user ID from context (assuming it's available)
	userId := gconv.String(ctx.Value("userId"))
	if userId == "" {
		userId = "system"
	}
	
	post := do.Posts{
		Title:      req.Title,
		Body:       req.Body,
		Slug:       req.Slug,
		PhotoUrl:   req.PhotoUrl,
		CreatedBy:  "admin", // TODO: Get from authenticated user
		Createbyid: userId,
	}
	
	result, err := dao.Posts.Ctx(ctx).Data(post).Insert()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to create post")
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get post ID")
	}
	
	return &v1.CreatePostRes{
		Success: true,
		Message: "Post created successfully",
		Data: &v1.PostID{
			ID: gconv.String(id),
		},
	}, nil
}

func (s *sPosts) UpdatePost(ctx context.Context, req *v1.UpdatePostReq) (*v1.UpdatePostRes, error) {
	// Check if post exists
	exists, err := dao.Posts.Ctx(ctx).Where("id", req.ID).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to check post existence")
	}
	if exists == 0 {
		return nil, gerror.New("post not found")
	}
	
	// Check if new slug already exists (if slug is being updated)
	if req.Slug != "" {
		exists, err = dao.Posts.Ctx(ctx).Where("slug", req.Slug).Where("id !=", req.ID).Count()
		if err != nil {
			return nil, gerror.Wrap(err, "failed to check slug existence")
		}
		if exists > 0 {
			return nil, gerror.New("slug already exists")
		}
	}
	
	post := do.Posts{}
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Body != "" {
		post.Body = req.Body
	}
	if req.Slug != "" {
		post.Slug = req.Slug
	}
	if req.PhotoUrl != "" {
		post.PhotoUrl = req.PhotoUrl
	}
	

	
	_, err = dao.Posts.Ctx(ctx).Data(post).Where("id", req.ID).Update()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to update post")
	}
	
	return &v1.UpdatePostRes{
		Success: true,
		Message: "Post updated successfully",
	}, nil
}

func (s *sPosts) DeletePost(ctx context.Context, req *v1.DeletePostReq) (*v1.DeletePostRes, error) {
	// Check if post exists
	exists, err := dao.Posts.Ctx(ctx).Where("id", req.ID).Where("deleted_at IS NULL").Count()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to check post existence")
	}
	if exists == 0 {
		return nil, gerror.New("post not found")
	}
	
	_, err = dao.Posts.Ctx(ctx).Where("id", req.ID).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "failed to delete post")
	}
	
	return &v1.DeletePostRes{
		Success: true,
		Message: "Post deleted successfully",
	}, nil
}