package esms

import (
	"github.com/ilmaruk/esms-go/internal/models"
	"github.com/ilmaruk/esms-go/internal/random"
)

// Given a shot on target (team a shot on team b's goal),
// was it a goal ?
func IsGoal(rnd random.Random, sh, gk *models.Player) bool {
	// Factors taken into account:
	// The shooter's Sh and fatigue against the GK's St
	//
	// The "median" is 0.35
	// Lower and upper bounds are 0.1 and 0.9 respectively
	//
	temp := sh.Sh*int(sh.Fatigue)*200 - gk.St*200 + 3500

	if temp > 9000 {
		temp = 9000
	} else if temp < 1000 {
		temp = 1000
	}

	return rnd.RandomP(temp)
}

func GoalIsCancelled(rnd random.Random) bool {
	return rnd.RandomP(500)
}
