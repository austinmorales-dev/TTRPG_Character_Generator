package engine

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/austinmorales/chargen/datastructs"
)

func GenJson(obj any) string { //quick helper function to convert characters to JSON
	jsonResp, err := json.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonResp)
}

func GenChar() string {
	rand.Seed(time.Now().UnixNano()) //seed the gen
	var first_name, last_name, race string
	roll_d6 := rand.Intn(6) + 1 //just to stay in theme haha
	if roll_d6 < 3 {
		first_name, last_name, race = GenName("human")
	} else {
		first_name, last_name, race = GenName("elf")
	}
	char := datastructs.Character{
		Name: datastructs.FullName{
			FirstName: first_name,
			LastName:  last_name,
		},
		Race: race,
		Stats: datastructs.Stats{
			HP:  rand.Intn(12) + 6,
			STR: rand.Intn(12) + 1,
			CHA: rand.Intn(12) + 1,
			INT: rand.Intn(12) + 1,
			DEX: rand.Intn(12) + 1,
			WIS: rand.Intn(12) + 1,
			CON: rand.Intn(12) + 1,
		},
		Alignment: GenAlignment(),
	}
	return GenJson(char)
}

func GWeapon() string {
	genweap := GenWeapon()
	eName, eDesc, eWeapName := GenEnchantment()
	fullTitle := strings.ToTitle(fmt.Sprintf("%v of %v", genweap.Name, eWeapName))
	genweap.Name = fullTitle
	genweap.Enchantments = datastructs.Enchantment{
		Name:        eName,
		Description: eDesc,
	}
	return GenJson(genweap)
}

func GenEnchantment() (string, string, string) {
	db := &Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	defer db.CloseDB()
	e_name, e_desc, eWeapName, err := db.GenerateEnchantment()
	if err != nil {
		log.Println("Failed to generate from table: ", err)
	}
	return e_name, e_desc, eWeapName
}

func GenName(r string) (string, string, string) {
	db := &Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	defer db.CloseDB()
	fName, lName, race, err := db.GenerateNames(r)
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	return fName, lName, race
}

func GenAlignment() string {
	db := &Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	defer db.CloseDB()
	alignment := db.GenerateAlignment()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	return alignment
}

func GenWeapon() *datastructs.Weapon {
	weapon := &datastructs.Weapon{}
	var props []string
	db := &Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	defer db.CloseDB()
	name, dt, dr, props := db.GenerateWeapon()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	weapon.Name = name
	weapon.DamageRoll = dr
	weapon.DamageType = dt
	weapon.Properties = props
	return weapon
}
