package service

import (
	"context"
	"gfbackend/internal/dao"
	"gfbackend/internal/model/do"
	"gfbackend/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

type IUsers interface {
	CreateUser(ctx context.Context, user *do.Users) error
	GetUserByID(ctx context.Context, id int) (*entity.Users, error)
	UpdateUser(ctx context.Context, user *do.Users) error
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]*entity.Users, error)
}

type sUsers struct{}

func NewUsers() *sUsers {
	return &sUsers{}
}

func (s *sUsers) CreateUser(ctx context.Context, user *do.Users) error {
	// Check if email already exists
	exists, err := dao.Users.ExistsByEmail(ctx, user.Email.(string))
	if err != nil {
		return gerror.Wrap(err, "failed to check email existence")
	}
	if exists {
		return gerror.New("email already exists")
	}

	// Check if username already exists (if provided)
	if user.Username != nil {
		exists, err = dao.Users.ExistsByUsername(ctx, user.Username.(string))
		if err != nil {
			return gerror.Wrap(err, "failed to check username existence")
		}
		if exists {
			return gerror.New("username already exists")
		}
	}

	_, err = dao.Users.Ctx(ctx).Data(user).Insert()
	if err != nil {
		return gerror.Wrap(err, "failed to insert user")
	}
	return nil
}

func (s *sUsers) GetUserByID(ctx context.Context, id int) (*entity.Users, error) {
	var user *entity.Users
	err := dao.Users.Ctx(ctx).Where("id", id).Scan(&user)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get user by ID")
	}
	return user, nil
}

func (s *sUsers) UpdateUser(ctx context.Context, user *do.Users) error {
	_, err := dao.Users.Ctx(ctx).Data(user).Where("id", user.Id).Update()
	if err != nil {
		return gerror.Wrap(err, "failed to update user")
	}
	return nil
}

func (s *sUsers) DeleteUser(ctx context.Context, id int) error {
	_, err := dao.Users.Ctx(ctx).Where("id", id).Delete()
	if err != nil {
		return gerror.Wrap(err, "failed to delete user")
	}
	return nil
}

func (s *sUsers) GetAllUsers(ctx context.Context) ([]*entity.Users, error) {
	var users []*entity.Users
	err := dao.Users.Ctx(ctx).Where("deleted_at IS NULL").Scan(&users)
	if err != nil {
		return nil, gerror.Wrap(err, "failed to get all users")
	}
	return users, nil
}
