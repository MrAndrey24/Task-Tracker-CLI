package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title     string
	Completed bool
	CreatedAt time.Time
	Status    string
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:     title,
		Status:    "Progress",
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index >= 0 && index < len(*todos) {
		err := errors.New("Index out of range")
		fmt.Println(err.Error())
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	(*todos)[index].Completed = !(*todos)[index].Completed
	if (*todos)[index].Completed {
		(*todos)[index].Status = "Done"
	} else {
		(*todos)[index].Status = "Progress"
	}
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) print() {
	for i, todo := range *todos {
		if todo.Completed {
			todo.Status = "Progress"
		}
		fmt.Printf("ID: %d\n\tTitle: %s\n\tCreated: %s\n\tStatus: %s\n", i, todo.Title, todo.CreatedAt.Format(time.RFC1123), todo.Status)
	}
}
