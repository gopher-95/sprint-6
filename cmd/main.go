package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	//логи будем записывать в файл, создаем его только для записи
	logFile, err := os.OpenFile("logfile.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//создание логгера
	logger := log.New(logFile, "serv", 0755)

	//создание сервера с помощью функции из пакета server
	server := server.NewServer(logger)

	//запуск сервера
	err = server.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
