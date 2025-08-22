package public

import (
	"embed"
	"io/fs"
)

//go:embed assets/*
var content embed.FS

var Content = content

func SubDir(name string) (fs.FS, error) {
	return fs.Sub(content, name)
}
