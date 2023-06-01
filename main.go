package main

import (
	"log"
	"net/http"
	"os"

	"github.com/austinmorales/chargen/engine"
)

func main() {
	log.Println("Service starting...")
	if os.Getenv("DB_URL") == "" {
		log.Println("DB_URL is empty...")
		os.Exit(88)
	}
	http.HandleFunc("/character", engine.RandomCharHandler)
	http.HandleFunc("/weapon", engine.GenWeaponHandler)
	http.HandleFunc("/npc", engine.GenNPCHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORTNO"), nil)
	if err != nil {
		log.Printf("Error starting server: %v", err)
	}
}
