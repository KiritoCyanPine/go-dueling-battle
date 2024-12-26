package player

// Attributes declares affinity to an element
type Attributes string

const (
	// Fire is an element that is associated with heat and destruction
	Fire Attributes = "FIRE"
	// Water is an element that is associated with fluidity and life
	Water Attributes = "WATER"
	// Air is an element that is associated with freedom and movement
	Air Attributes = "AIR"
	// Earth is an element that is associated with stability and growth
	Earth Attributes = "EARTH"
	// HOLY is an element that is associated with purity and divinity
	HOLY Attributes = "HOLY"
	// VOID is an element that is associated with emptiness and nothingness
	VOID Attributes = "VOID"
)

// AttributedAttack is a function that determines the damage that the player can deal
func (a Attributes) AttributedAttack(damage *Damage, p *Player) {
	switch a {
	case p.Attribute:
		damage.Attributed.Damage = int(float64(damage.Base) * 1.5)
	case ElementalWeakmess[p.Attribute]:
		damage.Attributed.Damage = damage.Base
		damage.Base -= int(float64(damage.Base) * 0.5)
	default:
		damage.Attributed.Damage = damage.Base
		damage.Base -= int(float64(damage.Base) * 0.25)
	}

}

// AttributedDefence is a function that determines the damage that the player can block
func (a Attributes) AttributedDefence(damage *Damage, p *Player) {
	p.Shielding.BlockDamage(damage)

	for _, effect := range p.statuseffect {
		effect.ShieldEffect(damage)
	}
}

// ElementalWeakmess is a map that determines the weakness of an element to another element
var ElementalWeakmess = map[Attributes]Attributes{
	Fire:  Water,
	Water: Fire,
	Air:   Earth,
	Earth: Air,
	HOLY:  VOID,
	VOID:  VOID,
}
