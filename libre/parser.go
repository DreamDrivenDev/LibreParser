package libre

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func Parse(fileMath string) {
	absPath, err := filepath.Abs(fileMath)
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
}