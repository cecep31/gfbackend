package auth

import (
	"gfbackend/api/auth"
)

type ControllerV1 struct{}

func NewV1() auth.IAuthV1 {
	return &ControllerV1{}
}