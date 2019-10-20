package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "hello")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))
	_ = http.ListenAndServe(":8100", nil)
}
