package main

import (
	"fmt"
	handlers "lww-elem-set-go/handlers"

	"net/http"
)

const (
	PORT = ":7777"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", handlers.Add)
	mux.HandleFunc("/lookup", handlers.Lookup)
	mux.HandleFunc("/list", handlers.List)
	mux.HandleFunc("/merge", handlers.Merge)
	mux.HandleFunc("/remove", handlers.Remove)

	fmt.Printf("Server is Running at %s\n", PORT)
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		fmt.Printf("Error starting server : %s\n", err)
	}

}
