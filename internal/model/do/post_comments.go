// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PostComments is the golang structure of table post_comments for DAO operations like Where/Data.
type PostComments struct {
	g.Meta           `orm:"table:post_comments, do:true"`
	Id               interface{} //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	DeletedAt        *gtime.Time //
	Text             interface{} //
	PostId           interface{} //
	ParrentCommentId interface{} //
	CreatedBy        interface{} //
}
