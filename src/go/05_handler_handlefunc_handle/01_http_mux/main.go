//TODO: SOURCE
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	th := &timeHandler{format: time.RFC1123}
	mux.Handle("/time", th)
	mux.HandleFunc("/penis", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PENIS!"))
	})
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/", indexHandler)
	log.Println("Listening")
	http.ListenAndServe(":3000", mux)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello API")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Index")
}

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is : " + tm))
}
