// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WorkspacesDao is the data access object for the table workspaces.
type WorkspacesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WorkspacesColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WorkspacesColumns defines and stores column names for the table workspaces.
type WorkspacesColumns struct {
	Id          string //
	Name        string //
	Description string //
	CreatedBy   string //
	DeletedAt   string //
}

// workspacesColumns holds the columns for the table workspaces.
var workspacesColumns = WorkspacesColumns{
	Id:          "id",
	Name:        "name",
	Description: "description",
	CreatedBy:   "created_by",
	DeletedAt:   "deleted_at",
}

// NewWorkspacesDao creates and returns a new DAO object for table data access.
func NewWorkspacesDao(handlers ...gdb.ModelHandler) *WorkspacesDao {
	return &WorkspacesDao{
		group:    "default",
		table:    "workspaces",
		columns:  workspacesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WorkspacesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WorkspacesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WorkspacesDao) Columns() WorkspacesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WorkspacesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WorkspacesDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *WorkspacesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
