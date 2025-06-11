package fsm

import "fmt"

type state int

const (
	red state = iota
	yellow
	green
)

type action func(Event)

func redState(e Event) {
	fmt.Println("Red -> Green")
}

func greenState(e Event) {
	fmt.Println("Green -> Yellow")
}

func yellowState(e Event) {
	fmt.Println("Yellow -> Red")
}
