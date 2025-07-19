
package service

import (
	"context"
	"gfbackend/internal/model/do"
	"gfbackend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
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
	_, err := g.Model("users").Ctx(ctx).Data(user).Insert()
	return err
}

func (s *sUsers) GetUserByID(ctx context.Context, id int) (*entity.Users, error) {
	var user *entity.Users
	err := g.Model("users").Ctx(ctx).Where("id", id).Scan(&user)
	return user, err
}

func (s *sUsers) UpdateUser(ctx context.Context, user *do.Users) error {
	_, err := g.Model("users").Ctx(ctx).Data(user).Where("id", user.Id).Update()
	return err
}

func (s *sUsers) DeleteUser(ctx context.Context, id int) error {
	_, err := g.Model("users").Ctx(ctx).Where("id", id).Delete()
	return err
}

func (s *sUsers) GetAllUsers(ctx context.Context) ([]*entity.Users, error) {
	var users []*entity.Users
	err := g.Model("users").Ctx(ctx).Scan(&users)
	return users, err
}
