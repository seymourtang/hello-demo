package main

import (
	"log"
	"net/http"

	"hello-demo/internal"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("world"))
	})
	http.HandleFunc("/uuid", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(internal.GetUuid()))
	})
	log.Printf("http server is starting,port:%d", 8081)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
