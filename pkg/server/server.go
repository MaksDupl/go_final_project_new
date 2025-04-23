package server

import (
	"fmt"
	"go_final_project/pkg/api"
	"net/http"
)

func Run() error {
	port := 7540
	api.Init()

	http.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
