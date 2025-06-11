package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dscholtz/traffic-lights/pkg/fsm"
)

func main() {
	// Create a cancellable context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize the FSM with the initial state (Red)
	machine := fsm.NewFSM()

	// Start the FSM loop
	go machine.Run(ctx)

	// Send TIMER_ELAPSED events every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Shut down after 12 seconds
	go func() {
		time.Sleep(12 * time.Second)
		fmt.Println("Main: initiating shutdown")
		machine.Close() // closes event channel
		cancel()        // signals context cancellation (optional)
	}()

	// Event loop
	for {
		select {
		case <-ticker.C:
			err := machine.Send(fsm.NewTimerEvent())
			if err != nil {
				fmt.Println("Main: send failed -", err)
			}

		case <-machine.Done():
			fmt.Println("Main: FSM terminated")
			return
		}
	}
}
