package random

import (
	"math/rand"
	"time"
)

type Random interface {
	// Generate a random number up to 10000. If the given p is
	// less than the generated number, return 1, otherwise return 0
	//
	// Used to "throw dice" and check if an event with some probability
	// happened. p is 0..10000 - for example 2000 means probability 0.2
	// So when 2000 is given, this function simulates an event with
	// probability 0.2 and tells if it happened (naturally it has
	// a prob. of 0.2 to happen)
	//
	RandomP(p int) bool

	// Returns a pseudo-random integer between 0 and N-1
	//
	MyRandom(n int) int
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

func (r DefaultRandomiser) RandomP(p int) bool {
	return r.MyRandom(10000) < p
}

func (r DefaultRandomiser) MyRandom(n int) int {
	return r.rnd.Intn(n)
}
