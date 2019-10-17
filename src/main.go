package main

import (
	"log"
	"net/http"
	"os"
	"strings"

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
	list, exist := os.LookupEnv("LIST_DIR")
	if !exist {
		list = "true"
	}

	fs := http.FileServer(blindFS{http.Dir(static)})
	if strings.ToLower(list) == "true" {
		fs = http.FileServer(http.Dir(static))
	}
	mux := http.NewServeMux()
	mux.Handle("/", fs)

	handler := cors.Default().Handler(mux)
	log.Printf("Listening at http://127.0.0.1%v...\n", port)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		panic(err)
	}
}

type blindFS struct {
	fs http.FileSystem
}

func (nfs blindFS) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := nfs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
