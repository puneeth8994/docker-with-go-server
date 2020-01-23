package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveHome)
	http.ListenAndServe(":8080", nil)
}

func serveHome(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		errorHandler(writer, request, http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")

	responseData := map[string]string{
		"message": "Hello World",
	}

	res, _ := json.Marshal(responseData)
	writer.Write(res)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "<h1>404 not found</h1>")
	}
}
