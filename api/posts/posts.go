// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package posts

import (
	"context"

	v1 "gfbackend/api/posts/v1"
)

type IPostsV1 interface {
	// Post management
	GetPosts(ctx context.Context, req *v1.GetPostsReq) (res *v1.GetPostsRes, err error)
	GetPost(ctx context.Context, req *v1.GetPostReq) (res *v1.GetPostRes, err error)
	GetPostBySlugAndUsername(ctx context.Context, req *v1.GetPostBySlugAndUsernameReq) (res *v1.GetPostBySlugAndUsernameRes, err error)
	CreatePost(ctx context.Context, req *v1.CreatePostReq) (res *v1.CreatePostRes, err error)
	UpdatePost(ctx context.Context, req *v1.UpdatePostReq) (res *v1.UpdatePostRes, err error)
	DeletePost(ctx context.Context, req *v1.DeletePostReq) (res *v1.DeletePostRes, err error)
	GetRandomPosts(ctx context.Context, req *v1.GetRandomPostsReq) (res *v1.GetRandomPostsRes, err error)
	GetMyPosts(ctx context.Context, req *v1.GetMyPostsReq) (res *v1.GetMyPostsRes, err error)
	GetPostsByUsername(ctx context.Context, req *v1.GetPostsByUsernameReq) (res *v1.GetPostsByUsernameRes, err error)
	GetPostsByTag(ctx context.Context, req *v1.GetPostsByTagReq) (res *v1.GetPostsByTagRes, err error)
	UploadPostImage(ctx context.Context, req *v1.UploadPostImageReq) (res *v1.UploadPostImageRes, err error)
	
	// Comment management
	GetPostComments(ctx context.Context, req *v1.GetPostCommentsReq) (res *v1.GetPostCommentsRes, err error)
	CreateComment(ctx context.Context, req *v1.CreateCommentReq) (res *v1.CreateCommentRes, err error)
	UpdateComment(ctx context.Context, req *v1.UpdateCommentReq) (res *v1.UpdateCommentRes, err error)
	DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (res *v1.DeleteCommentRes, err error)
	
	// View management
	RecordPostView(ctx context.Context, req *v1.RecordPostViewReq) (res *v1.RecordPostViewRes, err error)
	GetPostViews(ctx context.Context, req *v1.GetPostViewsReq) (res *v1.GetPostViewsRes, err error)
	GetPostViewStats(ctx context.Context, req *v1.GetPostViewStatsReq) (res *v1.GetPostViewStatsRes, err error)
	CheckUserViewedPost(ctx context.Context, req *v1.CheckUserViewedPostReq) (res *v1.CheckUserViewedPostRes, err error)
	
	// Like management
	LikePost(ctx context.Context, req *v1.LikePostReq) (res *v1.LikePostRes, err error)
	UnlikePost(ctx context.Context, req *v1.UnlikePostReq) (res *v1.UnlikePostRes, err error)
	GetPostLikes(ctx context.Context, req *v1.GetPostLikesReq) (res *v1.GetPostLikesRes, err error)
	GetPostLikeStats(ctx context.Context, req *v1.GetPostLikeStatsReq) (res *v1.GetPostLikeStatsRes, err error)
	CheckUserLikedPost(ctx context.Context, req *v1.CheckUserLikedPostReq) (res *v1.CheckUserLikedPostRes, err error)
}