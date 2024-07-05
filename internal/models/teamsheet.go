package models

type Teamsheet struct {
	Players   []*Player
	CurrentGK *Player

	TeamShooting float64
	TeamTackling float64
	TeamPassing  float64

	Aggression int

	ShotProb int

	Score         int
	FinalShotsOff int
}
