package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	y1, m1, d1 := date.Date()
	y2, m2, d2 := now.Date()
	return time.Date(y1, m1, d1, 0, 0, 0, 0, time.UTC).After(time.Date(y2, m2, d2, 0, 0, 0, 0, time.UTC))
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", errors.New("invalid dstart format")
	}

	parts := strings.Split(repeat, " ")
	if len(parts) == 0 {
		return "", errors.New("invalid repeat format")
	}

	switch parts[0] {
	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(dateFormat), nil

	case "d":
		if len(parts) < 2 {
			return "", errors.New("day interval is missing")
		}
		interval, err := strconv.Atoi(parts[1])
		if err != nil || interval <= 0 || interval > 366 {
			return "", errors.New("invalid day interval")
		}
		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(dateFormat), nil

	default:
		return "", errors.New("unsupported repeat format")
	}
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.FormValue("now")
	dstart := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var now time.Time
	var err error

	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFormat, nowStr)
		if err != nil {
			http.Error(w, "invalid now format", http.StatusBadRequest)
			return
		}
	}

	next, err := NextDate(now, dstart, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Возвращаем просто строку
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(next))
}
