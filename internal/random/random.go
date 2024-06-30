package random

import (
	"math/rand"
	"time"
)

type Random interface {
	FloatToMax(max float64) float64
}

type DefaultRandomiser struct {
	rnd *rand.Rand
}

func NewDefaultRandomiser() *DefaultRandomiser {
	return &DefaultRandomiser{
		rnd: rand.New(rand.NewSource(time.Now().UnixMicro())),
	}
}

var _ Random = &DefaultRandomiser{}

func (r DefaultRandomiser) FloatToMax(max float64) float64 {
	return r.rnd.Float64() * max
}
