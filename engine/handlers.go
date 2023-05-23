package engine

import (
	"io"
	"net/http"
)

func RandomCharHandler(w http.ResponseWriter, r *http.Request) {
	char := GenChar()
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, char)
}

func GenWeaponHandler(w http.ResponseWriter, r *http.Request) {
	weapon := GWeapon()
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, weapon)
}
