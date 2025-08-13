package debug

import (
	"context"
	v1 "gfbackend/api/debug/v1"
)

type IDebugV1 interface {
	GetProfile(ctx context.Context, req *v1.GetProfileReq) (res *v1.GetProfileRes, err error)
}