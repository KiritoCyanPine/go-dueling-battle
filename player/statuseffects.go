package player

import "math/rand"

// StatusEffect is a cause that has an emotional or physical connection to the player
// that influences the properties of the player
type StatusEffect string

const (
	// NoEffect is a status effect that does not have any effect on the player
	NoEffect StatusEffect = "NO"

	// PureEvilDamage is a status effect that increases the damage of the player
	PureEvilDamage StatusEffect = "PURE_EVIL" // totaldamage + 1/3
	// TheWettingCurseDamage is a status effect that decreases the damage of the player
	TheWettingCurseDamage StatusEffect = "WETTING_CURSE" // totaldamage - 1/4 critical - 1/4
	// SapOfEvilDamage is a status effect that increases the critical of the player
	SapOfEvilDamage StatusEffect = "SAP_OF_EVIL" // critical + 1/4
	// TrueKnightsSwordDamage is a status effect that increases the base damage of the player
	TrueKnightsSwordDamage StatusEffect = "TRUE_KNIGHTS_SWORD" // basedamage + 1/10
	// ShadowOfAssassinDamage is a status effect that increases the damage of the player
	ShadowOfAssassinDamage StatusEffect = "SHADOW_OF_ASSASSIN" // attribute.damage + 1/10 critical + 1/10 basedamage + 1/10
	// StudentOfMagicalTowerDamage is a status effect that increases the damage of the player
	StudentOfMagicalTowerDamage StatusEffect = "STUDENT_OF_MAGICAL_TOWER" // attribute.damage + 1/8 critical + 1/10
	// PriestsBlessingDamage is a status effect that increases the damage of the player
	PriestsBlessingDamage StatusEffect = "PRIESTS_BLESSING" // attribute.damage + 1/5
	// SnipersPrecisionDamage is a status effect that increases the damage of the player
	SnipersPrecisionDamage StatusEffect = "SNIPERS_PRECISION" // basedamage + 1/5

	// ReflectAttackShield is a status effect that decreases the damage to the player
	ReflectAttackShield StatusEffect = "REFLECT_ATTACK"
	// HolyNullificationShield is a status effect that decreases the damage to the player
	HolyNullificationShield StatusEffect = "HOLY_NULLIFICATION" // attribute.damage = 0
	// TanksShield is a status effect that decreases the damage to the player
	TanksShield StatusEffect = "TANKS_SHIELD" // totaldamage - 1/5
)

// ApplyEffect is a function that determines the effect of the status effect on the Damage given
func (s StatusEffect) ApplyEffect(d *Damage) {
	switch s {
	case PureEvilDamage:
		d.Base += (d.Base / 3)
		d.Attributed.Damage += (d.Attributed.Damage / 3)
	case TheWettingCurseDamage:
		d.Base -= (d.Base / 4)
		d.Attributed.Damage -= (d.Attributed.Damage / 4)
		d.Critical -= (d.Critical / 4)
	case SapOfEvilDamage:
		d.Critical += (d.TotalDamage() / 4)
	case TrueKnightsSwordDamage:
		d.Base += (d.Base / 10)
	case ShadowOfAssassinDamage:
		d.Base += (d.Base / 10)
		d.Attributed.Damage += (d.Attributed.Damage / 10)
		d.Critical += (d.Critical / 10)
	case StudentOfMagicalTowerDamage:
		d.Base += (d.Base / 8)
		d.Attributed.Damage += (d.Attributed.Damage / 10)
		d.Critical += (d.Critical / 10)
	case PriestsBlessingDamage:
		d.Attributed.Damage += (d.Attributed.Damage / 5)
	case SnipersPrecisionDamage:
		d.Base += (d.Base / 5)
	}
}

// ShieldEffect is a function that determines the effect of the status effect on the Damage taken
func (s StatusEffect) ShieldEffect(d *Damage) {
	switch s {
	case ReflectAttackShield:
		d.Base += (d.Base / 3)
		d.Attributed.Damage += (d.Attributed.Damage / 3)
	case HolyNullificationShield:
		d.Attributed.Damage = 0
	case TanksShield:
		d.Base -= (d.Base / 5)
	}
}

var commonStatusEffect = []StatusEffect{
	NoEffect,
	TanksShield,
	ShadowOfAssassinDamage,
	StudentOfMagicalTowerDamage,
	PriestsBlessingDamage,
	SnipersPrecisionDamage,
	TrueKnightsSwordDamage,
}

func getRandomStatusEffect() StatusEffect {
	return commonStatusEffect[rand.Intn(len(commonStatusEffect))]
}
