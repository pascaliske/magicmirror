package public

import (
	"embed"
	"io/fs"

	"github.com/pascaliske/magicmirror/logger"
)

//go:embed static/*
var staticFs embed.FS

func StaticFiles() fs.FS {
	staticFsSubtree, err := fs.Sub(staticFs, "static")

	if err != nil {
		logger.Error("Failed to get filesystem subtree for static files.")
		return nil
	}

	return staticFsSubtree
}
