package utils

import (
	"fmt"
	"os"
	"time"
)
func SaveLog () {
	filePath := "log.txt"
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
