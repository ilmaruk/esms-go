package esms

import (
	"github.com/ilmaruk/esms-go/internal/models"
	"github.com/ilmaruk/esms-go/internal/random"
)

// When a chance was generated for the team and assisted by the
// assister, who got the assist ?
//
// This is almost like who_did_it, but it also takes
// into account the side of the assister - a player on his side
// has a higher chance to get the assist.
//
// How it's done: if the side of the shooter (picked by who_did_it)
// is different from the side of the asssiter, who_did_it is run
// once again - but this happens only once. This increases the
// chance of the player on the same side to be picked, but leaves
// a possibility for other sides as well.
func WhoGotAssisted(rnd random.Random, ts *models.Teamsheet, a *models.Player) *models.Player {
	var shooter = a

	// Shooter and assister must be different, so re-run each time the same
	// one is generated
	//
	for shooter == a {
		shooter = WhoDidIt(rnd, ts, ACTION_SHOT)

		// if the side is different, re-run once
		//
		if shooter.Side != a.Side {
			shooter = WhoDidIt(rnd, ts, ACTION_SHOT)
		}
	}

	return shooter
}
