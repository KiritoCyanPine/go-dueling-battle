package duel

import (
	"github.com/kiritocyanpine/goduel/player"
)

// Duel is a fighting ritual where only two players engage in battle
type Duel struct {
	p1 *player.Player
	p2 *player.Player
}

// CreateDuel will initialize two players engaging in a duel
func CreateDuel(player1, player2 *player.Player) *Duel {
	return &Duel{
		p1: player1,
		p2: player2,
	}
}

// Player1TakesAction is the first player to take action attack or Shielding
func (d *Duel) Player1TakesAction(attribute player.Attributes, isSielding bool, statuseffect player.StatusEffect) Status {
	if isSielding {
		d.player1Shield(attribute)
		return d.matchStatus()
	}
	d.player1AttacksPlayer2(attribute, statuseffect)
	return d.matchStatus()
}

// Player2TakesAction is the second player to take action attack or Shielding
func (d *Duel) Player2TakesAction() Status {
	d.player2AttacksPlayer1()
	return d.matchStatus()
}

func (d *Duel) player1AttacksPlayer2(attribute player.Attributes, statuseffect player.StatusEffect) {
	damage := d.p1.InflictDamage(attribute, statuseffect)
	d.p2.TakeDamage(damage)
}

func (d *Duel) player2AttacksPlayer1() {
	damage := d.p2.InflictDamage(d.p2.Attribute, player.NoEffect)
	d.p1.TakeDamage(damage)
}

func (d *Duel) player1Shield(attribute player.Attributes) {
	d.p1.Shielding = player.NewShield(attribute, d.p1)
}

func (d *Duel) player2Shield(attribute player.Attributes) {
	d.p2.Shielding = player.NewShield(attribute, d.p2)
}

// matchStatus is a function that determines the status of the match
func (d *Duel) matchStatus() Status {
	if d.p1.IsDefeated() {
		return Player2Wins
	}

	if d.p2.IsDefeated() {
		return Player1Wins
	}

	return InProgress
}
