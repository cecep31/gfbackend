package main

import (
	_ "gfbackend/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gfbackend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
