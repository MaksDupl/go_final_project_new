package api

import (
	"go_final_project/pkg/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка получения задач: " + err.Error()})
		return
	}

	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}
