package engine

import (
	"io"
	"math/rand"
	"net/http"

	"github.com/austinmorales/chargen/datastructs"
)

func RandomCharHandler(w http.ResponseWriter, r *http.Request) {
	var first_name string
	var last_name string
	name := ImportNames()
	roll_d6 := rand.Intn(6) + 1 //just to stay in theme haha
	if roll_d6 < 3 {
		first_name = name.Human.FirstNames[rand.Intn(25)]
		last_name = name.Human.LastNames[rand.Intn(25)]
	} else {
		first_name = name.Elf.FirstNames[rand.Intn(25)]
		last_name = name.Elf.LastNames[rand.Intn(25)]
	}
	char := datastructs.Character{
		Name: datastructs.FullName{
			FirstName: first_name,
			LastName:  last_name,
		},
		Stats: datastructs.Stats{
			HP:  rand.Intn(12) + 6,
			STR: rand.Intn(12) + 1,
			CHA: rand.Intn(12) + 1,
			INT: rand.Intn(12) + 1,
			DEX: rand.Intn(12) + 1,
			WIS: rand.Intn(12) + 1,
			CON: rand.Intn(12) + 1,
		},
	}
	jsonResp := GenJson(char)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	io.WriteString(w, jsonResp)
}

func SampleCharHandler(w http.ResponseWriter, r *http.Request) { //POC of backend
	char := GenDefaultChar()
	jsonResp := GenJson(char)
	io.WriteString(w, jsonResp)
}
