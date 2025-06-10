package fsm

type eventType string

const (
	timerElapsed eventType = "TIMER_ELAPSED"
	emergency    eventType = "EMERGENCY"
	powerOutage  eventType = "POWER_OUTAGE"
)

type event struct {
	Type    eventType
	payload any // Payload to carry context such as elapsed time or something
}

func newTimerEvent() event {
	return event{Type: timerElapsed}
}

func newEmergencyEvent(reason string) event {
	return event{Type: emergency, payload: reason}
}
