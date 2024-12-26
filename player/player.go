package player

import "fmt"

// BaseGenesis defines the qulities that the player was born with.
// It is the basic determinant for potential
type BaseGenesis struct {
	healthpoints int
	attackpoint  int
}

func baseGenesisStats() BaseGenesis {
	return BaseGenesis{
		healthpoints: 100,
		attackpoint:  1,
	}
}

// CreateCustomGenesis will create a custom genesis for the player
func CreateCustomGenesis(healthpoints, attackpoint int) BaseGenesis {
	return BaseGenesis{
		healthpoints: healthpoints,
		attackpoint:  attackpoint,
	}
}

// Info is the basic information of the player
type Info struct {
	Name  string
	Title string
}

// Player is a character that'll engage in duel
type Player struct {
	info         Info
	baseStat     BaseGenesis
	statuseffect []StatusEffect
	class        Class
	level        int
	health       int
	Attribute    Attributes
	Shielding    Shield
	critical     Critical
}

// NewRandomPlayer will create a new player
func NewRandomPlayer(info Info, class Class, level int, attribute Attributes) *Player {
	basestat := baseGenesisStats()
	return &Player{
		info:         info,
		baseStat:     basestat,
		statuseffect: []StatusEffect{getRandomStatusEffect()},
		class:        class,
		level:        level,
		health:       basestat.healthpoints * level,
		Attribute:    attribute,
		critical:     NewCritical(),
	}
}

// CreateCustomPlayer will create a new player
func CreateCustomPlayer(info Info, class Class, level int, effect StatusEffect, attribute Attributes, basestat BaseGenesis, critical Critical) *Player {
	return &Player{
		info:         info,
		baseStat:     basestat,
		statuseffect: []StatusEffect{effect},
		class:        class,
		level:        level,
		health:       basestat.healthpoints * level,
		Attribute:    attribute,
		critical:     critical,
	}
}

// InflictDamage will calculate the damage that the player will inflict
func (p *Player) InflictDamage(atribute Attributes, inflictingEffect StatusEffect) *Damage {
	damage := Damage{
		Base: p.baseStat.attackpoint * p.level,
		Attributed: AttributedDamage{
			Attributes: p.Attribute,
			Damage:     0,
		},
		inflictingEffect: inflictingEffect,
	}
	atribute.AttributedAttack(&damage, p)
	p.critical.CriticalHit(&damage)

	for _, effect := range p.statuseffect {
		effect.ApplyEffect(&damage)
	}

	return &damage
}

// TakeDamage will calculate the damage that the player will take
func (p *Player) TakeDamage(damage *Damage) {
	p.Attribute.AttributedDefence(damage, p)
	p.health -= damage.TotalDamage()
	if damage.inflictingEffect != NoEffect {
		p.statuseffect = append(p.statuseffect, damage.inflictingEffect)
	}

	fmt.Println(p.info.Name, "has taken", damage.TotalDamage(), "damage")
	fmt.Printf("%v \n", p)
}

// MakeShield will create a shield for the player
func (p *Player) MakeShield(attribute Attributes) {
	p.Shielding = NewShield(p.Attribute, p)
}

// IsDefeated will check if the player is defeated
func (p *Player) IsDefeated() bool {
	return p.health <= 0
}
