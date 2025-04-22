package api

import (
	"errors"
	"fmt"
	"go_final_project/pkg/db"
	"time"
)

func checkDate(task *db.Task) error {
	now := time.Now()
	layout := "20060102"

	// если дата не указана — берём текущую
	if task.Date == "" {
		task.Date = now.Format(layout)
	}

	t, err := time.Parse(layout, task.Date)
	if err != nil {
		return fmt.Errorf("Дата указана в неверном формате: %v", err)
	}

	// если есть repeat — проверим и пересчитаем дату
	if task.Repeat != "" {
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return errors.New("Неверное правило повторения: " + err.Error())
		}
		// если дата меньше сегодня — используем next
		if afterNow(now, t) {
			task.Date = next
		}
	} else {
		// если без repeat и дата в прошлом — ставим сегодняшнюю
		if afterNow(now, t) {
			task.Date = now.Format(layout)
		}
	}

	return nil
}
