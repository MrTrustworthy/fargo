package events

import "github.com/MrTrustworthy/fargo/entities"

const (
	INVENTORY_SHOW_EVENT = "ShowInventory"

)

type ShowInventory struct {
	*entities.Unit
}

func (ce ShowInventory) Type() string { return INVENTORY_SHOW_EVENT }

func (ce ShowInventory) AsLogMessage() string {
	return "Showing dialog"
}