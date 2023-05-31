package main

import (
	"log"
	"net/http"

	"github.com/austinmorales/chargen/engine"
)

func main() {
	log.Println("Service starting...")
	http.HandleFunc("/character", engine.RandomCharHandler)
	http.HandleFunc("/weapon", engine.GenWeaponHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error starting server: %v", err)
	}
}
