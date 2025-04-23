package main

import (
	"fmt"
	"log"

	"go_final_project/pkg/db"
	"go_final_project/pkg/server"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных: %v", err)
	}

	err = server.Run()
	if err != nil {
		fmt.Print("Ошибка запуска сервера")
	}
}
