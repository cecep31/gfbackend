// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	v1 "gfbackend/api/auth/v1"
)

type IAuthV1 interface {
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	CheckUsername(ctx context.Context, req *v1.CheckUsernameReq) (res *v1.CheckUsernameRes, err error)
}