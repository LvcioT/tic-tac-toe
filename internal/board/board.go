package board

type Board [][]Stone

func NewBoard(d int) Board {
	b := make(Board, d)

	for i := 0; i < d; i++ {
		b[i] = make([]Stone, d)
	}

	return b
}

func NewRandomBoard(d int) Board {
	b := NewBoard(d)

	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			b[i][j] = NewStoneRandom()
		}
	}

	return b
}
