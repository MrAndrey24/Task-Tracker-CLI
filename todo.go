package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Todo struct {
	Title     string
	Status    string
	Completed bool
	CreatedAt time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:     title,
		Status:    "Not Started",
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)

	fmt.Println("Task added successfully ID: " + strconv.Itoa(len(*todos)-1))
}

func (t *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("index out of range")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	fmt.Println("Task deleted successfully")
	return nil
}

func (todos *Todos) markDone(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	(*todos)[index].Completed = true
	(*todos)[index].Status = "Done"
	fmt.Println("Task marked as done successfully")
	return nil
}

func (todos *Todos) markProgress(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	(*todos)[index].Completed = false
	(*todos)[index].Status = "In Progress"
	fmt.Println("Task marked as in progress successfully")
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	(*todos)[index].Title = title
	fmt.Println("Task updated successfully")
	return nil
}

func (todos *Todos) printFiltered(filter func(todo Todo) bool) {
	for i, todo := range *todos {
		if filter(todo) {
			fmt.Printf("ID: %d\n\tTitle: %s\n\tCreated: %s\n\tStatus: %s\n", i, todo.Title, todo.CreatedAt.Format(time.RFC1123), todo.Status)
		}
	}
}

func (todos *Todos) printAll() {
	todos.printFiltered(func(todo Todo) bool {
		return true
	})
}

func (todos *Todos) printDone() {
	todos.printFiltered(func(todo Todo) bool {
		return todo.Completed
	})
}

func (todos *Todos) printNotStarted() {
	todos.printFiltered(func(todo Todo) bool {
		return !todo.Completed && todo.Status == "Not Started"
	})
}

func (todos *Todos) printInProgress() {
	todos.printFiltered(func(todo Todo) bool {
		return !todo.Completed && todo.Status == "In Progress"
	})
}
