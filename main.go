package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	odtFile := "libreEx.odt"
	txtFile := "document.txt"

	// Получение абсолютного пути к ODT файлу
	absPath, err := filepath.Abs(odtFile)
	if err != nil {
		fmt.Println("Ошибка получения абсолютного пути:", err)
		return
	}

	// Конвертация ODT файла в текстовый формат с помощью LibreOffice
	cmd := exec.Command("libreoffice", "--headless", "--convert-to", "txt:Text", absPath)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при конвертации файла:", err)
		return
	}

	// Построение пути к новому txt файлу
	txtFilePath := filepath.Join(filepath.Dir(absPath), txtFile)

	// Чтение сконвертированного текстового файла
	data, err := os.ReadFile(txtFilePath)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Вывод содержимого файла
	fmt.Println(string(data))
}