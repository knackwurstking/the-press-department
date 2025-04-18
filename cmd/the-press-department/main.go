package main

import (
	"log"
	"os"
	"the-press-department/ui"

	"github.com/hajimehoshi/ebiten/v2"
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
	if os.Getenv("GOOS") != "js" && os.Getenv("GOARCH") != "wasm" {
		e := echo.New()

		e.GET(ServerPath+"/*", echo.StaticDirectoryHandler(ui.Public(), false))

		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
			Output: os.Stderr,
		}))

		if err := e.Start(ServerAddr); err != nil {
			e.Logger.Fatal(err)
		}

		os.Exit(0)
	}

	ebiten.SetWindowSize(940, 470)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("The Press Department")

	if err := ebiten.RunGame(NewGame(DefaultScale * 1.5)); err != nil {
		log.Fatalf("Run game failed: %s", err.Error())
	}
}
