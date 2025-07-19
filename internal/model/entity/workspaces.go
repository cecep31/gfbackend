// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Workspaces is the golang structure for table workspaces.
type Workspaces struct {
	Id          string      `json:"id"          orm:"id"          ` //
	Name        string      `json:"name"        orm:"name"        ` //
	Description string      `json:"description" orm:"description" ` //
	CreatedBy   string      `json:"createdBy"   orm:"created_by"  ` //
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  ` //
}
