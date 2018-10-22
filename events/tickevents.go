package events
const (
	TICK_EVENT = "TickEvent"
)

type TickEvent struct{}

func (se TickEvent) Type() string { return TICK_EVENT }

func (se TickEvent) AsLogMessage() string {
	return "Tick"
}