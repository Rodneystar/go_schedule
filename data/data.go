package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type TimerActiveSpan struct {
	AtTime time.Time 		`json: "atTime"`
	Duration time.Duration 	`json: "duration"`
}

type DataAccess interface {
	GetAllTimers() ([]TimerActiveSpan, error)
	AddTimer(TimerActiveSpan) error
	RemoveByAtTime( time.Time) error
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


func(da *FileAccess) RemoveByAtTime(t time.Time) error {
	timers, _ := da.GetAllTimers()
	var tIdx int
	for i, e := range timers {
		if e.AtTime.Equal(t) {
			tIdx = i
		}
	}

	firstPart := timers[0:tIdx]
	secondPart := timers[tIdx+1:]

	encoded, err := toJSONArr(append(firstPart, secondPart...))
	if err != nil {
		return err
	}

	err  = ioutil.WriteFile(da.filename, encoded, 0644)
	if err != nil {
		return err
	}
	return nil

}

func(da *FileAccess) Init(filename string) {
	da.filename = filename
	err := da.createIfNotExist()
	if err != nil {
		log.Fatalf("%s -- error creating file if doesnt exist: %s", time.Now(), err)
	}
}

func (da *FileAccess) DelAll() error {
	empty, err := toJSONArr(make([]TimerActiveSpan, 0))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(da.filename, empty, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (da *FileAccess) createIfNotExist() error {
	if _, err := os.Stat(da.filename); os.IsNotExist(err) {
		file, err := os.Create(da.filename)
		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (da *FileAccess) GetAllTimers() ([]TimerActiveSpan, error) {
	jsonIn, err := ioutil.ReadFile(da.filename)
	if err != nil {
		return nil, err
	}
	return fromJSONArr(jsonIn), nil
}

func (da *FileAccess) AddTimer(t TimerActiveSpan) error {
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

func toJSON(event TimerActiveSpan) ([]byte, error) {
	bytes, err := json.MarshalIndent(event, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func toJSONArr(event []TimerActiveSpan) ([]byte, error) {
	bytes, err := json.MarshalIndent(event, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func fromJSONArr(jsonArr []byte) []TimerActiveSpan {
	var timerEvent []TimerActiveSpan
	json.Unmarshal(jsonArr, &timerEvent)
	return timerEvent
}
