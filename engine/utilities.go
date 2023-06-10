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
	class, baseHP := GenClass()
	race, subrace, mods := GenRace()
	stats := GenerateStats(baseHP)
	// fmt.Println(stats)
	rm_stats := ApplyRaceMods(mods, stats)
	// fmt.Println(rm_stats)
	// fmt.Println(mods)
	first_name, last_name := GenName(race)
	if subrace != "" {
		race = fmt.Sprintf("%v %v", subrace, race)
	}
	char := datastructs.Character{
		Name: datastructs.FullName{
			FirstName: first_name,
			LastName:  last_name,
		},
		Race:       race,
		Stats:      *rm_stats,
		Alignment:  GenFeature("alignment"),
		Class:      class,
		Background: GenFeature("background"),
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

func GenClass() (string, int) {
	db := &Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	defer db.CloseDB()
	class, baseHP := db.GenerateClass()
	return class, baseHP
}

func GenRace() (string, string, []int) {
	db := &Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("failed to connected to db:", err)
	}
	defer db.CloseDB()
	name, subrace, mods := db.GenerateRace()
	return name, subrace, mods
}

func GenName(r string) (string, string) {
	db := &Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	defer db.CloseDB()
	fName, lName, err := db.GenerateNames(r)
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	return fName, lName
}

func GenFeature(query string) string {
	db := &Database{}
	err := db.ConnectToDB()
	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	defer db.CloseDB()
	switch query {
	case "alignment":
		return db.GenerateAlignment()
	case "background":
		return db.GenerateBackground()
	}

	if err != nil {
		log.Println("Failed to connect to DB: ", err)
	}
	return ""
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

func GenTestMon() *datastructs.Statblock {
	mon := &datastructs.Statblock{
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
			Stats: datastructs.Stats{
				HP:  RollAbility(),
				STR: RollAbility(),
				CHA: RollAbility(),
				INT: RollAbility(),
				DEX: RollAbility(),
				WIS: RollAbility(),
				CON: RollAbility(),
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
