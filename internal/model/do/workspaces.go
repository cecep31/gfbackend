// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Workspaces is the golang structure of table workspaces for DAO operations like Where/Data.
type Workspaces struct {
	g.Meta      `orm:"table:workspaces, do:true"`
	Id          interface{} //
	Name        interface{} //
	Description interface{} //
	CreatedBy   interface{} //
	DeletedAt   *gtime.Time //
}
