package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("it works"))
	})
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
