package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/austinmorales/chargen/datastructs"
)

func GenDefaultChar() datastructs.Character { //generates a default character to test JSON response
	char := datastructs.Character{
		Name: datastructs.FullName{
			FirstName: "Link",
			LastName:  "of Hyrule",
		},
		Stats: datastructs.Stats{
			HP:  12,
			STR: 12,
			CHA: 12,
			INT: 12,
			DEX: 12,
			WIS: 12,
			CON: 12,
		},
	}
	return char
}

func GenJson(char datastructs.Character) string { //quick helper function to convert characters to JSON
	jsonResp, err := json.Marshal(char)
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

func RandomName() {
	name := ImportNames()
	fmt.Println(name.Human.FirstNames[1])
}
