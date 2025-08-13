// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package users

import (
	"context"

	v1 "gfbackend/api/users/v1"
)

type IUsersV1 interface {
	// User management
	GetUsers(ctx context.Context, req *v1.GetUsersReq) (res *v1.GetUsersRes, err error)
	GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error)
	GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserReq) (res *v1.GetCurrentUserRes, err error)
	DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error)
	
	// Follow system
	FollowUser(ctx context.Context, req *v1.FollowUserReq) (res *v1.FollowUserRes, err error)
	UnfollowUser(ctx context.Context, req *v1.UnfollowUserReq) (res *v1.UnfollowUserRes, err error)
	CheckFollowStatus(ctx context.Context, req *v1.CheckFollowStatusReq) (res *v1.CheckFollowStatusRes, err error)
	GetMutualFollows(ctx context.Context, req *v1.GetMutualFollowsReq) (res *v1.GetMutualFollowsRes, err error)
	GetUserFollowers(ctx context.Context, req *v1.GetUserFollowersReq) (res *v1.GetUserFollowersRes, err error)
	GetUserFollowing(ctx context.Context, req *v1.GetUserFollowingReq) (res *v1.GetUserFollowingRes, err error)
	GetFollowStats(ctx context.Context, req *v1.GetFollowStatsReq) (res *v1.GetFollowStatsRes, err error)
}
