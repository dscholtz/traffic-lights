package fsm

import (
	"context"
	"fmt"
)

type stateMachine struct {
	currentState state
	transitions  map[state]map[Event]state
	actions      map[state]map[Event]action
	eventChan    chan Event
	doneChan     chan struct{}
}

// Constructor
func New() *stateMachine {
	sm := &stateMachine{
		currentState: red,
		transitions:  make(map[state]map[Event]state),
		actions:      make(map[state]map[Event]action),
		eventChan:    make(chan Event, 10), // buffered for responsiveness
		doneChan:     make(chan struct{}),
	}

	// ---- Transitions and action in red ----
	sm.transitions[red] = map[Event]state{} // initialize once
	sm.transitions[red][Event{Type: timerElapsed}] = green
	sm.transitions[red][Event{Type: emergency}] = red

	sm.actions[red] = map[Event]action{}
	sm.actions[red][Event{Type: timerElapsed}] = redState
	sm.actions[red][Event{Type: emergency}] = func(e Event) {
		fmt.Printf("Emergency avoided, we are already in state red. Reason: %v\n", e.payload)
	}

	// ---- Transitions and action in green ----
	sm.transitions[green] = map[Event]state{
		{Type: timerElapsed}: yellow,
	}

	sm.actions[green] = map[Event]action{
		{Type: timerElapsed}: greenState,
	}

	// ---- Transitions and action in yellow ----
	sm.transitions[yellow] = map[Event]state{
		{Type: timerElapsed}: red,
	}

	sm.actions[yellow] = map[Event]action{
		{Type: timerElapsed}: yellowState,
	}

	return sm
}

// Event sender — thread-safe and controlled
func (f *stateMachine) Send(event Event) error {
	select {
	case f.eventChan <- event:
		return nil
	case <-f.doneChan:
		return fmt.Errorf("FSM is stopped")
	}
}

// FSM runner — with optional context for shutdown
func (sm *stateMachine) Run(ctx context.Context) {
	for {
		select {
		case event, ok := <-sm.eventChan:
			if !ok {
				sm.shutdown()
				return
			}

			fmt.Printf("Event is: %v", event)
			newState, ok := sm.transitions[sm.currentState][event]
			if ok {
				sm.currentState = newState
				action, ok := sm.actions[sm.currentState][event]
				if ok {
					action(event)
				}
			} else {
				fmt.Println("Invalid transition")
			}

		case <-ctx.Done():
			fmt.Println("FSM received context cancellation")
			sm.shutdown()
			return
		}
	}
}

// Graceful shutdown
func (f *stateMachine) Close() {
	close(f.eventChan)
}

func (f *stateMachine) Done() <-chan struct{} {
	return f.doneChan
}

func (f *stateMachine) shutdown() {
	close(f.doneChan)
}
