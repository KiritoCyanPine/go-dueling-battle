package player

// AttributedDamage is a struct that determines the damage of the player
// in that specific attribute
type AttributedDamage struct {
	Attributes Attributes
	Damage     int
}

// Damage is a struct that determines the damage of the player
type Damage struct {
	Base             int
	Attributed       AttributedDamage
	Critical         int
	inflictingEffect StatusEffect
}

// TotalDamage is a function that calculates the total damage of the player
func (d Damage) TotalDamage() int {
	return d.Base + d.Attributed.Damage + d.Critical
}
