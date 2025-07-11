package main

import (
	"fmt"
	"strconv"
	"todo/todoList"
	"todo/utils"
	"github.com/fatih/color"
)
func main () {
	utils.SaveLog()

	myList := todoList.NewList()
	myList.Hello(true)
	myList.ReadFromFile(myList)

	for {
		var answear int
		color.Magenta("Выбарите действие: ")
		fmt.Scan(&answear)

		switch answear {
		case 1:
			myList.GetAll()
			myList.Hello(false)

		case 2:
			var task string
			color.Magenta("Введи название задачи: ")
			fmt.Scan(&task)
			myList.Create(task)
			myList.Hello(false)

		case 3:
			var indexStr string
			color.Magenta("Введи номер задачи: ")
			fmt.Scan(&indexStr)
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				color.Red("Ошибка: нужно ввести число!")
					return
			}
			myList.Delete(index)
			myList.Hello(false)

		case 4:
			var indexStr string
			color.Magenta("Введи номер задачи: ")
			fmt.Scan(&indexStr)
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				color.Red("Ошибка: нужно ввести число!")
					return
			}
			myList.SelectTask(index)
			myList.Hello(false)

		case 5:
			myList.DeleteSelected()
			myList.Hello(false)

		case 6:
			var indexStr string
			color.Magenta("Введи номер задачи: ")
			fmt.Scan(&indexStr)
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				color.Red("Ошибка: нужно ввести число!")
					return
			}
			myList.Edit(index - 1)
			myList.Hello(false)

		case 7:
			myList.DelAll()
			myList.Hello(false)
		case 8:
			color.Blue("Выход...")
			return

		default:
			color.Red("Ошибка: нужно ввести число от 1 до 5!")
		}
	}
}