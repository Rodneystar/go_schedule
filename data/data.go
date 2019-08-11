package data

import (
	"encoding/json"
	"time"
)

type TimerEventDescription struct {
	AtTime time.Time `json: atTIme`
	State  bool      `json: state`
}

type DataAccess interface {
	GetAllTimers() []TimerEventDescription
	AddTimer(TimerEventDescription) bool
}

func toJson(event TimerEventDescription) ([]byte, error) {
	bytes, err := json.MarshalIndent(event, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
