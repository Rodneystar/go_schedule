package data

import (
	"bytes"
	"encoding/json"
	"os"
	"schedule/clock"
	"time"
)

type TimerActiveSpan struct {
	AtTime   clock.Clock   `json: "atTime"`
	Duration time.Duration `json: "duration"`
}

type DataAccess interface {
	GetAllTimers() ([]TimerActiveSpan, error)
	AddTimer(TimerActiveSpan) error
	RemoveByAtTime(clock.Clock) error
	DelAll() error
}

func NewDataAccess(f *os.File) DataAccess {
	fileaccess := &FileAccess{
		file: f,
	}
	fileaccess.Init(f)
	return fileaccess

}

type FileAccess struct {
	file *os.File
}

func (da *FileAccess) writeAll(bs []byte) error {
	da.file.Truncate(0)
	_, err := da.file.Write(bs)
	return err
}

func (da *FileAccess) RemoveByAtTime(t clock.Clock) error {
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

	err = da.writeAll(encoded)
	if err != nil {
		return err
	}
	return nil

}

func (da *FileAccess) Init(f *os.File) {
	da.file = f

}

func (da *FileAccess) DelAll() error {
	empty, err := toJSONArr(make([]TimerActiveSpan, 0))
	if err != nil {
		return err
	}
	da.file.Truncate(0)
	_, err = da.file.Write(empty)
	if err != nil {
		return err
	}
	return nil
}

func (da *FileAccess) readAll() ([]byte, error) {
	chunk := make([]byte, 64)
	buffer := make([]byte, 0)
	var err error
	done := false
	for i := 0; !done; {
		i, err = da.file.Read(chunk)
		if i == 0 || err != nil {
			done = true
		}
		buffer = append(buffer, chunk...)
	}
	return buffer, err
}

func (da *FileAccess) GetAllTimers() ([]TimerActiveSpan, error) {
	bytes.NewBuffer(nil)
	jsonIn, err := da.readAll()
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
	return da.writeAll(encodedTimers)
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
