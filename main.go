package main

import (
	"fmt"
	"log"

	"github.com/kiritocyanpine/goduel/duel"
	"github.com/kiritocyanpine/goduel/player"
)

type attackInput struct {
	attribute    player.Attributes
	isSielding   bool
	statuseffect player.StatusEffect
}

func main() {
	p1 := player.CreateCustomPlayer(player.Info{Name: "Hisoka", Title: "The Transputor"}, player.Mage, 1, player.NoEffect, player.Fire, player.CreateCustomGenesis(110, 2), player.NewCustomCritical(100, 100))
	p2 := player.NewRandomPlayer(player.Info{Name: "JonSnow", Title: "The Bastard"}, player.Knight, 1, player.Fire)

	fmt.Printf("your opponent is %v", p2)
	status := duel.InProgress
	for {
		dueling := duel.CreateDuel(p1, p2)

		var w1, w2, w3 string
		_, err := fmt.Scanln(&w1, &w2, &w3)
		if err != nil {
			log.Fatal(err)
		}

		attacks := readAttacks(w1, w2, w3)

		status = dueling.Player1TakesAction(attacks.attribute, attacks.isSielding, attacks.statuseffect)
		if status != duel.InProgress {
			break
		}

		status = dueling.Player2TakesAction()
		if status != duel.InProgress {
			break
		}

	}

	fmt.Println(status)
}

func readAttacks(w1, w2, w3 string) attackInput {
	a := attackInput{}
	a.attribute = player.Attributes(w1)
	a.isSielding = w2 == "1"
	a.statuseffect = player.StatusEffect(w3)
	return a
}
