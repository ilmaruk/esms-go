package models

type Player struct {
	St int
	Sh int

	ShContrib float64
	TkContrib float64
	PsContrib float64

	Ag int

	Side string

	KeyPasses int
	Tackles   int
	ShotsOff  int
	Saves     int
	Assists   int
	Goals     int
	Conceded  int

	Fatigue float64
}
