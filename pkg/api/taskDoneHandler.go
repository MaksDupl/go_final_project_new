package api

import (
	"go_final_project/pkg/db"
	"net/http"
	"time"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	// Получаем задачу из базы
	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Задача не найдена"})
		return
	}

	if task.Repeat == "" {
		// Одноразовая задача — удаляем
		if err := db.DeleteTask(id); err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка удаления: " + err.Error()})
			return
		}
	} else {
		// Периодическая задача — обновляем дату
		now := time.Now()
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка расчёта следующей даты: " + err.Error()})
			return
		}

		if err := db.UpdateDate(next, id); err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка обновления даты: " + err.Error()})
			return
		}
	}

	writeJSON(w, map[string]string{})
}
