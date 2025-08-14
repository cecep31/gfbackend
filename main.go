package main

import (
	_ "gfbackend/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"log"

	"github.com/gogf/gf/v2/os/gctx"

	"gfbackend/internal/cmd"
	"gfbackend/internal/utility"
)

func main() {
	// Load environment files in order of precedence
	err := utility.LoadEnvFiles()
	if err != nil {
		log.Printf("Warning: Error loading .env files: %v", err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}
