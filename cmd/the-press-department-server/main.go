package main

import (
	"os"
	"the-press-department/ui"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	ServerAddr = os.Getenv("THEPRESSDEPARTMENT_SERVER_ADDR")
	ServerPath = os.Getenv("THEPRESSDEPARTMENT_SERVER_PATH")
)

func init() {
	if ServerAddr == "" {
		panic("Environment variable missing: THEPRESSDEPARTMENT_SERVER_ADDR")
	}
}

func main() {
	e := echo.New()

	e.GET(ServerPath+"/*", echo.StaticDirectoryHandler(ui.Public(), false))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: os.Stderr,
	}))

	if err := e.Start(ServerAddr); err != nil {
		e.Logger.Fatal(err)
	}
}
