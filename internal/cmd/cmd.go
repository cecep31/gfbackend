package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gfbackend/internal/controller/auth"
	"gfbackend/internal/controller/debug"
	"gfbackend/internal/controller/posts"
	"gfbackend/internal/controller/tags"
	"gfbackend/internal/controller/users"
	"gfbackend/internal/controller/workspaces"
	"gfbackend/internal/utility"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// Configure server with environment variables
			port := utility.GetEnvInt("SERVER_PORT", 8080)
			s.SetAddr(fmt.Sprintf(":%d", port))

			s.Group("/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					auth.NewV1(),
					users.NewV1(),
					posts.NewV1(),
					tags.NewV1(),
					workspaces.NewV1(),
					debug.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
