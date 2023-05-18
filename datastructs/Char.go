package datastructs

type Character struct {
	Name  FullName `json:"Name"`
	Stats Stats    `json:"Stats"`
}

type FullName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Stats struct {
	HP  int `json:"HP"`
	STR int `json:"STR"`
	DEX int `json:"DEX"`
	CON int `json:"CON"`
	WIS int `json:"WIS"`
	CHA int `json:"CHA"`
	INT int `json:"INT"`
}

type ImportedName struct {
	Human struct {
		FirstNames []string `json:"firstNames"`
		LastNames  []string `json:"lastNames"`
	} `json:"human"`
	Elf struct {
		FirstNames []string `json:"firstNames"`
		LastNames  []string `json:"lastNames"`
	} `json:"elf"`
}
