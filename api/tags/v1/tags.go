package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateTagReq struct {
	g.Meta `path:"/tags" tags:"Tags" method:"post" summary:"Create new tag"`
	Name   string `json:"name" v:"required" description:"Tag name"`
}

type CreateTagRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Tag   `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type GetAllTagsReq struct {
	g.Meta `path:"/tags" tags:"Tags" method:"get" summary:"Get all tags"`
	Offset int `json:"offset" d:"0" description:"Number of records to skip"`
	Limit  int `json:"limit" d:"10" description:"Number of records to return"`
}

type GetAllTagsRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []Tag     `json:"data"`
	Meta    *MetaData `json:"meta,omitempty"`
}

type GetTagReq struct {
	g.Meta `path:"/tags/{id}" tags:"Tags" method:"get" summary:"Get tag by ID"`
	ID     string `json:"id" v:"required" description:"Tag ID"`
}

type GetTagRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Tag   `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type UpdateTagReq struct {
	g.Meta `path:"/tags/{id}" tags:"Tags" method:"put" summary:"Update tag"`
	ID     string `json:"id" v:"required" description:"Tag ID"`
	Name   string `json:"name" v:"required" description:"Tag name"`
}

type UpdateTagRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Tag   `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type DeleteTagReq struct {
	g.Meta `path:"/tags/{id}" tags:"Tags" method:"delete" summary:"Delete tag"`
	ID     string `json:"id" v:"required" description:"Tag ID"`
}

type DeleteTagRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Data models
type Tag struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type MetaData struct {
	TotalItems int `json:"total_items"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}