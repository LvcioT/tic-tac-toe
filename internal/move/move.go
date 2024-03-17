package move

import "fmt"

type Move struct {
	X int
	Y int
}

func (m Move) String() string {
	return fmt.Sprintf("(%d,%d)", m.X, m.Y)
}
