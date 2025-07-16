package users

import (
	"context"
	"time"

	v1 "gfbackend/api/users/v1"
)

func (c *ControllerV1) GetUsers(ctx context.Context, req *v1.GetUsersReq) (res *v1.GetUsersRes, err error) {
	// Mock data for demonstration
	users := []v1.User{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			ID:        2,
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	return &v1.GetUsersRes{
		Users: users,
		Total: len(users),
	}, nil
}

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	// Mock data for demonstration
	user := &v1.User{
		ID:        req.ID,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &v1.GetUserRes{
		User: user,
	}, nil
}

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	// Mock data for demonstration - in real app, this would save to database
	user := &v1.User{
		ID:        3,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &v1.CreateUserRes{
		User: user,
	}, nil
}

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	// Mock data for demonstration - in real app, this would update in database
	user := &v1.User{
		ID:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &v1.UpdateUserRes{
		User: user,
	}, nil
}

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	// Mock deletion - in real app, this would delete from database
	return &v1.DeleteUserRes{
		Success: true,
	}, nil
}
