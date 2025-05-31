package main

import "fmt"
func main () {
	list := []string{}

	greeting()
	enter(&list)

}

func enter (list *[]string) {
	for {
		var answear int
		fmt.Print("Выбарите действие: ")
		fmt.Scan(&answear)

		if answear == 1 {
			if len(*list) == 0 {
				fmt.Println("Список задач пуст")
			} else {
				for index, el := range *list {
					fmt.Println(index + 1, el)
				}
			}
		}

		if answear == 2 {
			var task string
			fmt.Println("Введи название задачи: ")
			fmt.Scan(&task)
			*list = append(*list, task)
		}

		if answear == 3 {
			if len(*list) == 0 {
				fmt.Println("Список задач пуст, удалять нечего!")
				continue
			}
			var id int
			fmt.Println("Введи номер задачи которую нужно удалить: ")
			fmt.Scan(&id)

			if id < 1 || id > len(*list) {
				fmt.Println("Неверный номер задачи!")
				continue
			}

			*list = append((*list)[:id-1], (*list)[id:]...)
			fmt.Println("Задача удалена!")
		}

		if answear == 4 {
			fmt.Println("Выход...")
			return
		}
	}
}

func greeting () {
	fmt.Println("")
	fmt.Println("<=== Todo List ===>")
	fmt.Println("")

	fmt.Println("1 - Список задач")
	fmt.Println("2 - Добавить задачу")
	fmt.Println("3 - Удалить задачу")
	fmt.Println("4 - Выход")
}
