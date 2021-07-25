package storage

import (
    "github.com/Pawilonek/nozbe-cli/tasks"
)

type storage interface {
    LoadTasks() (tasks.List, error)
    SaveTasks(tasks.List) error
}


