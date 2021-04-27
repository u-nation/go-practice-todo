package output

import "time"

type (
	Todos struct {
		Title    string
		Deadline time.Time
	}

	Todo struct {
		Id       string
		Title    string
		isDone   bool
		Deadline time.Time
	}
)
