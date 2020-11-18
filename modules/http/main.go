package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var s Server

type Server struct{}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.Marshal(struct {
		Method string
		URL    interface{}
		Header interface{}
		Body   interface{}
	}{
		Method: r.Method,
		URL:    r.URL,
		Header: r.Header,
		Body:   r.Form,
	})

	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}

func main() {
	err := http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatalln(err)
	}
}
