package workspaces

import (
	"gfbackend/api/workspaces"
)

type ControllerV1 struct{}

func NewV1() workspaces.IWorkspacesV1 {
	return &ControllerV1{}
}