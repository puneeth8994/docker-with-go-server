package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveHome)
	http.ListenAndServe(":8080", nil)
}

func serveHome(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

	responseData := map[string]string{
		"message": "Hello World",
	}

	res, _ := json.Marshal(responseData)
	writer.Write(res)
}
