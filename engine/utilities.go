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
		Class:     GenClass(),
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

func GenTestMon() *datastructs.Monster {
	mon := &datastructs.Monster{
		ID: datastructs.IDProps{
			Name:         "Ancient Black Dragon",
			CreatureType: "Dragon",
			Size:         "Gargantuan",
			Alignment:    "Chaotic evil",
		},
		CProps: datastructs.CombatProps{
			AC: datastructs.StrInt{
				Name:  "Natural Armor",
				Value: 22,
			},
			Statblock: datastructs.Stats{
				HP:  rand.Intn(20) + 6,
				STR: rand.Intn(20) + 6,
				CHA: rand.Intn(20) + 6,
				INT: rand.Intn(20) + 6,
				DEX: rand.Intn(20) + 6,
				WIS: rand.Intn(20) + 6,
				CON: rand.Intn(20) + 6,
			},
			Movement: []datastructs.StrInt{
				{Name: "Speed", Value: 40},
				{Name: "Fly", Value: 80},
				{Name: "Swim", Value: 40},
			},
		},
		Attr: datastructs.Attributes{
			SavingThrows: []datastructs.StrInt{
				{Name: "DEX", Value: 9},
				{Name: "CON", Value: 14},
				{Name: "WIS", Value: 9},
				{Name: "CHA", Value: 11},
			},
			Skills: []datastructs.StrInt{
				{Name: "Perception", Value: 16},
				{Name: "Stealth", Value: 9},
			},
			DamageImmunities: []string{
				"acid",
			},
			Senses: []datastructs.StrInt{
				{Name: "blindsight", Value: 60},
				{Name: "darkvision", Value: 120},
				{Name: "passive Perception", Value: 26},
			},
			Languages: []string{
				"common",
				"draconic",
			},
			Challenge: datastructs.Challenge{
				Value: 20,
				XP:    33000,
			},
		},
		SpecTraits: []datastructs.DString{
			{Name: "Amphibious", Desc: "Can breath water and air"},
			{Name: "Legendary Resistance", Desc: "If this fails a saving throw, it can choose to succeed instead."},
		},
		Actions: []datastructs.DString{
			{Name: "Multiattack", Desc: "Can attack once with it's bite, and twice with it's claws."},
			{Name: "Bite", Desc: "Melee Weapon Attack, +15 to hit, reach 15ft."},
			{Name: "Claw", Desc: "Melee Weapon Attack, +15 to hit, reach 10ft."},
		},
		LActions: []datastructs.DString{
			{Name: "Detect", Desc: "Make a Widsom (Perception) check."},
		},
	}
	return mon
}

func GenClass() string {
	db := &Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	defer db.CloseDB()
	class := db.GenerateClass()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	return class
}

func capFirst(word string) string {
	splitWord := strings.Split(word, "")
	splitWord[0] = strings.ToUpper(splitWord[0])
	// f := append(splitWord, convFirst)
	return strings.Join(splitWord, "")
}
