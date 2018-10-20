package events

import "github.com/MrTrustworthy/fargo/entities"

const (
	INVENTORY_SHOW_EVENT      = "ShowInventoryEvent"
	INVENTORY_ELEMENT_CLICKED = "InventoryElementClickedEvent"
)

type ShowInventoryEvent struct {
	*entities.Unit
}

func (ce ShowInventoryEvent) Type() string { return INVENTORY_SHOW_EVENT }

func (ce ShowInventoryEvent) AsLogMessage() string {
	return "Showing inventory for selected unit"
}

type InventoryElementClickedEvent struct {
	*entities.Unit
	*entities.Item
}

func (ce InventoryElementClickedEvent) Type() string { return INVENTORY_ELEMENT_CLICKED }

func (ce InventoryElementClickedEvent) AsLogMessage() string {
	return "Inventory element " + ce.Item.Name + " clicked"
}
