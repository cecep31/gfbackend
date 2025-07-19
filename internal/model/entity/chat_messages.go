// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatMessages is the golang structure for table chat_messages.
type ChatMessages struct {
	Id               string      `json:"id"               orm:"id"                ` //
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        ` //
	ConversationId   string      `json:"conversationId"   orm:"conversation_id"   ` //
	UserId           string      `json:"userId"           orm:"user_id"           ` //
	Role             string      `json:"role"             orm:"role"              ` //
	Content          string      `json:"content"          orm:"content"           ` //
	Model            string      `json:"model"            orm:"model"             ` //
	PromptTokens     int         `json:"promptTokens"     orm:"prompt_tokens"     ` //
	CompletionTokens int         `json:"completionTokens" orm:"completion_tokens" ` //
	TotalTokens      int         `json:"totalTokens"      orm:"total_tokens"      ` //
}
