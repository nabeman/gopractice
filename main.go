package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", fooHandleFunc)

	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func fooHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/fooページです")
}

