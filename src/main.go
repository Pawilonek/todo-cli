package main

import (
	"fmt"

	"github.com/Pawilonek/nozbe-cli/tasks"
)

func main() {

	list := tasks.NewList()

	list.Add("test 1")
	list.Add("A small second task")
	list.Add("A third one!")
	list.Add("Why i'm counting this?")
	list.Add("Hello!")

	list.ToggleDone(1)
	list.ToggleDone(0)
	list.ToggleDone(1)

	list.ToggleDone(3)

	taskList := list.List()
	var doneCharacter string
	for i := 0; i < len(taskList); i++ {
		if taskList[i].Done {
			doneCharacter = "\033[1;32m☑\033[0m"
		} else {
			doneCharacter = "\033[1;31m☐\033[0m"
		}

		fmt.Printf("%s %s \n", doneCharacter, taskList[i].Name)
	}
}
