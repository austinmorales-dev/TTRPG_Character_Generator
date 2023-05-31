package datastructs

type Weapon struct {
	Name         string      `json:"Name"`
	DamageRoll   string      `json:"DamageRoll"`
	DamageType   string      `json:"DamageType"`
	Enchantments Enchantment `json:"Enchantments"`
	Properties   []string    `json:"Properties"`
}

type Enchantment struct {
	Name        string
	Description string
}
