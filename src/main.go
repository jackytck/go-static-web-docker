package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
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
	mux := http.NewServeMux()
	mux.Handle("/", fs)

	handler := cors.Default().Handler(mux)
	log.Printf("Listening at %v...\n", port)
	http.ListenAndServe(port, handler)
}
