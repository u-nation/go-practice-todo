package output

import "gopkg.in/guregu/null.v4"

type Error struct {
	Code       int
	Link       string
	Message    string
	Validation null.String
}
