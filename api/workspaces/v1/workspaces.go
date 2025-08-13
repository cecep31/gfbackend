package v1

import "github.com/gogf/gf/v2/frame/g"

// Workspace management endpoints
type CreateWorkspaceReq struct {
	g.Meta      `path:"/workspaces" tags:"Workspaces" method:"post" summary:"Create new workspace"`
	Name        string `json:"name" v:"required" description:"Workspace name"`
	Description string `json:"description" description:"Workspace description"`
}

type CreateWorkspaceRes struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *Workspace `json:"data,omitempty"`
	Error   string     `json:"error,omitempty"`
}

type GetAllWorkspacesReq struct {
	g.Meta `path:"/workspaces" tags:"Workspaces" method:"get" summary:"Get all workspaces"`
	Offset int `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int `json:"limit" d:"10" description:"Number of records to return"`
}

type GetAllWorkspacesRes struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    []Workspace `json:"data"`
	Meta    *MetaData   `json:"meta,omitempty"`
}

type GetUserWorkspacesReq struct {
	g.Meta `path:"/workspaces/user" tags:"Workspaces" method:"get" summary:"Get user workspaces"`
	Offset int `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int `json:"limit" d:"10" description:"Number of records to return"`
}

type GetUserWorkspacesRes struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    []Workspace `json:"data"`
	Meta    *MetaData   `json:"meta,omitempty"`
}

type GetWorkspaceReq struct {
	g.Meta `path:"/workspaces/{id}" tags:"Workspaces" method:"get" summary:"Get workspace by ID"`
	ID     string `json:"id" v:"required" description:"Workspace ID"`
}

type GetWorkspaceRes struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *Workspace `json:"data,omitempty"`
	Error   string     `json:"error,omitempty"`
}

type UpdateWorkspaceReq struct {
	g.Meta      `path:"/workspaces/{id}" tags:"Workspaces" method:"put" summary:"Update workspace"`
	ID          string `json:"id" v:"required" description:"Workspace ID"`
	Name        string `json:"name" description:"Workspace name"`
	Description string `json:"description" description:"Workspace description"`
}

type UpdateWorkspaceRes struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    *Workspace `json:"data,omitempty"`
	Error   string     `json:"error,omitempty"`
}

type DeleteWorkspaceReq struct {
	g.Meta `path:"/workspaces/{id}" tags:"Workspaces" method:"delete" summary:"Delete workspace"`
	ID     string `json:"id" v:"required" description:"Workspace ID"`
}

type DeleteWorkspaceRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Workspace member management endpoints
type AddWorkspaceMemberReq struct {
	g.Meta      `path:"/workspaces/{id}/members" tags:"Workspace Members" method:"post" summary:"Add workspace member"`
	ID          string `json:"id" v:"required" description:"Workspace ID"`
	UserID      string `json:"user_id" v:"required" description:"User ID to add"`
	Role        string `json:"role" v:"required|in:admin,member,viewer" description:"Member role"`
	Permissions string `json:"permissions" description:"Member permissions"`
}

type AddWorkspaceMemberRes struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    *WorkspaceMember  `json:"data,omitempty"`
	Error   string            `json:"error,omitempty"`
}

type GetWorkspaceMembersReq struct {
	g.Meta `path:"/workspaces/{id}/members" tags:"Workspace Members" method:"get" summary:"Get workspace members"`
	ID     string `json:"id" v:"required" description:"Workspace ID"`
	Offset int    `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int    `json:"limit" d:"10" description:"Number of records to return"`
}

type GetWorkspaceMembersRes struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    []WorkspaceMember  `json:"data"`
	Meta    *MetaData          `json:"meta,omitempty"`
}

type UpdateMemberRoleReq struct {
	g.Meta      `path:"/workspaces/{id}/members/{user_id}" tags:"Workspace Members" method:"put" summary:"Update member role"`
	ID          string `json:"id" v:"required" description:"Workspace ID"`
	UserID      string `json:"user_id" v:"required" description:"User ID"`
	Role        string `json:"role" v:"required|in:admin,member,viewer" description:"New member role"`
	Permissions string `json:"permissions" description:"New member permissions"`
}

type UpdateMemberRoleRes struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    *WorkspaceMember `json:"data,omitempty"`
	Error   string           `json:"error,omitempty"`
}

type RemoveWorkspaceMemberReq struct {
	g.Meta `path:"/workspaces/{id}/members/{user_id}" tags:"Workspace Members" method:"delete" summary:"Remove workspace member"`
	ID     string `json:"id" v:"required" description:"Workspace ID"`
	UserID string `json:"user_id" v:"required" description:"User ID to remove"`
}

type RemoveWorkspaceMemberRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Page management endpoints
type CreatePageReq struct {
	g.Meta      `path:"/workspaces/{workspace_id}/pages" tags:"Pages" method:"post" summary:"Create new page"`
	WorkspaceID string `json:"workspace_id" v:"required" description:"Workspace ID"`
	Title       string `json:"title" v:"required" description:"Page title"`
	Content     string `json:"content" description:"Page content"`
	ParentID    string `json:"parent_id" description:"Parent page ID"`
}

type CreatePageRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Page  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type GetPageReq struct {
	g.Meta      `path:"/workspaces/{workspace_id}/pages/{id}" tags:"Pages" method:"get" summary:"Get page by ID"`
	WorkspaceID string `json:"workspace_id" v:"required" description:"Workspace ID"`
	ID          string `json:"id" v:"required" description:"Page ID"`
}

type GetPageRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Page  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type UpdatePageReq struct {
	g.Meta      `path:"/workspaces/{workspace_id}/pages/{id}" tags:"Pages" method:"put" summary:"Update page"`
	WorkspaceID string `json:"workspace_id" v:"required" description:"Workspace ID"`
	ID          string `json:"id" v:"required" description:"Page ID"`
	Title       string `json:"title" description:"Page title"`
	Content     string `json:"content" description:"Page content"`
	ParentID    string `json:"parent_id" description:"Parent page ID"`
}

type UpdatePageRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Page  `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type DeletePageReq struct {
	g.Meta      `path:"/workspaces/{workspace_id}/pages/{id}" tags:"Pages" method:"delete" summary:"Delete page"`
	WorkspaceID string `json:"workspace_id" v:"required" description:"Workspace ID"`
	ID          string `json:"id" v:"required" description:"Page ID"`
}

type DeletePageRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type GetWorkspacePagesReq struct {
	g.Meta      `path:"/workspaces/{workspace_id}/pages" tags:"Pages" method:"get" summary:"Get workspace pages"`
	WorkspaceID string `json:"workspace_id" v:"required" description:"Workspace ID"`
	Offset      int    `json:"offset" d:"0" description:"Number of records to skip"`
	Limit       int    `json:"limit" d:"10" description:"Number of records to return"`
}

type GetWorkspacePagesRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Page    `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

type GetChildPagesReq struct {
	g.Meta      `path:"/workspaces/{workspace_id}/pages/{parent_id}/children" tags:"Pages" method:"get" summary:"Get child pages"`
	WorkspaceID string `json:"workspace_id" v:"required" description:"Workspace ID"`
	ParentID    string `json:"parent_id" v:"required" description:"Parent page ID"`
	Offset      int    `json:"offset" d:"0" description:"Number of records to skip"`
	Limit       int    `json:"limit" d:"10" description:"Number of records to return"`
}

type GetChildPagesRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Page    `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

// Data models
type Workspace struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	OwnerID     string `json:"owner_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type WorkspaceMember struct {
	ID          string `json:"id"`
	WorkspaceID string `json:"workspace_id"`
	UserID      string `json:"user_id"`
	Role        string `json:"role"`
	Permissions string `json:"permissions,omitempty"`
	User        *User  `json:"user,omitempty"`
	JoinedAt    string `json:"joined_at"`
}

type Page struct {
	ID          string `json:"id"`
	WorkspaceID string `json:"workspace_id"`
	Title       string `json:"title"`
	Content     string `json:"content,omitempty"`
	ParentID    string `json:"parent_id,omitempty"`
	AuthorID    string `json:"author_id"`
	Author      *User  `json:"author,omitempty"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type MetaData struct {
	TotalItems int `json:"total_items"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}