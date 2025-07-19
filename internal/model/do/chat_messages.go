// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatMessages is the golang structure of table chat_messages for DAO operations like Where/Data.
type ChatMessages struct {
	g.Meta           `orm:"table:chat_messages, do:true"`
	Id               interface{} //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	ConversationId   interface{} //
	UserId           interface{} //
	Role             interface{} //
	Content          interface{} //
	Model            interface{} //
	PromptTokens     interface{} //
	CompletionTokens interface{} //
	TotalTokens      interface{} //
}
