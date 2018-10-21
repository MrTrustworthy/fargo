package events

import (
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/ui"
)

const (
	DIALOG_SHOW_EVENT  = "DialogShowEvent"
	DIALOG_HIDE_EVENT  = "DialogHideEvent"
	DIALOG_CLICK_EVENT = "DialogClickEvent"
)

type DialogShowEvent struct {
	*ui.Dialog
}

func (ce DialogShowEvent) Type() string { return DIALOG_SHOW_EVENT }

func (ce DialogShowEvent) AsLogMessage() string {
	return "Showing dialog"
}

type DialogHideEvent struct{}

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
