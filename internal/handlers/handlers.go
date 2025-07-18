package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// первый обработчик записывает в тело ответа html форму
func FirstHandler(w http.ResponseWriter, r *http.Request) {
	//получаем абсолютный путь файла index.html
	filePath := "..\\index.html"

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		log.Fatal(err)
	}

	//читаем содержимое файла, передав абсолютный путь файла
	data, err := os.ReadFile(absPath)
	if err != nil {
		fmt.Println("не удалось открыть файл")
		log.Fatal(err)
	}

	//устанавливаем заголовок
	w.Header().Set("Content-Type", "text/html")

	//передаем содержимое в ответ сервера
	w.Write(data)
}

// второй обработчик достает переданный в html форму файл
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//парсим html форму, получаем файл и читаем из него данные (сохраняем в переменную dataFromFile)
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при загрузке", http.StatusInternalServerError)
	}

	defer file.Close()

	dataFromFile, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при загрузке", http.StatusInternalServerError)
	}

	//конвертация полученных данных в слово, либо в Морзе
	text := service.MorseOrWord(string(dataFromFile))

	//создаем локальный файл
	fileName := time.Now().UTC().Format("2006-01-02_15-04-05") + ".txt"

	createdFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//записываем в этот файл результат конвератции
	_, err = createdFile.Write([]byte(text))
	if err != nil {
		log.Fatal(err)
	}

	//закрываем файл
	defer createdFile.Close()

	//устанавливаем заголовок
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	//возвращаем результат конвератции строки
	w.Write([]byte(text))
}
