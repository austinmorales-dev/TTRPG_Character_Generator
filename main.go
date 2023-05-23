package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/austinmorales/chargen/engine"
)

func main() {
	log.Println("Service starting...")
	http.HandleFunc("/random", engine.RandomCharHandler)
	http.HandleFunc("/weapon", engine.GenWeaponHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error starting server: %v", err)
	}
}

func dbCheck() {
	db := &engine.Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	defer db.CloseDB()
	gen, _ := db.GenerateFromDB("weapon_types")
	if err != nil {
		log.Println("Failed to generate from table: ", err)
	}
	fmt.Println(gen)
}
