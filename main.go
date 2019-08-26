package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"schedule/data"
	"schedule/physical"
	"schedule/service"
)

type modeResponse struct{ Mode service.TimerMode }

var serv *service.TimerService
var mux http.ServeMux
var LOGGER *log.Logger

type LoggingResponseWriter struct {
	http.ResponseWriter
	responseBody []byte
	status       int
}

func (w *LoggingResponseWriter) Response() []byte {
	return w.responseBody
}
func (w *LoggingResponseWriter) WriteHeader(st int) {
	w.status = st
	w.ResponseWriter.WriteHeader(st)
}
func (w *LoggingResponseWriter) Write(str []byte) (int, error) {
	w.responseBody = str
	return w.ResponseWriter.Write(str)
}

func main() {
	LOGGER = log.New(os.Stdout, "controller/main: ", 0)
	serv = &service.TimerService{}
	serv.Init(data.NewDataAccess(), physical.NewFakeSwitchable(false))

	mux := http.NewServeMux()

	mux.HandleFunc("/mode", loggingWrapper(http.HandlerFunc(modeHandler)))

	http.ListenAndServe("0.0.0.0:8080", mux)
}

func loggingWrapper(h http.Handler) http.HandlerFunc {

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
func modeHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getMode(res, req)
	case http.MethodPut:
		putMode(res, req)
	}
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
