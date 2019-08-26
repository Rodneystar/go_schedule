package main

import "net/http"

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
