package esms

type Action uint

const (
	ACTION_SHOT Action = iota
	ACTION_FOUL
	ACTION_TACKLE
	ACTION_ASSIST
)
