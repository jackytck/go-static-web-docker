package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	static, exist := os.LookupEnv("STATIC_DIR")
	if !exist {
		static = "static"
	}
	port, exist := os.LookupEnv("PORT")
	if !exist {
		port = ":3000"
	}

	fs := http.FileServer(http.Dir(static))
	http.Handle("/", fs)

	log.Printf("Listening at %v...\n", port)
	http.ListenAndServe(port, nil)
}
