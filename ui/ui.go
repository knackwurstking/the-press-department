package ui

import (
	"embed"
	"io/fs"
)

//go:embed public
var public embed.FS

func Public() fs.FS {
	f, err := fs.Sub(public, "public")
	if err != nil {
		panic(err)
	}

	return f
}
