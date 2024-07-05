package esms

import (
	"github.com/ilmaruk/esms-go/internal/models"
	"github.com/ilmaruk/esms-go/internal/random"
)

// OnTarget tells whether the shot is on target.
func OnTarget(rnd random.Random, s *models.Player) bool {
	return rnd.RandomP(int(5800 * s.Fatigue))
}
