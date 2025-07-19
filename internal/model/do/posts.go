// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Posts is the golang structure of table posts for DAO operations like Where/Data.
type Posts struct {
	g.Meta     `orm:"table:posts, do:true"`
	Id         interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
	Title      interface{} //
	CreatedBy  interface{} //
	Body       interface{} //
	Slug       interface{} //
	Createbyid interface{} //
	PhotoUrl   interface{} //
	Published  interface{} //
}
