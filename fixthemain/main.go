package main

import "github.com/01-edu/z01"

type Door struct {
	state int
}

const (
	CLOSE = 0
	OPEN  = 1
)

func printLine(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) bool {
	printLine("Door Opening...")
	ptrDoor.state = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	printLine("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	printLine("is the Door opened ?")
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	printLine("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
