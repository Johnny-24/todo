package utils

import (
	"fmt"
	"time"
	"os"
)
func SaveLog () {
	today := time.Now()
	str := "Последняя дата запуска " + today.Format("02.01.2006 15:04")
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()


}
