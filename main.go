package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"schedule/clock"
	"schedule/data"
	"schedule/physical"
	"schedule/service"
	"time"
)

type modeResponse struct{ Mode service.TimerMode }

var serv *service.TimerService
var mux http.ServeMux
var LOGGER *log.Logger
var storageFile os.File

func main() {
	LOGGER = log.New(os.Stdout, "controller/main: ", 0)
	serv = &service.TimerService{}
	storageFile, err := openFile("timer.dat")
	if err != nil {
		panic(err)
	}
	serv.Init(data.NewDataAccess(storageFile), physical.NewFakeSwitchable(false))

	mux := http.NewServeMux()

	mux.HandleFunc("/mode", loggingHandler(modeHandler()))
	serv.AddTimer(*clock.NewClock(time.Now()), 2*time.Hour)
	http.ListenAndServe("0.0.0.0:8080", mux)
}

func loggingHandler(h http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		loggedRes := &LoggingResponseWriter{
			ResponseWriter: res,
			responseBody:   make([]byte, 0),
			status:         0,
		}
		LOGGER.Printf("req received: %v, %v", req.Method, req.URL)
		h.ServeHTTP(loggedRes, req)

		LOGGER.Printf("res sent: %d, %s", loggedRes.status, loggedRes.Response())
	})
}
func modeHandler() http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			getMode(res, req)
		case http.MethodPut:
			putMode(res, req)
		}
	})
}

func putMode(res http.ResponseWriter, req *http.Request) {
	writeHeaders(res)
	jsonBody := make([]byte, req.ContentLength)
	req.Body.Read(jsonBody)
	reqBody := &modeResponse{}
	err := json.Unmarshal(jsonBody, reqBody)

	if err != nil {
		fmt.Println("we failed in putMode")
	}
	serv.SetMode(reqBody.Mode)
	res.Write(jsonBody)
}

func getMode(res http.ResponseWriter, req *http.Request) {
	writeHeaders(res)
	currMode := serv.GetMode()
	resBody := &modeResponse{Mode: currMode}
	jsonBody, _ := json.Marshal(resBody)
	res.Write(jsonBody)
}

func writeHeaders(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("content-type", "application/json")
	res.WriteHeader(200)
}

func openFile(name string) (*os.File, error) {
	dir := fmt.Sprintf("%s\\.heatingTimer", os.Getenv("HOME"))
	os.Mkdir(dir, os.ModeDir)
	filepath := fmt.Sprintf("%s\\%s", dir, name)
	return os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
}
