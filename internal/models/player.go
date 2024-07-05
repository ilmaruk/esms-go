package models

type Player struct {
	ShContrib float64
	TkContrib float64
	PsContrib float64

	Ag int

	Side string

	KeyPasses int
	Tackles   int
	ShotsOff  int

	Fatigue float64
}
