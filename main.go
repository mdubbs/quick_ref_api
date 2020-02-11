package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type headersResponse struct {
	Headers map[string][]string `json:"headers"`
	Message string `json:"message"`
}

type baseResponse struct {
	Message string `json:"message"`
}

func ping(w http.ResponseWriter, _ *http.Request) {
	blah, _ := json.Marshal(&baseResponse{Message:"pong"})
	_, _ = fmt.Fprintf(w, string(blah))
}

func headers(w http.ResponseWriter, req *http.Request) {
	headerMap := map[string][]string{}
	for name, headers := range req.Header {
		for _, h := range headers {
			headerMap[name] = append(headerMap[name], h)
		}
	}
	blah, _ := json.Marshal(&headersResponse{
		Headers: headerMap,
		Message: "hi",
	})
	fmt.Fprintf(w, string(blah))
}

func main() {
	http.HandleFunc("/", ping)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/health", ping)
	http.HandleFunc("/headers", headers)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
