package data

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type TimerEventDescription struct {
	AtTime time.Time `json: atTIme`
	State  bool      `json: state`
}

type DataAccess interface {
	GetAllTimers() ([]TimerEventDescription, error)
	AddTimer(TimerEventDescription) error
}

func NewDataAccess() DataAccess {
	return &FileAccess{
		filename: "timers.dat",
	}
}

type FileAccess struct {
	filename string
}

func (da *FileAccess) GetAllTimers() ([]TimerEventDescription, error) {
	jsonIn, err := ioutil.ReadFile(da.filename)
	if err != nil {
		return nil, err
	}
	return fromJSONArr(jsonIn), nil
}

func (da *FileAccess) AddTimer(t TimerEventDescription) error {
	currTimers, err := da.GetAllTimers()
	if err != nil {
		return err
	}

	currTimers = append(currTimers, t)
	encodedTimers, err := toJSONArr(currTimers)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(da.filename, encodedTimers, 0644)
}

func toJSON(event TimerEventDescription) ([]byte, error) {
	bytes, err := json.MarshalIndent(event, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func toJSONArr(event []TimerEventDescription) ([]byte, error) {
	bytes, err := json.MarshalIndent(event, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func fromJSONArr(jsonArr []byte) []TimerEventDescription {
	var timerEvent []TimerEventDescription
	json.Unmarshal(jsonArr, &timerEvent)
	return timerEvent
}
