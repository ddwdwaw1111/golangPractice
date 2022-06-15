package Taskstore

import (
	"time"
)

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type Taskstore struct {
	tasks  map[int]Task
	nextId int
}

func New() *Taskstore {
	ts := &Taskstore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

func (ts *Taskstore) CreateTask(text string, tags []string, duo time.Time) int {

	return 0
}
