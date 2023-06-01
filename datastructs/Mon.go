package datastructs

type Monster struct {
	ID         IDProps     `json:"ID"`
	CProps     CombatProps `json:"CProps"`
	Attr       Attributes  `json:"Attributes"`
	SpecTraits []DString   `json:"SpecialTraits"`
	Actions    []DString   `json:"Actions"`
	LActions   []DString   `json:"LActions"`
}

type IDProps struct {
	Name         string `json:"Name"`
	CreatureType string `json:"CreatureType"`
	Size         string `json:"Size"`
	Alignment    string `json:"Alignment"`
}

type CombatProps struct {
	AC        StrInt   `json:"AC"`
	Statblock Stats    `json:"Stats"`
	Movement  []StrInt `json:"Movement"`
}

type Attributes struct {
	SavingThrows        []StrInt  `json:"SavingThrows"`
	Skills              []StrInt  `json:"Skills"`
	DamageImmunities    []string  `json:"DamageImmunities"`
	ConditionImmunities []string  `json:"ConditionImmunities"`
	DamageResistances   []string  `json:"DamageResistances"`
	Senses              []StrInt  `json:"Senses"`
	Languages           []string  `json:"Languages"`
	Challenge           Challenge `json:"Challenge"`
}

type Challenge struct {
	Value int `json:"Value"`
	XP    int `json:"XP"`
}

type StrInt struct {
	Name  string `json:"Name"`
	Value int    `json:"Value"`
}

type DString struct { //double string
	Name string `json:"Name"`
	Desc string `json:"Desc"`
}
