package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"strings"
	"time"
)

func main() {
	reader := newAlphaReader(strings.NewReader("Hello! It's 9 am maMFF()(§)NFuck3r! What'ssu uuuup b111444tvch!"))
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				//TODO Handle remaining Bytes
				fmt.Println(string(p[:n]))
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println((string(p[:n])))
	}
	fmt.Println()
}

type alphaReader struct {
	reader io.Reader
	cur    int
	src    int
}

// Custom Read using io.Reader
func (a *alphaReader) Read(p []byte) (int, error) {

	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}

	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

func alpha(r byte) byte {
	if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
		return r
	}
	return 0
}

func newAlphaReader(reader io.Reader) *alphaReader {
	return &alphaReader{reader: reader}
}

func stringReaderExample() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 8)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				//TODO Handle remaining Bytes
				fmt.Println(string(p[:n]))
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println((string(p[:n])))
	}
}

func serverExample() {
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
