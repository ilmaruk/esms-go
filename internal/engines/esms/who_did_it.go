package esms

import (
	"slices"

	"github.com/ilmaruk/esms-go/internal/models"
	"github.com/ilmaruk/esms-go/internal/random"
)

func whoDidIt(r random.Random, ts *models.Teamsheet, a Action) *models.Player {
	var total, weight float64
	var ar = make([]float64, 11)

	for i, p := range ts.Players {
		switch a {
		case ACTION_SHOT:
			weight += p.ShContrib * 100.
			total += ts.TeamShooting * 100.
		case ACTION_FOUL:
			weight += p.Ag
			total += ts.Aggression
		case ACTION_TACKLE:
			weight += p.TkContrib * 100.
			total += ts.TeamTackling * 100.
		case ACTION_ASSIST:
			weight += p.PsContrib * 100.
			total += ts.TeamPassing * 100.
		}

		ar[i] = weight
	}

	rv := r.FloatToMax(total)
	k, _ := slices.BinarySearch(ar, rv)
	return ts.Players[k]
}
