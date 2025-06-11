package fsm

type eventType int

const (
	timerElapsed eventType = iota
	emergency
	powerOutage
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
