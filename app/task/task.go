package task

import "time"

// task handles task operation

type cattype int

const (
	CategoryEliminate cattype = iota
	CategoryDelegate
	CategoryDoYourself
)

type valtype int

const (
	ValueImportant valtype = iota
	ValueUrgent
	ValueSignificant
)

type progtype int

const (
	TaskUnstarted progtype = iota
	TaskStarted
	TaskDoing
	TaskDone
	TaskObservation
)

type Task struct {
	id         int64
	title      string
	descriptio string
	category   cattype
	value      valtype
	startTime  time.Time
	endTime    time.Time
	Note       string
	progress   progtype
}

func (t *Task) ShowAll() []*Task {
	return nil
}

func (t *Task) Add(task *Task) error {
	return nil
}

func (t *Task) Update(id int64) error {
	return nil
}

func (t *Task) Remove(id int64) error {
	return nil
}
