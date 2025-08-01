// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PostsToTagsDao is the data access object for the table posts_to_tags.
type PostsToTagsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PostsToTagsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PostsToTagsColumns defines and stores column names for the table posts_to_tags.
type PostsToTagsColumns struct {
	PostId string //
	TagId  string //
}

// postsToTagsColumns holds the columns for the table posts_to_tags.
var postsToTagsColumns = PostsToTagsColumns{
	PostId: "post_id",
	TagId:  "tag_id",
}

// NewPostsToTagsDao creates and returns a new DAO object for table data access.
func NewPostsToTagsDao(handlers ...gdb.ModelHandler) *PostsToTagsDao {
	return &PostsToTagsDao{
		group:    "default",
		table:    "posts_to_tags",
		columns:  postsToTagsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PostsToTagsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PostsToTagsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PostsToTagsDao) Columns() PostsToTagsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PostsToTagsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PostsToTagsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PostsToTagsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
