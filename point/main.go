package main

import "github.com/01-edu/z01"

type point struct{ x, y int }

func setPoint(ptr *point) { ptr.x, ptr.y = 42, 21 }

func main() {
	var p point
	setPoint(&p)

	s := []rune{120, 32, 61, 32, 52, 50, 44, 32, 121, 32, 61, 32, 50, 49, 10}
	for i := 0; i < 15; i++ {
		z01.PrintRune(s[i])
	}
}
