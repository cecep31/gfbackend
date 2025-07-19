// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Sessions is the golang structure for table sessions.
type Sessions struct {
	Token     string      `json:"token"     orm:"token"      ` //
	UserId    string      `json:"userId"    orm:"user_id"    ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UserAgent string      `json:"userAgent" orm:"user_agent" ` //
	ExpiresAt *gtime.Time `json:"expiresAt" orm:"expires_at" ` //
}
