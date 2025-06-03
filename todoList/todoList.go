package todoList

import (
	"fmt"
	"todo/utils"

	"github.com/fatih/color"
)

  type Task struct {
    id   int
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
      id:   len(l.tasks) + 1,
      Text: text,
      Done: false,
    }
    l.tasks = append(l.tasks, newTask)
    color.Cyan("Задача добавлена")
  }

  func (l *List) Delete(index int) {
    utilsArray := utils.Array[Task]{}
    newArray := utilsArray.DelByIdx(l.tasks, index-1)
    l.tasks = newArray
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

  func (l *List) SelectTask(index int) {
    if index > len(l.tasks) {
      color.Red("Такой задачи нет")
      return
    }
    l.tasks[index-1].Done = !l.tasks[index-1].Done
  }

  func (l *List) DeleteSelected () {
    filteredTasks := make([]Task, 0, len(l.tasks))
    for _, task := range l.tasks {
      if !task.Done {
        filteredTasks = append(filteredTasks, task)
      }
    }
    l.tasks = filteredTasks
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
    color.Yellow("4 - Отметить задачу")
    color.Yellow("5 - Удалить все отмеченные задачи")
    color.Yellow("6 - Выход")
    color.Yellow("")
  }