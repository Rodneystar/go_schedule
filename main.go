package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// header := res.Header()
		type nameReq struct {
			Name string `json:"name"`
		}
		nameR := nameReq{}

		body := make([]byte, req.ContentLength)
		req.Body.Read(body)
		json.Unmarshal(body, &nameR)
		fmt.Printf("reqBody: %s\n", nameR.Name)
	})
	mux.HandleFunc("/mode", modeHandler)

	http.ListenAndServe("0.0.0.0:8080", mux)

}

func modeHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getMode(res, req)
	case http.MethodPut:
		// putMode(res, req)
	}
}

func getMode(res http.ResponseWriter, req *http.Request) {
	headers := res.Header()
	headers.Add("content-type", "application/json")
	res.WriteHeader(200)
	resBody := struct{ Mode bool }{Mode: true}
	json, _ := json.Marshal(resBody)
	res.Write(json)
}
