package events

const (
	DIALOG_SHOW_EVENT = "DialogShowEvent"
	DIALOG_HIDE_EVENT = "DialogHideEvent"

)

type DialogShowEvent struct {}

func (ce DialogShowEvent) Type() string { return DIALOG_SHOW_EVENT }

func (ce DialogShowEvent) AsLogMessage() string {
	return "Showing dialog"
}

type DialogHideEvent struct {}

func (ce DialogHideEvent) Type() string { return DIALOG_HIDE_EVENT }

func (ce DialogHideEvent) AsLogMessage() string {
	return "Hiding dialog"
}