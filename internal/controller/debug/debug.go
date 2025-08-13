package debug

import (
	"gfbackend/api/debug"
)

type ControllerV1 struct{}

func NewV1() debug.IDebugV1 {
	return &ControllerV1{}
}