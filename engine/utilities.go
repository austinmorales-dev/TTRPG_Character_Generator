package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
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

func ImportNames() *datastructs.ImportedName {
	jsonFile, err := os.Open("./etc/names.json")
	if err != nil {
		log.Fatal("Couldn't open JSON file...")
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var importedName datastructs.ImportedName
	json.Unmarshal(byteValue, &importedName)

	return &importedName
}

func GenChar() string {
	rand.Seed(time.Now().UnixNano()) //seed the gen
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
	return GenJson(char)
}

func GenDamage() string {
	sides := []int{6, 8, 10, 12}
	rand.Seed(time.Now().UnixNano())
	amount := rand.Intn(2) + 1
	dRoll := fmt.Sprintf("%vd%v", amount, sides[rand.Intn(len(sides))])
	return dRoll
}

type ProtoWeapon struct {
	Name       string
	DamageType string
}

func GWeapon() string {
	var genweap ProtoWeapon
	var candidates []ProtoWeapon
	rand.Seed(time.Now().UnixNano())
	weaponNouns := []string{"Destruction",
		"Annihilation",
		"Doom",
		"Justice",
		"Shadows",
		"Wrath",
		"Chaos",
		"Light",
		"Darkness",
		"Vengeance",
		"Heroes",
		"Souls",
		"Power",
		"Thunder",
		"Frost",
		"Flames",
		"Life",
		"Death",
		"Dragons",
		"Legends",
	}
	meleeDT := []string{"slashing", "bludgeoning"}
	rangedDT := []string{"piercing"}
	// magicDT := []string{"cold", "poison", "fire", "acid", "psychic", "necrotic", "radiant", "force", "thunder", "lightning"}

	weapons := []ProtoWeapon{
		{Name: "axe", DamageType: meleeDT[0]},
		{Name: "sword", DamageType: meleeDT[0]},
		{Name: "warhammer", DamageType: meleeDT[1]},
		{Name: "club", DamageType: meleeDT[1]},
		{Name: "quarterstaff", DamageType: meleeDT[1]},
		{Name: "shield", DamageType: meleeDT[1]},
		{Name: "bow", DamageType: rangedDT[0]},
		{Name: "short bow", DamageType: rangedDT[0]},
		{Name: "dart", DamageType: rangedDT[0]},
		{Name: "crossbow", DamageType: rangedDT[0]},
	}

	d6roll := rand.Intn(6) + 1
	if d6roll > 3 {
		for _, weapon := range weapons {
			if weapon.DamageType == meleeDT[0] || weapon.DamageType == meleeDT[1] {
				candidates = append(candidates, weapon)
			}
		}
	} else {
		for _, weapon := range weapons {
			if weapon.DamageType == rangedDT[0] {
				candidates = append(candidates, weapon)
			}
		}

	}
	if len(candidates) > 0 {
		genweap = candidates[rand.Intn(len(candidates))]
	}
	randNoun := weaponNouns[rand.Intn(len(weaponNouns))]
	fullTitle := strings.ToTitle(fmt.Sprintf("%v of %v", genweap.Name, randNoun))

	weapon := datastructs.Weapon{
		Name:       fullTitle,
		Damage:     GenDamage(),
		Properties: genweap.DamageType,
	}
	return GenJson(weapon)
}
