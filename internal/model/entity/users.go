// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id           string      `json:"id"           orm:"id"             ` //
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     ` //
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"     ` //
	FirstName    string      `json:"firstName"    orm:"first_name"     ` //
	LastName     string      `json:"lastName"     orm:"last_name"      ` //
	Email        string      `json:"email"        orm:"email"          ` //
	Password     string      `json:"password"     orm:"password"       ` //
	Image        string      `json:"image"        orm:"image"          ` //
	IsSuperAdmin bool        `json:"isSuperAdmin" orm:"is_super_admin" ` //
	Username     string      `json:"username"     orm:"username"       ` //
	GithubId     int64       `json:"githubId"     orm:"github_id"      ` //
}
