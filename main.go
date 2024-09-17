package main

import "fmt"

func main() {
	fmt.Print("Hello, World!")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	todos.add("Buy milk")
	todos.add("Buy eggs")
	todos.add("Buy bread")
	todos.add("Buy butter")
	todos.delete(2)
	todos.print()
	storage.Save(todos)

}
