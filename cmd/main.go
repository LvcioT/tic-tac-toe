package main

import (
	"fmt"
	"tictactoe/internal/board"
	"tictactoe/internal/move"
)

const (
	BoardDimension = 3
)

func main() {
	b := board.NewRandomBoard(BoardDimension)

	fmt.Println("Board:")
	fmt.Println(b.String())

	moves := move.Available(b)
	fmt.Println("Moves:")
	for _, m := range moves {
		fmt.Println(m.String())
	}
}
