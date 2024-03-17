package board

import "math/rand"

type Stone int

const (
	StoneO Stone = -1
	StoneX Stone = 1
	Empty  Stone = 0
)

func NewStoneFromRune(r rune) Stone {
	switch r {
	case -1:
		return StoneO
	case 1:
		return StoneX
	default:
		return Empty
	}
}

func NewStoneFromInt(n int) Stone {
	switch n {
	case -1:
		return StoneO
	case 1:
		return StoneX
	default:
		return Empty
	}
}

func NewStoneRandom() Stone {
	n := rand.Intn(3) - 1

	return NewStoneFromInt(n)
}

func (s Stone) Rune() rune {
	switch s {
	case StoneO:
		return 'O'
	case StoneX:
		return 'X'
	default:
		return ' '
	}
}
