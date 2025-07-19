// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Likes is the golang structure for table likes.
type Likes struct {
	Id        int         `json:"id"        orm:"id"         ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	PostId    string      `json:"postId"    orm:"post_id"    ` //
	UserId    string      `json:"userId"    orm:"user_id"    ` //
}
