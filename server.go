package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case "GET":
		fmt.Fprintf(w, "hello GET")
	case "POST":
		user, err := JsonDecoder(req.Body)

		if err != nil {
			fmt.Println(err.Error())
		}

		json.NewEncoder(w).Encode(user)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		fmt.Fprintf(w, "%v\n%v\n\n", name, headers)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8089", nil)
}
