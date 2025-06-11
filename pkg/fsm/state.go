package fsm

import "fmt"

type stateFn func(event Event) stateFn

func redState(event Event) stateFn {
	switch event.Type {
	case timerElapsed:
		fmt.Println("Red -> Green")
		return greenState
	case emergency:
		fmt.Println("Already in red; stay on red due to emergency")
		return redState
	case powerOutage:
		fmt.Println("Power outage")
		return nil
	default:
		fmt.Printf("redState: Unhandled event: %v\n", event.Type)
		return redState
	}
}

func greenState(event Event) stateFn {
	switch event.Type {
	case timerElapsed:
		fmt.Println("Green -> Yellow")
		return yellowState
	case emergency:
		fmt.Println("Emergency! Green -> Red")
		return redState
	case powerOutage:
		fmt.Println("Power outage")
		return nil
	default:
		fmt.Printf("greenState: Unhandled event: %v\n", event.Type)
		return greenState
	}
}

func yellowState(event Event) stateFn {
	switch event.Type {
	case timerElapsed:
		fmt.Println("Yellow -> Red")
		return greenState
	case emergency:
		fmt.Println("Emergency! Yellow -> Red")
		return redState
	case powerOutage:
		fmt.Println("Power outage")
		return nil
	default:
		fmt.Printf("yellowState: Unhandled event: %v\n", event.Type)
		return yellowState
	}
}
