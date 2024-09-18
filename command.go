package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add            string
	Del            int
	Edit           string
	MarkInProgress int
	MarkInDone     int
	List           bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{
		Add:            "",
		Del:            -1,
		Edit:           "",
		MarkInProgress: -1,
		MarkInDone:     -1,
		List:           false,
	}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "update", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify todo by index to delete")
	flag.IntVar(&cf.MarkInProgress, "progress", -1, "Specify todo by index to mark in progress")
	flag.IntVar(&cf.MarkInDone, "done", -1, "Specify todo by index to mark in done")
	flag.BoolVar(&cf.List, "list", false, "List all to-dos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		if index, newTitle, err := parseEdit(cf.Edit); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			todos.edit(index, newTitle)
		}
	case cf.MarkInProgress != -1:
		todos.markProgress(cf.MarkInProgress)
	case cf.MarkInDone != -1:
		todos.markDone(cf.MarkInDone)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Invalid command")
	}
}

func parseEdit(edit string) (int, string, error) {
	parts := strings.SplitN(edit, ":", 2)
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("error: invalid format for edit. Please use index:new_title")
	}
	index, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", fmt.Errorf("error: invalid index for edit")
	}
	return index, parts[1], nil
}
