package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:     title,
		Completed: false,
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
		(*todos)[index].CompletedAt = time.Now()
	} else {
		(*todos)[index].CompletedAt = time.Time{}
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
		status := "Incomplete"
		completedAt := "N/A"
		if todo.Completed {
			status = "Completed"
			completedAt = todo.CompletedAt.Format(time.RFC1123)
		}
		fmt.Printf("%d: %s\n\tStatus: %s\n\tCreated: %s\n\tCompleted: %s\n", i, todo.Title, status, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}
}
