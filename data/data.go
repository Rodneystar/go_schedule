package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type TimerEventDescription struct {
	AtTime time.Time `json: atTIme`
	State  bool      `json: state`
}

type DataAccess interface {
	GetAllTimers() ([]TimerEventDescription, error)
	AddTimer(TimerEventDescription) error
	DelAll() error
}

func NewDataAccess() DataAccess {
	return &FileAccess{
		filename: "timers.dat",
	}
}

type FileAccess struct {
	filename string
}

func (da *FileAccess) DelAll() error {
	da.createIfNotExist()
	empty, err := toJSONArr(make([]TimerEventDescription, 0))
	if err != nil {
		return err
	}
	ioutil.WriteFile(da.filename, empty, 0644)
	return nil
}
func (da *FileAccess) createIfNotExist() {
	if _, err := os.Stat(da.filename); os.IsNotExist(err) {
		os.Create(da.filename)
	}
}
func (da *FileAccess) GetAllTimers() ([]TimerEventDescription, error) {
	da.createIfNotExist()
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
