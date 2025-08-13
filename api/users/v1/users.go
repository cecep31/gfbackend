package v1

import "github.com/gogf/gf/v2/frame/g"

// User management endpoints
type GetUsersReq struct {
	g.Meta `path:"/users" tags:"Users" method:"get" summary:"Get all users (Admin only)"`
	Offset int `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int `json:"limit" d:"10" description:"Number of records to return"`
}

type GetUsersRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []User    `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

type GetUserReq struct {
	g.Meta `path:"/users/{id}" tags:"Users" method:"get" summary:"Get user by ID"`
	ID     string `json:"id" v:"required" description:"User ID"`
}

type GetUserRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *User  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type GetCurrentUserReq struct {
	g.Meta `path:"/users/me" tags:"Users" method:"get" summary:"Get current user"`
}

type GetCurrentUserRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *User  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type DeleteUserReq struct {
	g.Meta `path:"/users/{id}" tags:"Users" method:"delete" summary:"Delete user (Admin only)"`
	ID     string `json:"id" v:"required" description:"User ID"`
}

type DeleteUserRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Follow system endpoints
type FollowUserReq struct {
	g.Meta `path:"/users/follow" tags:"Users" method:"post" summary:"Follow a user"`
	UserID string `json:"user_id" v:"required" description:"User ID to follow"`
}

type FollowUserRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type UnfollowUserReq struct {
	g.Meta `path:"/users/{id}/follow" tags:"Users" method:"delete" summary:"Unfollow a user"`
	ID     string `json:"id" v:"required" description:"User ID to unfollow"`
}

type UnfollowUserRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type CheckFollowStatusReq struct {
	g.Meta `path:"/users/{id}/follow-status" tags:"Users" method:"get" summary:"Check follow status"`
	ID     string `json:"id" v:"required" description:"User ID"`
}

type CheckFollowStatusRes struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *FollowStatus `json:"data,omitempty"`
}

type GetMutualFollowsReq struct {
	g.Meta `path:"/users/{id}/mutual-follows" tags:"Users" method:"get" summary:"Get mutual follows"`
	ID     string `json:"id" v:"required" description:"User ID"`
}

type GetMutualFollowsRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type GetUserFollowersReq struct {
	g.Meta `path:"/users/{id}/followers" tags:"Users" method:"get" summary:"Get user followers"`
	ID     string `json:"id" v:"required" description:"User ID"`
	Offset int    `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int    `json:"limit" d:"10" description:"Number of records to return"`
}

type GetUserFollowersRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []User    `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

type GetUserFollowingReq struct {
	g.Meta `path:"/users/{id}/following" tags:"Users" method:"get" summary:"Get user following"`
	ID     string `json:"id" v:"required" description:"User ID"`
	Offset int    `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int    `json:"limit" d:"10" description:"Number of records to return"`
}

type GetUserFollowingRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []User    `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

type GetFollowStatsReq struct {
	g.Meta `path:"/users/{id}/follow-stats" tags:"Users" method:"get" summary:"Get follow statistics"`
	ID     string `json:"id" v:"required" description:"User ID"`
}

type GetFollowStatsRes struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *FollowStats `json:"data,omitempty"`
}

// Data models
type User struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Image          string `json:"image,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	IsSuperAdmin   bool   `json:"is_super_admin,omitempty"`
	FollowersCount int    `json:"followers_count,omitempty"`
	FollowingCount int    `json:"following_count,omitempty"`
	IsFollowing    *bool  `json:"is_following,omitempty"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type MetaData struct {
	TotalItems int `json:"total_items"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}

type FollowStatus struct {
	IsFollowing bool   `json:"is_following"`
	UserID      string `json:"user_id"`
	FollowedID  string `json:"followed_id"`
}

type FollowStats struct {
	UserID         string `json:"user_id"`
	FollowersCount int    `json:"followers_count"`
	FollowingCount int    `json:"following_count"`
}
