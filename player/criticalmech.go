package player

import (
	"math/rand"
	"time"
)

// Critical is a struct that determines the critical hit of the player
type Critical struct {
	// CriticalFactor is a factor that determines the critical hit of the player
	CriticalFactor int
	// CritRate is the rate of the player to hit a critical hit
	CritRate int
}

// NewCritical is a function that creates a new Critical struct
func NewCritical() Critical {
	return Critical{
		CriticalFactor: 10,
		CritRate:       1,
	}
}

// NewCustomCritical is a function that creates a new Critical struct that is customized
func NewCustomCritical(factor, rate int) Critical {
	return Critical{
		CriticalFactor: factor,
		CritRate:       rate,
	}
}

// CriticalHit is a function that determines the critical hit
// of the player based on the current damage
func (c *Critical) CriticalHit(damage *Damage) {
	rand.Seed(time.Now().UnixNano())
	probability := rand.Intn(100)

	if probability > c.CritRate {
		return
	}

	damage.Critical = damage.TotalDamage() * (c.CriticalFactor / 100)
}
