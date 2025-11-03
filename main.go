package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	port := ":4038"
	log.Printf("🚀 Server running at http://127.0.0.1%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
