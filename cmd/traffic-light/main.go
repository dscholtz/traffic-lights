package main

import (
	"context"

	"github.com/dscholtz/traffic-lights/pkg/cli"
	"github.com/dscholtz/traffic-lights/pkg/fsm"
)

func main() {
	// Create a cancellable context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize the FSM with the initial state (Red)
	machine := fsm.New()

	// Start the FSM loop
	go machine.Run(ctx)

	// Start cli sender
	cli.StartCli(machine)
}
