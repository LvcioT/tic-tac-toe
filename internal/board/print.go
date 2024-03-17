package board

import (
	"strings"
)

const (
	VL rune = '│'
	HL rune = '─'
	CL rune = '┼'
)

func (b Board) String() string {
	d := len(b)
	rows := make([]string, d)
	separator := "\n" + horizontalSeparator(d) + "\n"

	for i := 0; i < d; i++ {
		rows[i] = horizontalRow(b[i])
	}

	s := strings.Join(rows, separator)
	return s
}

func horizontalSeparator(n int) string {
	d := n*2 - 1
	runes := make([]rune, d)

	for i := 0; i < d; i++ {
		if i%2 == 0 {
			runes[i] = HL
		} else {
			runes[i] = CL
		}
	}

	s := string(runes)
	return s
}

func horizontalRow(row []Stone) string {
	d := len(row)*2 - 1
	runes := make([]rune, d)

	j := 0
	for i := 0; i < d; i++ {
		if i%2 == 0 {
			runes[i] = row[j].Rune()
			j++
		} else {
			runes[i] = VL
		}
	}

	s := string(runes)
	return s
}
