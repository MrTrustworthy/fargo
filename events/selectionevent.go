package events

import "github.com/MrTrustworthy/fargo/entities"

const (
	SELECT_EVENT_NAME            = "SelectionEvent"
	SELECT_EVENT_ACTION_SELECTED = "Selected"
	SELECT_EVENT_ACTION_DESELECT = "Deselected"
)

type SelectionEvent struct {
	Action string
	*entities.Unit
}

func (se SelectionEvent) Type() string { return SELECT_EVENT_NAME }

func (se SelectionEvent) AsLogMessage() string {
	return "Action[" + se.Action + "] for unit " + se.Unit.Name
}
