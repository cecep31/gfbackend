package workspaces

import (
	"context"
	v1 "gfbackend/api/workspaces/v1"
)

type IWorkspacesV1 interface {
	// Workspace management
	CreateWorkspace(ctx context.Context, req *v1.CreateWorkspaceReq) (res *v1.CreateWorkspaceRes, err error)
	GetAllWorkspaces(ctx context.Context, req *v1.GetAllWorkspacesReq) (res *v1.GetAllWorkspacesRes, err error)
	GetUserWorkspaces(ctx context.Context, req *v1.GetUserWorkspacesReq) (res *v1.GetUserWorkspacesRes, err error)
	GetWorkspace(ctx context.Context, req *v1.GetWorkspaceReq) (res *v1.GetWorkspaceRes, err error)
	UpdateWorkspace(ctx context.Context, req *v1.UpdateWorkspaceReq) (res *v1.UpdateWorkspaceRes, err error)
	DeleteWorkspace(ctx context.Context, req *v1.DeleteWorkspaceReq) (res *v1.DeleteWorkspaceRes, err error)

	// Workspace member management
	AddWorkspaceMember(ctx context.Context, req *v1.AddWorkspaceMemberReq) (res *v1.AddWorkspaceMemberRes, err error)
	GetWorkspaceMembers(ctx context.Context, req *v1.GetWorkspaceMembersReq) (res *v1.GetWorkspaceMembersRes, err error)
	UpdateMemberRole(ctx context.Context, req *v1.UpdateMemberRoleReq) (res *v1.UpdateMemberRoleRes, err error)
	RemoveWorkspaceMember(ctx context.Context, req *v1.RemoveWorkspaceMemberReq) (res *v1.RemoveWorkspaceMemberRes, err error)

	// Page management
	CreatePage(ctx context.Context, req *v1.CreatePageReq) (res *v1.CreatePageRes, err error)
	GetPage(ctx context.Context, req *v1.GetPageReq) (res *v1.GetPageRes, err error)
	UpdatePage(ctx context.Context, req *v1.UpdatePageReq) (res *v1.UpdatePageRes, err error)
	DeletePage(ctx context.Context, req *v1.DeletePageReq) (res *v1.DeletePageRes, err error)
	GetWorkspacePages(ctx context.Context, req *v1.GetWorkspacePagesReq) (res *v1.GetWorkspacePagesRes, err error)
	GetChildPages(ctx context.Context, req *v1.GetChildPagesReq) (res *v1.GetChildPagesRes, err error)
}