package tags

import (
	"context"
	v1 "gfbackend/api/tags/v1"
)

type ITagsV1 interface {
	CreateTag(ctx context.Context, req *v1.CreateTagReq) (res *v1.CreateTagRes, err error)
	GetAllTags(ctx context.Context, req *v1.GetAllTagsReq) (res *v1.GetAllTagsRes, err error)
	GetTag(ctx context.Context, req *v1.GetTagReq) (res *v1.GetTagRes, err error)
	UpdateTag(ctx context.Context, req *v1.UpdateTagReq) (res *v1.UpdateTagRes, err error)
	DeleteTag(ctx context.Context, req *v1.DeleteTagReq) (res *v1.DeleteTagRes, err error)
}