package fsm

import (
	"context"
	"fmt"
)

type FSM struct {
	currentState stateFn
	eventChan    chan event
	doneChan     chan struct{}
}

// Constructor
func NewFSM() *FSM {
	return &FSM{
		currentState: redState,
		eventChan:    make(chan event, 10), // buffered for responsiveness
		doneChan:     make(chan struct{}),
	}
}

// Event sender — thread-safe and controlled
func (f *FSM) Send(event event) error {
	select {
	case f.eventChan <- event:
		return nil
	case <-f.doneChan:
		return fmt.Errorf("FSM is stopped")
	}
}

// FSM runner — with optional context for shutdown
func (f *FSM) Run(ctx context.Context) {
	for {
		select {
		case event, ok := <-f.eventChan:
			if !ok {
				f.shutdown()
				return
			}
			if f.currentState == nil {
				f.shutdown()
				return
			}
			f.currentState = f.currentState(event)

		case <-ctx.Done():
			fmt.Println("FSM received context cancellation")
			f.shutdown()
			return
		}
	}
}

// Graceful shutdown
func (f *FSM) Close() {
	close(f.eventChan)
}

func (f *FSM) Done() <-chan struct{} {
	return f.doneChan
}

func (f *FSM) shutdown() {
	close(f.doneChan)
}
