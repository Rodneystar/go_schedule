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

	http.ListenAndServe("0.0.0.0:8080", mux)

}
