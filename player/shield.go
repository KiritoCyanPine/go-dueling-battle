package player

// Shield is a struct that creates the shield of the player
type Shield struct {
	defpoints int
	attribute Attributes
}

// NewShield will create a new shield
func NewShield(attribute Attributes, player *Player) Shield {
	defpoints := 0

	switch attribute {
	case player.Attribute:
		defpoints = 3
	case ElementalWeakmess[player.Attribute]:
		defpoints = 1
	default:
		defpoints = 2
	}

	return Shield{
		defpoints: defpoints,
		attribute: attribute,
	}
}

// BlockDamage is a function that determines the damage that the shield can block
func (s *Shield) BlockDamage(d *Damage) {
	if s.defpoints == 0 {
		return
	}

	s.defpoints--

	if s.attribute == d.Attributed.Attributes || s.attribute == VOID {
		d.Attributed.Damage -= int(float32(d.Attributed.Damage) / 2)
	}

	d.Base -= int(float32(d.Base) / 4)
	d.Critical -= int(float32(d.Critical) / 3)
}
