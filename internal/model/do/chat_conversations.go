// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatConversations is the golang structure of table chat_conversations for DAO operations like Where/Data.
type ChatConversations struct {
	g.Meta    `orm:"table:chat_conversations, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Title     interface{} //
	UserId    interface{} //
}
