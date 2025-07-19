// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PostsToTags is the golang structure of table posts_to_tags for DAO operations like Where/Data.
type PostsToTags struct {
	g.Meta `orm:"table:posts_to_tags, do:true"`
	PostId interface{} //
	TagId  interface{} //
}
