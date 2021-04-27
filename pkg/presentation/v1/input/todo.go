package input

import "time"

type (
	TodoPostParameter struct {
		Title    string
		Deadline time.Time
	}

	TodoPutParameter struct {
		TodoPostParameter
		isDone bool
	}
)
