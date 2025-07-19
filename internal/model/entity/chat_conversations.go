// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatConversations is the golang structure for table chat_conversations.
type ChatConversations struct {
	Id        string      `json:"id"        orm:"id"         ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
	Title     string      `json:"title"     orm:"title"      ` //
	UserId    string      `json:"userId"    orm:"user_id"    ` //
}
