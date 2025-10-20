package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]strings.Builder, numRows)
	row, step := 0, 1

	for _, ch := range s { // rune-safe
		rows[row].WriteRune(ch)
		// bounce at the top/bottom
		if row == 0 {
			step = 1
		} else if row == numRows-1 {
			step = -1
		}
		row += step
	}

	var out strings.Builder
	for i := 0; i < numRows; i++ {
		out.WriteString(rows[i].String())
	}
	return out.String()
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows)) // print "PAYPALISHIRING"
}
