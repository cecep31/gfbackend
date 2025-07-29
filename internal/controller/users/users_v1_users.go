package users

import (
	"context"
	"strconv"
	"strings"

	v1 "gfbackend/api/users/v1"
	"gfbackend/internal/dao"
	"gfbackend/internal/model/do"
	"gfbackend/internal/model/entity"
	"gfbackend/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
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
		
		// Convert ID from string to int
		id, err := strconv.Atoi(entityUser.Id)
		if err != nil {
			continue // Skip invalid IDs
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
			ID:    id,
			Name:  name,
			Email: entityUser.Email,
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
		Users: users,
		Total: len(users),
	}, nil
}

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	// Get user from database
	userService := service.NewUsers()
	entityUser, err := userService.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get user from database")
	}
	
	if entityUser == nil {
		return nil, gerror.New("user not found")
	}
	
	// Convert ID from string to int
	id, err := strconv.Atoi(entityUser.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "invalid user ID format")
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
		ID:    id,
		Name:  name,
		Email: entityUser.Email,
	}
	
	// Format timestamps
	if entityUser.CreatedAt != nil {
		user.CreatedAt = entityUser.CreatedAt.Format("2006-01-02 15:04:05")
	}
	if entityUser.UpdatedAt != nil {
		user.UpdatedAt = entityUser.UpdatedAt.Format("2006-01-02 15:04:05")
	}

	return &v1.GetUserRes{
		User: user,
	}, nil
}

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	// Parse name into first and last name
	var firstName, lastName string
	nameParts := strings.Fields(req.Name)
	if len(nameParts) > 0 {
		firstName = nameParts[0]
		if len(nameParts) > 1 {
			lastName = strings.Join(nameParts[1:], " ")
		}
	}
	
	// Create user data object
	userData := &do.Users{
		FirstName: firstName,
		LastName:  lastName,
		Email:     req.Email,
		Password:  req.Password, // In production, this should be hashed
		Username:  req.Email, // Use email as username for now
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}
	
	// Save to database
	userService := service.NewUsers()
	err = userService.CreateUser(ctx, userData)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to create user in database")
	}
	
	// Get the created user to return with ID
	// Since we don't have the ID from create, we'll find by email
	var createdUser *entity.Users
	err = dao.Users.Ctx(ctx).Where("email", req.Email).OrderDesc("created_at").Scan(&createdUser)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to retrieve created user")
	}
	
	if createdUser == nil {
		return nil, gerror.New("failed to find created user")
	}
	
	// Convert ID from string to int
	id, err := strconv.Atoi(createdUser.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "invalid user ID format")
	}
	
	user := &v1.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}
	
	// Format timestamps
	if createdUser.CreatedAt != nil {
		user.CreatedAt = createdUser.CreatedAt.Format("2006-01-02 15:04:05")
	}
	if createdUser.UpdatedAt != nil {
		user.UpdatedAt = createdUser.UpdatedAt.Format("2006-01-02 15:04:05")
	}

	return &v1.CreateUserRes{
		User: user,
	}, nil
}

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	// First check if user exists
	userService := service.NewUsers()
	existingUser, err := userService.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get user from database")
	}
	
	if existingUser == nil {
		return nil, gerror.New("user not found")
	}
	
	// Parse name into first and last name if provided
	var firstName, lastName string
	if req.Name != "" {
		nameParts := strings.Fields(req.Name)
		if len(nameParts) > 0 {
			firstName = nameParts[0]
			if len(nameParts) > 1 {
				lastName = strings.Join(nameParts[1:], " ")
			}
		}
	} else {
		// Keep existing names if not provided
		firstName = existingUser.FirstName
		lastName = existingUser.LastName
	}
	
	// Create update data object
	updateData := &do.Users{
		Id:        existingUser.Id,
		FirstName: firstName,
		LastName:  lastName,
		UpdatedAt: gtime.Now(),
	}
	
	// Update email if provided
	if req.Email != "" {
		updateData.Email = req.Email
	}
	
	// Update password if provided
	if req.Password != "" {
		updateData.Password = req.Password // In production, this should be hashed
	}
	
	// Update in database
	err = userService.UpdateUser(ctx, updateData)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to update user in database")
	}
	
	// Get updated user
	updatedUser, err := userService.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get updated user")
	}
	
	// Convert ID from string to int
	id, err := strconv.Atoi(updatedUser.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "invalid user ID format")
	}
	
	// Combine first and last name
	name := updatedUser.FirstName
	if updatedUser.LastName != "" {
		if name != "" {
			name += " " + updatedUser.LastName
		} else {
			name = updatedUser.LastName
		}
	}
	if name == "" {
		name = updatedUser.Username
	}
	
	user := &v1.User{
		ID:    id,
		Name:  name,
		Email: updatedUser.Email,
	}
	
	// Format timestamps
	if updatedUser.CreatedAt != nil {
		user.CreatedAt = updatedUser.CreatedAt.Format("2006-01-02 15:04:05")
	}
	if updatedUser.UpdatedAt != nil {
		user.UpdatedAt = updatedUser.UpdatedAt.Format("2006-01-02 15:04:05")
	}

	return &v1.UpdateUserRes{
		User: user,
	}, nil
}

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	// First check if user exists
	userService := service.NewUsers()
	existingUser, err := userService.GetUserByID(ctx, req.ID)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get user from database")
	}
	
	if existingUser == nil {
		return &v1.DeleteUserRes{
			Success: false,
		}, gerror.New("user not found")
	}
	
	// Delete user from database
	err = userService.DeleteUser(ctx, req.ID)
	if err != nil {
		return &v1.DeleteUserRes{
			Success: false,
		}, gerror.Wrap(err, "failed to delete user from database")
	}
	
	return &v1.DeleteUserRes{
		Success: true,
	}, nil
}
