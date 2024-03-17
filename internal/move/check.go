package move

import "tictactoe/internal/board"

func Available(b board.Board) (m []Move) {
	d := len(b)

	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			if b[i][j] == board.Empty {
				m = append(m, Move{i, j})
			}
		}
	}

	return
}
