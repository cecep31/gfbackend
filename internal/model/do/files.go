// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Files is the golang structure of table files for DAO operations like Where/Data.
type Files struct {
	g.Meta    `orm:"table:files, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Name      interface{} //
	Path      interface{} //
	Size      interface{} //
	Type      interface{} //
	CreatedBy interface{} //
}
