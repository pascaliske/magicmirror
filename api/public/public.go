package public

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/pascaliske/magicmirror/logger"
)

//go:embed static/*
var staticFs embed.FS

func StaticFiles() http.FileSystem {
	staticFsSubtree, err := fs.Sub(staticFs, "static")

	if err != nil {
		logger.Error("Failed to get filesystem subtree for static files.")
		return nil
	}

	return http.FS(staticFsSubtree)
}
