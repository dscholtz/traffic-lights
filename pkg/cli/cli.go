package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dscholtz/traffic-lights/pkg/fsm"
)

type stateMachine interface {
	Send(event fsm.Event) error
}

func StartCli(sm stateMachine) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an event: 'timer' or 'emergency <reason>'")

	for scanner.Scan() {
		text := scanner.Text()
		if strings.TrimSpace(text) == "" {
			continue
		}

		switch true {
		case strings.EqualFold(text, "timer"):
			sm.Send(fsm.NewTimerEvent())

		case strings.HasPrefix(strings.ToLower(text), "emergency"):
			parts := strings.SplitN(text, " ", 2)
			reason := "unknown"
			if len(parts) > 1 {
				reason = parts[1]
			}
			sm.Send(fsm.NewEmergencyEvent(reason))

		default:
			fmt.Println("Unrecognized input. Use 'timer' or 'emergency'")
		}
	}
}
