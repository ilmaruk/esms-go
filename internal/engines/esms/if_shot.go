package esms

import (
	"github.com/ilmaruk/esms-go/internal/models"
	"github.com/ilmaruk/esms-go/internal/random"
)

func IfShot(rnd random.Random, attacking *models.Teamsheet) {
	// Did a scoring chance occur ?
	if !rnd.RandomP(attacking.ShotProb) {
		return
	}

	var assister *models.Player
	var chanceAssisted = false

	// There's a 0.75 probability that a chance was assisted, and
	// 0.25 that it's a solo
	if !rnd.RandomP(7500) {
		assister = WhoDidIt(rnd, attacking, ACTION_ASSIST)
		chanceAssisted = true
		shooter = WhoGotAssist(rnd, attacking, assister)
	}
}
