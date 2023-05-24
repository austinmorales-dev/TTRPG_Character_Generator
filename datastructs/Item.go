package datastructs

type Weapon struct {
	Name         string      `json:"name"`
	DamageRoll   string      `json:"damageRoll"`
	DamageType   string      `json:"damageType"`
	Enchantments Enchantment `json:"enchantments"`
	Properties   []string    `json:"properties"`
}

type Enchantment struct {
	Name        string
	Description string
}
