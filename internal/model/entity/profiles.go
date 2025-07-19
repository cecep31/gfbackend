// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Profiles is the golang structure for table profiles.
type Profiles struct {
	Id        int         `json:"id"        orm:"id"         ` //
	UserId    string      `json:"userId"    orm:"user_id"    ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	Bio       string      `json:"bio"       orm:"bio"        ` //
	Website   string      `json:"website"   orm:"website"    ` //
	Phone     string      `json:"phone"     orm:"phone"      ` //
	Location  string      `json:"location"  orm:"location"   ` //
}
