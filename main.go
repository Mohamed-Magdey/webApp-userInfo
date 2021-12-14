package main

import (
	"net/http"
	"webApp/userInfo/handlers"
)

func main() {

	http.HandleFunc("/", handlers.IndexFunc)

	http.ListenAndServe(":8080", nil)
}
