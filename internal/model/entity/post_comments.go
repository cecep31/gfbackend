// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PostComments is the golang structure for table post_comments.
type PostComments struct {
	Id               string      `json:"id"               orm:"id"                 ` //
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         ` //
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         ` //
	Text             string      `json:"text"             orm:"text"               ` //
	PostId           string      `json:"postId"           orm:"post_id"            ` //
	ParrentCommentId int64       `json:"parrentCommentId" orm:"parrent_comment_id" ` //
	CreatedBy        string      `json:"createdBy"        orm:"created_by"         ` //
}
