package main

import (
	"net/http"
	"webApp/userInfo/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexFunc)

	http.HandleFunc("/showuser/show", handlers.ShowUserFunc)
	http.HandleFunc("/showuser/", handlers.ShowUser)
	http.HandleFunc("/showuser/notsuccededshow/", handlers.NotSuccededShow)

	http.ListenAndServe(":8080", nil)
}
