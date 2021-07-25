package storage

import (
    "encoding/json"
    "os"
    "io/ioutil"

    "github.com/Pawilonek/nozbe-cli/tasks"
)

type Disk struct {
    filename string
}

func NewDisk(filename string) Disk {
    return Disk{filename: filename}
}

func (d Disk) LoadTasks() (tasks.List, error) {
    _, err := os.Stat(d.filename)
    if os.IsNotExist(err) {
        // Not yet initialized sotrage, return empty list of tasks

        return tasks.NewList(), nil
    }

    data, err := ioutil.ReadFile(d.filename)
    if err != nil {
         return tasks.NewList(), err
    }

    var tasksArray []tasks.Task

    err = json.Unmarshal(data, &tasksArray)
    if err != nil {
         return tasks.NewList(), err
    }

    list := tasks.FromStorage(tasksArray)

    return list, nil
}

func (d Disk) SaveTasks(list tasks.List) error {
    b, err := json.Marshal(list.List())
	if err != nil {
		return err
	}

    err = ioutil.WriteFile(d.filename, b, 0644)
    if err != nil {
        return err
    }

    return nil
}

