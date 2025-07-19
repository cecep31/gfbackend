// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PostCommentsDao is the data access object for the table post_comments.
type PostCommentsDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  PostCommentsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// PostCommentsColumns defines and stores column names for the table post_comments.
type PostCommentsColumns struct {
	Id               string //
	CreatedAt        string //
	UpdatedAt        string //
	DeletedAt        string //
	Text             string //
	PostId           string //
	ParrentCommentId string //
	CreatedBy        string //
}

// postCommentsColumns holds the columns for the table post_comments.
var postCommentsColumns = PostCommentsColumns{
	Id:               "id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	Text:             "text",
	PostId:           "post_id",
	ParrentCommentId: "parrent_comment_id",
	CreatedBy:        "created_by",
}

// NewPostCommentsDao creates and returns a new DAO object for table data access.
func NewPostCommentsDao(handlers ...gdb.ModelHandler) *PostCommentsDao {
	return &PostCommentsDao{
		group:    "default",
		table:    "post_comments",
		columns:  postCommentsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PostCommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PostCommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PostCommentsDao) Columns() PostCommentsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PostCommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PostCommentsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PostCommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
