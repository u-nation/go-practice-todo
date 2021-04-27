package record

import "time"

type (
	TodoRecord struct {
		CommonRecordSoftDelete
		Title    string
		Deadline time.Time
		isDone   bool
	}
)
