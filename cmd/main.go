package main

import (
	"log"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("world"))
	})
	slog.Info("http server is starting", "port", 8081)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
