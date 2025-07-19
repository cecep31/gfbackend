// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Files is the golang structure for table files.
type Files struct {
	Id        string      `json:"id"        orm:"id"         ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
	Name      string      `json:"name"      orm:"name"       ` //
	Path      string      `json:"path"      orm:"path"       ` //
	Size      int         `json:"size"      orm:"size"       ` //
	Type      string      `json:"type"      orm:"type"       ` //
	CreatedBy string      `json:"createdBy" orm:"created_by" ` //
}
