// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Profiles is the golang structure of table profiles for DAO operations like Where/Data.
type Profiles struct {
	g.Meta    `orm:"table:profiles, do:true"`
	Id        interface{} //
	UserId    interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Bio       interface{} //
	Website   interface{} //
	Phone     interface{} //
	Location  interface{} //
}
