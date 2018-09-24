package events

import "github.com/MrTrustworthy/fargo/entities"

const (
	SELECTION_SELECTED_EVENT_NAME   = "SelectionSelectedEvent"
	SELECTION_DESELECTED_EVENT_NAME = "SelectionDeselectedEvent"
)

type SelectionSelectedEvent struct {
	Action string
	*entities.Unit
}

func (se SelectionSelectedEvent) Type() string { return SELECTION_SELECTED_EVENT_NAME }

func (se SelectionSelectedEvent) AsLogMessage() string {
	return "for unit " + se.Unit.Name
}

type SelectionDeselectedEvent struct {
	Action string
	*entities.Unit
}

func (se SelectionDeselectedEvent) Type() string { return SELECTION_DESELECTED_EVENT_NAME }

func (se SelectionDeselectedEvent) AsLogMessage() string {
	return "for unit " + se.Unit.Name
}
