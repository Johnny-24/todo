package todoList

import (
	"fmt"
	"github.com/fatih/color"
)

type Task struct {
    ID   int
    Text string
    Done bool
}

type List struct {
    tasks []Task
}

func NewList() *List {
    return &List{}
}

func (l *List) Create(text string) {
    newTask := Task{
        ID:   len(l.tasks) + 1,
        Text: text,
        Done: false,
    }
    l.tasks = append(l.tasks, newTask)
    color.Cyan("Задача добавлена")
}

func (l *List) Delete(index int) {
    l.tasks = append(l.tasks[:index-1], l.tasks[index:]...)
    color.Cyan("Задача удалена")
}

func (l *List) GetAll() {
    if len(l.tasks) == 0 {
        color.Red("Список задач пуст")
        return
    }
    for index, el := range l.tasks {
        status := "[]"
        if el.Done { status = "[X]" }
        fmt.Println(index + 1, status, el.Text)
    }

    fmt.Println("")
}

func (l *List) Hello(title bool) {
    if title {
        fmt.Println("")
        color.Blue("<=== Todo List ===>")
        fmt.Println("")
    }

    color.Yellow("---------------")
	color.Yellow("1 - Список задач")
	color.Yellow("2 - Добавить задачу")
	color.Yellow("3 - Удалить задачу")
	color.Yellow("4 - Выход")
	color.Yellow("")
}