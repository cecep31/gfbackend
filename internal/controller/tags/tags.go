package tags

import (
	"gfbackend/api/tags"
)

type ControllerV1 struct{}

func NewV1() tags.ITagsV1 {
	return &ControllerV1{}
}