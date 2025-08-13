package v1

import "github.com/gogf/gf/v2/frame/g"

// Post management endpoints
type GetPostsReq struct {
	g.Meta `path:"/posts" tags:"Posts" method:"get" summary:"Get all posts"`
	Offset int `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int `json:"limit" d:"10" description:"Number of records to return"`
}

type GetPostsRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Post    `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

type GetPostReq struct {
	g.Meta `path:"/posts/{id}" tags:"Posts" method:"get" summary:"Get post by ID"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type GetPostRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Post  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type GetPostBySlugAndUsernameReq struct {
	g.Meta   `path:"/posts/u/{username}/{slug}" tags:"Posts" method:"get" summary:"Get post by username and slug"`
	Username string `json:"username" v:"required" description:"Author username"`
	Slug     string `json:"slug" v:"required" description:"Post slug"`
}

type GetPostBySlugAndUsernameRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Post  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type CreatePostReq struct {
	g.Meta   `path:"/posts" tags:"Posts" method:"post" summary:"Create new post"`
	Title    string   `json:"title" v:"required|min:7" description:"Post title"`
	PhotoUrl string   `json:"photo_url" description:"Post photo URL"`
	Slug     string   `json:"slug" v:"required|min:7" description:"Post slug"`
	Body     string   `json:"body" v:"required|min:10" description:"Post content"`
	Tags     []string `json:"tags" description:"Post tags"`
}

type CreatePostRes struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    *PostID  `json:"data,omitempty"`
	Error   string   `json:"error,omitempty"`
}

type UpdatePostReq struct {
	g.Meta   `path:"/posts/{id}" tags:"Posts" method:"put" summary:"Update post"`
	ID       string   `json:"id" v:"required" description:"Post ID"`
	Title    string   `json:"title" v:"min:7" description:"Post title"`
	PhotoUrl string   `json:"photo_url" description:"Post photo URL"`
	Slug     string   `json:"slug" v:"min:7" description:"Post slug"`
	Body     string   `json:"body" v:"min:10" description:"Post content"`
	Tags     []string `json:"tags" description:"Post tags"`
}

type UpdatePostRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Post  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type DeletePostReq struct {
	g.Meta `path:"/posts/{id}" tags:"Posts" method:"delete" summary:"Delete post"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type DeletePostRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type GetRandomPostsReq struct {
	g.Meta `path:"/posts/random" tags:"Posts" method:"get" summary:"Get random posts"`
	Limit  int `json:"limit" d:"9" description:"Number of posts to return (max 50)"`
}

type GetRandomPostsRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []Post `json:"data"`
}

type GetMyPostsReq struct {
	g.Meta `path:"/posts/mine" tags:"Posts" method:"get" summary:"Get current user's posts"`
	Offset int `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int `json:"limit" d:"10" description:"Number of records to return"`
}

type GetMyPostsRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Post    `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

type GetPostsByUsernameReq struct {
	g.Meta   `path:"/posts/username/{username}" tags:"Posts" method:"get" summary:"Get posts by username"`
	Username string `json:"username" v:"required" description:"Author username"`
	Offset   int    `json:"offset" d:"0" description:"Number of records to skip"`
	Limit    int    `json:"limit" d:"10" description:"Number of records to return"`
}

type GetPostsByUsernameRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Post    `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

type GetPostsByTagReq struct {
	g.Meta `path:"/posts/tag/{tag}" tags:"Posts" method:"get" summary:"Get posts by tag"`
	Tag    string `json:"tag" v:"required" description:"Tag name"`
	Offset int    `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int    `json:"limit" d:"10" description:"Number of records to return"`
}

type GetPostsByTagRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Post    `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

type UploadPostImageReq struct {
	g.Meta `path:"/posts/image" tags:"Posts" method:"post" summary:"Upload post image"`
	// File upload will be handled in the controller
}

type UploadPostImageRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Image `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// Comment endpoints
type GetPostCommentsReq struct {
	g.Meta `path:"/posts/{id}/comments" tags:"Comments" method:"get" summary:"Get post comments"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type GetPostCommentsRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Comment `json:"data"`
}

type CreateCommentReq struct {
	g.Meta  `path:"/posts/{id}/comments" tags:"Comments" method:"post" summary:"Create comment"`
	ID      string `json:"id" v:"required" description:"Post ID"`
	Content string `json:"content" v:"required" description:"Comment content"`
}

type CreateCommentRes struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    *Comment `json:"data,omitempty"`
	Error   string   `json:"error,omitempty"`
}

type UpdateCommentReq struct {
	g.Meta    `path:"/posts/{id}/comments/{comment_id}" tags:"Comments" method:"put" summary:"Update comment"`
	ID        string `json:"id" v:"required" description:"Post ID"`
	CommentID string `json:"comment_id" v:"required" description:"Comment ID"`
	Content   string `json:"content" v:"required" description:"Comment content"`
}

type UpdateCommentRes struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    *Comment `json:"data,omitempty"`
	Error   string   `json:"error,omitempty"`
}

type DeleteCommentReq struct {
	g.Meta    `path:"/posts/{id}/comments/{comment_id}" tags:"Comments" method:"delete" summary:"Delete comment"`
	ID        string `json:"id" v:"required" description:"Post ID"`
	CommentID string `json:"comment_id" v:"required" description:"Comment ID"`
}

type DeleteCommentRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// View endpoints
type RecordPostViewReq struct {
	g.Meta `path:"/posts/{id}/view" tags:"Views" method:"post" summary:"Record post view"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type RecordPostViewRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type GetPostViewsReq struct {
	g.Meta `path:"/posts/{id}/views" tags:"Views" method:"get" summary:"Get post views"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type GetPostViewsRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []View `json:"data"`
}

type GetPostViewStatsReq struct {
	g.Meta `path:"/posts/{id}/view-stats" tags:"Views" method:"get" summary:"Get post view statistics"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type GetPostViewStatsRes struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *ViewStats `json:"data,omitempty"`
}

type CheckUserViewedPostReq struct {
	g.Meta `path:"/posts/{id}/viewed" tags:"Views" method:"get" summary:"Check if user viewed post"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type CheckUserViewedPostRes struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *ViewStatus `json:"data,omitempty"`
}

// Like endpoints
type LikePostReq struct {
	g.Meta `path:"/posts/{id}/like" tags:"Likes" method:"post" summary:"Like post"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type LikePostRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error,omitempty"`
}

type UnlikePostReq struct {
	g.Meta `path:"/posts/{id}/like" tags:"Likes" method:"delete" summary:"Unlike post"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type UnlikePostRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error,omitempty"`
}

type GetPostLikesReq struct {
	g.Meta `path:"/posts/{id}/likes" tags:"Likes" method:"get" summary:"Get post likes"`
	ID     string `json:"id" v:"required" description:"Post ID"`
	Offset int    `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int    `json:"limit" d:"10" description:"Number of records to return"`
}

type GetPostLikesRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    *LikeData `json:"data,omitempty"`
}

type GetPostLikeStatsReq struct {
	g.Meta `path:"/posts/{id}/like-stats" tags:"Likes" method:"get" summary:"Get post like statistics"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type GetPostLikeStatsRes struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *LikeStats `json:"data,omitempty"`
}

type CheckUserLikedPostReq struct {
	g.Meta `path:"/posts/{id}/liked" tags:"Likes" method:"get" summary:"Check if user liked post"`
	ID     string `json:"id" v:"required" description:"Post ID"`
}

type CheckUserLikedPostRes struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *LikeStatus `json:"data,omitempty"`
}

// Data models
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	PhotoUrl  string    `json:"photo_url,omitempty"`
	Body      string    `json:"body"`
	Slug      string    `json:"slug"`
	ViewCount int       `json:"view_count,omitempty"`
	LikeCount int       `json:"like_count,omitempty"`
	Creator   *Creator  `json:"creator,omitempty"`
	Tags      []Tag     `json:"tags,omitempty"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type Creator struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID        string   `json:"id"`
	Content   string   `json:"content"`
	Author    *Creator `json:"author,omitempty"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostID struct {
	ID string `json:"id"`
}

type Image struct {
	URL string `json:"url"`
}

type MetaData struct {
	TotalItems int `json:"total_items"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}

type View struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
	CreatedAt string `json:"created_at"`
}

type ViewStats struct {
	PostID    string `json:"post_id"`
	ViewCount int    `json:"view_count"`
}

type ViewStatus struct {
	HasViewed bool   `json:"has_viewed"`
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
}

type Like struct {
	ID        string   `json:"id"`
	PostID    string   `json:"post_id"`
	UserID    string   `json:"user_id"`
	User      *Creator `json:"user,omitempty"`
	CreatedAt string   `json:"created_at"`
}

type LikeData struct {
	Likes  []Like `json:"likes"`
	Total  int    `json:"total"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type LikeStats struct {
	PostID     string `json:"post_id"`
	TotalLikes int    `json:"total_likes"`
}

type LikeStatus struct {
	HasLiked bool   `json:"has_liked"`
	PostID   string `json:"post_id"`
	UserID   string `json:"user_id"`
}