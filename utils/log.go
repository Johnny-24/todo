package utils

import (
	"fmt"
	"os"
	"time"
	"path/filepath"
)
func SaveLog () {
	filePath := "logs/log.txt"
	dirPath := filepath.Dir(filePath)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Println("Ошибка создания директории:", err)
			return
		}
	}

	var oldData string
	fileExist := true
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fileExist = false
	}

	if fileExist {
		data, _ := os.ReadFile(filePath)
		oldData = string(data)
		fileSize := fileInfo.Size()
		fmt.Println(fileSize > 248)
		if fileSize > 248 {
			oldData = ""
			err := os.Remove(filePath)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	today := time.Now()
	str := "Программа была запущена " + today.Format("02.01.2006 15:04") + "\n" + oldData

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}
