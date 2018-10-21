package eventsystem

type BaseEvent interface {
	Type() string
	AsLogMessage() string
}
