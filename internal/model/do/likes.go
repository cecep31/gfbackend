// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Likes is the golang structure of table likes for DAO operations like Where/Data.
type Likes struct {
	g.Meta    `orm:"table:likes, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time //
	PostId    interface{} //
	UserId    interface{} //
}
