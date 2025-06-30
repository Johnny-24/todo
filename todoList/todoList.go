package todoList

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/utils"

	"github.com/fatih/color"
)

  type Task struct {
    Id   int `json:"id"`
    Text string `json:"text"`
    Done bool `json:"done"`
  }

  type List struct {
    Tasks []Task `json:"Tasks"`
  }

  func createJsonFile (list *List) {
    jsonFile, err := json.Marshal(&list)
    if err != nil {
      fmt.Println(err)
      return
    }

    file, err := os.Create("data.json")
    if err != nil {
      fmt.Println(err)
      return
    }

    _, err = file.Write(jsonFile)
    if err != nil {
      fmt.Println(err)
    }
    defer file.Close()
  }

  func NewList() *List {
    return &List{}
  }

  func (l *List) Create(text string) {
    newTask := Task{
      Id:   len(l.Tasks) + 1,
      Text: text,
      Done: false,
    }
    l.Tasks = append(l.Tasks, newTask)
    createJsonFile(l)
    color.Cyan("Успех!")
  }

  func (l *List) Delete(index int) {
    utilsArray := utils.Array[Task]{}
    newArray := utilsArray.DelByIdx(l.Tasks, index-1)
    l.Tasks = newArray
    createJsonFile(l)
    color.Cyan("Успех!")
  }

  func (l *List) GetAll() {
    if len(l.Tasks) == 0 {
      color.Red("Список задач пуст")
      return
    }
    fmt.Println("")
    for index, el := range l.Tasks {
      status := "[]"
      if el.Done { status = "[X]" }
      fmt.Println("№",index + 1, status, el.Text)
    }

    fmt.Println("")
  }

  func (l *List) SelectTask(index int) {
    if index > len(l.Tasks) {
      color.Red("Такой задачи нет")
      return
    }
    l.Tasks[index-1].Done = !l.Tasks[index-1].Done
    createJsonFile(l)
  }

  func (l *List) DeleteSelected() {
    filteredTasks := make([]Task, 0, len(l.Tasks))
    for _, task := range l.Tasks {
      if !task.Done {
        filteredTasks = append(filteredTasks, task)
      }
    }
    l.Tasks = filteredTasks
    createJsonFile(l)
  }

  func (l *List) Edit(idx int) {
    var name string
    color.Magenta("Введи новое название: ")
    fmt.Scan(&name)
    l.Tasks[idx].Text = name
    createJsonFile(l)
    color.Cyan("Успех!")
  }

  func (l *List) DelAll() {
    l.Tasks = make([]Task, 0)
    createJsonFile(l)
    color.Cyan("Успех!")
  }

  func (l *List) ReadFromFile (list *List) ([]byte, error) {
    file, err := os.ReadFile("data.json")
    fmt.Println("file", file)
    if err != nil {
      fmt.Println("err1", err)
      return nil, err
    }

    err = json.Unmarshal(file, &list)
    if err != nil {
      return nil, err
    }
    return file, nil
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
    color.Yellow("6 - Редактировать задачу")
    color.Yellow("7 - Удалить все задачи")
    color.Yellow("8 - Выход")
    color.Yellow("")
  }