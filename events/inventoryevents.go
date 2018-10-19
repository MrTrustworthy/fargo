package events

import "github.com/MrTrustworthy/fargo/entities"

const (
	INVENTORY_SHOW_EVENT = "ShowInventoryEvent"

)

type ShowInventoryEvent struct {
	*entities.Unit
}

func (ce ShowInventoryEvent) Type() string { return INVENTORY_SHOW_EVENT }

func (ce ShowInventoryEvent) AsLogMessage() string {
	return "Showing inventory for selected unit"
}