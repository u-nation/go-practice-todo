package model

import (
	"encoding/json"
	"time"

	"golang.org/x/xerrors"
)

type Deadline time.Time

func (d Deadline) Validate() error {
	if time.Now().After(time.Time(d)) {
		return xerrors.New("過去時間は設定できません")
	}
	return nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return xerrors.Errorf("failed to Deadline Unmarshal: %w", err)
	}

	tm, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return xerrors.Errorf("failed to Deadline Parse: %w", err)
	}
	*d = Deadline(tm.In(time.UTC))
	return nil
}

func (d Deadline) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(time.Time(d).Format(time.RFC3339))
	if err != nil {
		return nil, xerrors.Errorf("failed to Deadline Marshal: %w", err)
	}

	return bytes, nil
}
