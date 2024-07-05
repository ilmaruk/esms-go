package esms

import (
	"github.com/ilmaruk/esms-go/internal/models"
	"github.com/ilmaruk/esms-go/internal/random"
)

func IfShot(rnd random.Random, attacking, defending *models.Teamsheet) {
	// Did a scoring chance occur ?
	if !rnd.RandomP(attacking.ShotProb) {
		return
	}

	var assister, shooter *models.Player
	var chanceAssisted = false

	// There's a 0.75 probability that a chance was assisted, and
	// 0.25 that it's a solo
	if !rnd.RandomP(7500) {
		// Assister
		assister = WhoDidIt(rnd, attacking, ACTION_ASSIST)
		chanceAssisted = true
		assister.KeyPasses++

		// Shooter
		shooter = WhoGotAssist(rnd, attacking, assister)
	} else {
		shooter = WhoDidIt(rnd, attacking, ACTION_SHOT)
	}

	chanceTackled := (int)(4000.0 * ((defending.TeamTackling * 3.0) / (attacking.TeamPassing*2.0 + attacking.TeamShooting)))
	if rnd.RandomP(chanceTackled) {
		tackler := WhoDidIt(rnd, defending, ACTION_TACKLE)
		tackler.Tackles++
		return
	}

	if !OnTarget(rnd, shooter) {
		shooter.ShotsOff++
		attacking.FinalShotsOff++
		return
	}

	if !IsGoal(rnd, shooter) {
		defending.CurrentGK.Saves++
		return
	}

	if GoalIsCancelled() {
		return
	}

	attacking.Score++
	shooter.Goals++
	defending.CurrentGK.Conceded++

	// RF: assister != nil should be enough
	if chanceAssisted && assister != shooter {
		assister.Assists++
	}
}
