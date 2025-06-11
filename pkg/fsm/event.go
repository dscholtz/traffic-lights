package fsm

type eventType string

const (
	timerElapsed eventType = "TIMER_ELAPSED"
	emergency    eventType = "EMERGENCY"
	powerOutage  eventType = "POWER_OUTAGE"
)

type Event struct {
	Type    eventType
	payload any // Payload to carry context such as elapsed time or something
}

func NewTimerEvent() Event {
	return Event{Type: timerElapsed}
}

func NewEmergencyEvent(reason string) Event {
	return Event{Type: emergency, payload: reason}
}
