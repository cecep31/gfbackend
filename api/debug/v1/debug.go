package v1

import "github.com/gogf/gf/v2/frame/g"

type GetProfileReq struct {
	g.Meta `path:"/debug/pprof/{profile}" tags:"Debug" method:"get" summary:"Performance profiling"`
	Profile string `json:"profile" v:"required|in:goroutine,heap,threadcreate,block,mutex,profile,trace" description:"Profile type"`
}

type GetProfileRes struct {
	// This will return raw profiling data, handled directly by the controller
	// The response format depends on the profile type requested
}