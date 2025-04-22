package handlers

import (
	"net/http"
	"path/filepath"
)

const webDir = "/Users/maks/Golang/go_final_project/web"

func ServeFile(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		http.ServeFile(res, req, filepath.Join(webDir, "index.html"))
		return
	}

	filePath := filepath.Join(webDir, req.URL.Path)

	if _, err := http.Dir(webDir).Open(req.URL.Path); err != nil {
		http.Error(res, "Файл не найден", http.StatusNotFound)
		return
	}

	http.ServeFile(res, req, filePath)
}
