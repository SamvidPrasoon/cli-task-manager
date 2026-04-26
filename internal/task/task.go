package task

import "time"

//enums
type Status string

const (
	StatusPending Status = "pending"
	StatusDone    Status = "done"
)

type Task struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func New(id int, title string) Task {
	return Task{
		Id:        id,
		Title:     title,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
}

func (t Task) IsDone() bool {
	return t.Status == StatusDone
}
func (t Task) IsPending() bool {
	return t.Status == StatusPending
}
func (t Task) MarkDone() Task {
	t.Status = StatusDone
	return t
}
