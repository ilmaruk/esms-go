package esms

import (
	"slices"

	"github.com/ilmaruk/esms-go/internal/models"
	"github.com/ilmaruk/esms-go/internal/random"
)

// Given a team and an event (eg. SHOT)
// picks one player at (weighted) random
// that performed this event.
//
// For example, for SHOT, pick a player
// at weighted random according to his
// shooting skill
func WhoDidIt(r random.Random, ts *models.Teamsheet, a Action) *models.Player {
	var total, weight float64
	var ar = make([]float64, 11)

	for i, p := range ts.Players {
		switch a {
		case ACTION_SHOT:
			weight += p.ShContrib * 100
			total += ts.TeamShooting * 100.
		case ACTION_FOUL:
			weight += float64(p.Ag)
			total += float64(ts.Aggression)
		case ACTION_TACKLE:
			weight += p.TkContrib * 100
			total += ts.TeamTackling * 100.
		case ACTION_ASSIST:
			weight += p.PsContrib * 100
			total += ts.TeamPassing * 100.
		}

		ar[i] = weight
	}

	rv := r.MyRandom(int(total))
	k, _ := slices.BinarySearch(ar, float64(rv))
	return ts.Players[k]
}
