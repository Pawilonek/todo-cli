package tasks

type List struct {
	tasks []Task
}

func NewList() List {
	return List{
		tasks: []Task{},
	}
}

func (list *List) Add(name string) {
	list.tasks = append(list.tasks, Task{
		Name: name,
	})
}

func (list *List) ToggleDone(taskId int) {

	list.tasks[taskId].Done = !list.tasks[taskId].Done
}

func (list List) List() []Task {
	return list.tasks
}
