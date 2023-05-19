package main

import (
	"log"
	"net/http"

	"github.com/austinmorales/chargen/engine"
)

func main() {
	log.Println("Service starting...")
	http.HandleFunc("/random", engine.RandomCharHandler)
	http.ListenAndServe(":8080", nil)
}
