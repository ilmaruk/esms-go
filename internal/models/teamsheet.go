package models

type Teamsheet struct {
	Players []*Player

	TeamShooting int
	TeamTackling int
	TeamPassing  int

	Aggression int

	ShotProb int
}
