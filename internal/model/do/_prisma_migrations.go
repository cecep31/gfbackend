// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PrismaMigrations is the golang structure of table _prisma_migrations for DAO operations like Where/Data.
type PrismaMigrations struct {
	g.Meta            `orm:"table:_prisma_migrations, do:true"`
	Id                interface{} //
	Checksum          interface{} //
	FinishedAt        *gtime.Time //
	MigrationName     interface{} //
	Logs              interface{} //
	RolledBackAt      *gtime.Time //
	StartedAt         *gtime.Time //
	AppliedStepsCount interface{} //
}
