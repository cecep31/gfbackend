// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Sessions is the golang structure of table sessions for DAO operations like Where/Data.
type Sessions struct {
	g.Meta    `orm:"table:sessions, do:true"`
	Token     interface{} //
	UserId    interface{} //
	CreatedAt *gtime.Time //
	UserAgent interface{} //
	ExpiresAt *gtime.Time //
}
