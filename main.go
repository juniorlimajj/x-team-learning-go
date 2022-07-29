package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application-json")
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(`{"message":"Hello World!"}`))
	})

	fmt.Println("http server listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
