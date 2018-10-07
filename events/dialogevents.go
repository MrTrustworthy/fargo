package events

import "engo.io/engo"

const (
	DIALOG_SHOW_INVENTORY_EVENT = "DialogShowInventoryEvent"
	DIALOG_HIDE_EVENT           = "DialogHideEvent"
	DIALOG_CLICK_EVENT          = "DialogClickEvent"
)

type DialogShowInventoryEvent struct {}

func (ce DialogShowInventoryEvent) Type() string { return DIALOG_SHOW_INVENTORY_EVENT }

func (ce DialogShowInventoryEvent) AsLogMessage() string {
	return "Showing dialog"
}

type DialogHideEvent struct {}

func (ce DialogHideEvent) Type() string { return DIALOG_HIDE_EVENT }

func (ce DialogHideEvent) AsLogMessage() string {
	return "Hiding dialog"
}

type DialogClickEvent struct {
	engo.Point
}

func (ce DialogClickEvent) Type() string { return DIALOG_CLICK_EVENT }

func (ce DialogClickEvent) AsLogMessage() string {
	return "Clicked on dialog"
}