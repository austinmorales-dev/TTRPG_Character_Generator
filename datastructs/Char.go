package datastructs

type Character struct {
	Name       FullName `json:"Name"`
	Stats      Stats    `json:"Stats"`
	Race       string   `json:"Race"`
	Alignment  string   `json:"Alignment"`
	Class      string   `json:"Class"`
	Background string   `json:"Background"`
}

type FullName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Stats struct {
	HP     int `json:"HP"`
	STR    int `json:"STR"`
	STRmod int `json:"STRm"`
	DEX    int `json:"DEX"`
	DEXmod int `json:"DEXm"`
	CON    int `json:"CON"`
	CONmod int `json:"CONm"`
	INT    int `json:"INT"`
	INTmod int `json:"INTm"`
	WIS    int `json:"WIS"`
	WISmod int `json:"WISm"`
	CHA    int `json:"CHA"`
	CHAmod int `json:"CHAm"`
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
