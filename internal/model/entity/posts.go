// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Posts is the golang structure for table posts.
type Posts struct {
	Id         string      `json:"id"         orm:"id"         ` //
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at" ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at" ` //
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
	Title      string      `json:"title"      orm:"title"      ` //
	CreatedBy  string      `json:"createdBy"  orm:"created_by" ` //
	Body       string      `json:"body"       orm:"body"       ` //
	Slug       string      `json:"slug"       orm:"slug"       ` //
	Createbyid string      `json:"createbyid" orm:"createbyid" ` //
	PhotoUrl   string      `json:"photoUrl"   orm:"photo_url"  ` //
	Published  bool        `json:"published"  orm:"published"  ` //
}
