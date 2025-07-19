// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatMessagesDao is the data access object for the table chat_messages.
type ChatMessagesDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  ChatMessagesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// ChatMessagesColumns defines and stores column names for the table chat_messages.
type ChatMessagesColumns struct {
	Id               string //
	CreatedAt        string //
	UpdatedAt        string //
	ConversationId   string //
	UserId           string //
	Role             string //
	Content          string //
	Model            string //
	PromptTokens     string //
	CompletionTokens string //
	TotalTokens      string //
}

// chatMessagesColumns holds the columns for the table chat_messages.
var chatMessagesColumns = ChatMessagesColumns{
	Id:               "id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	ConversationId:   "conversation_id",
	UserId:           "user_id",
	Role:             "role",
	Content:          "content",
	Model:            "model",
	PromptTokens:     "prompt_tokens",
	CompletionTokens: "completion_tokens",
	TotalTokens:      "total_tokens",
}

// NewChatMessagesDao creates and returns a new DAO object for table data access.
func NewChatMessagesDao(handlers ...gdb.ModelHandler) *ChatMessagesDao {
	return &ChatMessagesDao{
		group:    "default",
		table:    "chat_messages",
		columns:  chatMessagesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ChatMessagesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ChatMessagesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ChatMessagesDao) Columns() ChatMessagesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ChatMessagesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ChatMessagesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ChatMessagesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
