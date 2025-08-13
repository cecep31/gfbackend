package users

import (
	"context"

	v1 "gfbackend/api/users/v1"
	"gfbackend/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetUsers(ctx context.Context, req *v1.GetUsersReq) (res *v1.GetUsersRes, err error) {
	// Get users from database
	userService := service.NewUsers()
	entityUsers, err := userService.GetAllUsers(ctx)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get users from database")
	}

	// Convert entity users to API response format
	users := make([]v1.User, 0, len(entityUsers))
	for _, entityUser := range entityUsers {
		if entityUser == nil {
			continue
		}
		

		
		// Combine first and last name
		name := entityUser.FirstName
		if entityUser.LastName != "" {
			if name != "" {
				name += " " + entityUser.LastName
			} else {
				name = entityUser.LastName
			}
		}
		if name == "" {
			name = entityUser.Username
		}
		
		user := v1.User{
			ID:        entityUser.Id,
			Email:     entityUser.Email,
			Name:      name,
			Username:  entityUser.Username,
			Image:     entityUser.Image,
			FirstName: entityUser.FirstName,
			LastName:  entityUser.LastName,
			CreatedAt: entityUser.CreatedAt.String(),
			UpdatedAt: entityUser.UpdatedAt.String(),
		}
		
		// Format timestamps
		if entityUser.CreatedAt != nil {
			user.CreatedAt = entityUser.CreatedAt.Format("2006-01-02 15:04:05")
		}
		if entityUser.UpdatedAt != nil {
			user.UpdatedAt = entityUser.UpdatedAt.Format("2006-01-02 15:04:05")
		}
		
		users = append(users, user)
	}

	return &v1.GetUsersRes{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    users,
		Meta: &v1.MetaData{
			TotalItems: len(users),
			Offset:     req.Offset,
			Limit:      req.Limit,
			TotalPages: (len(users) + req.Limit - 1) / req.Limit,
		},
	}, nil
}

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	// Get user from database
	userService := service.NewUsers()
	entityUser, err := userService.GetUserByIDString(ctx, req.ID)
	if err != nil {
		return &v1.GetUserRes{
			Success: false,
			Message: "User not found",
			Error:   err.Error(),
		}, nil
	}
	
	if entityUser == nil {
		return &v1.GetUserRes{
			Success: false,
			Message: "User not found",
			Error:   "User does not exist",
		}, nil
	}
	
	// Combine first and last name
	name := entityUser.FirstName
	if entityUser.LastName != "" {
		if name != "" {
			name += " " + entityUser.LastName
		} else {
			name = entityUser.LastName
		}
	}
	if name == "" {
		name = entityUser.Username
	}
	
	user := &v1.User{
		ID:        entityUser.Id,
		Email:     entityUser.Email,
		Name:      name,
		Username:  entityUser.Username,
		Image:     entityUser.Image,
		FirstName: entityUser.FirstName,
		LastName:  entityUser.LastName,
		CreatedAt: entityUser.CreatedAt.String(),
		UpdatedAt: entityUser.UpdatedAt.String(),
	}
	
	return &v1.GetUserRes{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	}, nil
}

func (c *ControllerV1) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserReq) (res *v1.GetCurrentUserRes, err error) {
	// TODO: Get current user from authentication context
	return &v1.GetCurrentUserRes{
		Success: false,
		Message: "Not implemented",
		Error:   "Method not implemented",
	}, nil
}

func (c *ControllerV1) FollowUser(ctx context.Context, req *v1.FollowUserReq) (res *v1.FollowUserRes, err error) {
	return &v1.FollowUserRes{
		Success: false,
		Message: "Not implemented",
	}, nil
}

func (c *ControllerV1) UnfollowUser(ctx context.Context, req *v1.UnfollowUserReq) (res *v1.UnfollowUserRes, err error) {
	return &v1.UnfollowUserRes{
		Success: false,
		Message: "Not implemented",
	}, nil
}

func (c *ControllerV1) CheckFollowStatus(ctx context.Context, req *v1.CheckFollowStatusReq) (res *v1.CheckFollowStatusRes, err error) {
	return &v1.CheckFollowStatusRes{
		Success: false,
		Message: "Not implemented",
	}, nil
}

func (c *ControllerV1) GetMutualFollows(ctx context.Context, req *v1.GetMutualFollowsReq) (res *v1.GetMutualFollowsRes, err error) {
	return &v1.GetMutualFollowsRes{
		Success: false,
		Message: "Not implemented",
	}, nil
}

func (c *ControllerV1) GetUserFollowers(ctx context.Context, req *v1.GetUserFollowersReq) (res *v1.GetUserFollowersRes, err error) {
	return &v1.GetUserFollowersRes{
		Success: false,
		Message: "Not implemented",
	}, nil
}

func (c *ControllerV1) GetUserFollowing(ctx context.Context, req *v1.GetUserFollowingReq) (res *v1.GetUserFollowingRes, err error) {
	return &v1.GetUserFollowingRes{
		Success: false,
		Message: "Not implemented",
	}, nil
}

func (c *ControllerV1) GetFollowStats(ctx context.Context, req *v1.GetFollowStatsReq) (res *v1.GetFollowStatsRes, err error) {
	return &v1.GetFollowStatsRes{
		Success: false,
		Message: "Not implemented",
	}, nil
}

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	// Check if user exists first
	userService := service.NewUsers()
	existingUser, err := userService.GetUserByIDString(ctx, req.ID)
	if err != nil {
		return &v1.DeleteUserRes{
			Success: false,
			Message: "Failed to delete user",
		}, nil
	}
	
	if existingUser == nil {
		return &v1.DeleteUserRes{
			Success: false,
			Message: "User not found",
		}, nil
	}
	
	// Delete the user
	err = userService.DeleteUserByString(ctx, req.ID)
	if err != nil {
		return &v1.DeleteUserRes{
			Success: false,
			Message: "Failed to delete user",
		}, nil
	}
	
	return &v1.DeleteUserRes{
		Success: true,
		Message: "User deleted successfully",
		Data:    nil,
	}, nil
}
