package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
	"github.com/MrTrustworthy/fargo/ui"
)

type DialogSystem struct {
	world         *ecs.World
	currentDialog *ui.Dialog
}

func (ds *DialogSystem) New(world *ecs.World) {
	ds.world = world
	eventsystem.Mailbox.Listen(events.DIALOG_SHOW_EVENT, ds.getHandleShowDialog())
	eventsystem.Mailbox.Listen(events.DIALOG_HIDE_EVENT, ds.getHandleHideDialog())
	eventsystem.Mailbox.Listen(events.SELECTION_DESELECTED_EVENT_NAME, ds.getHandleHideDialog())
	eventsystem.Mailbox.Listen(events.DIALOG_CLICK_EVENT, ds.getHandleDialogClick())
}

func (ds *DialogSystem) getHandleShowDialog() func(msg eventsystem.BaseEvent) {
	return func(msg eventsystem.BaseEvent) {
		dsmsg, ok := msg.(events.DialogShowEvent)
		if !ok {
			return
		}
		ds.EnsureCurrentDialogClosed()
		ds.currentDialog = dsmsg.Dialog
		ds.ShowCurrentDialog()
		eventsystem.Mailbox.Dispatch(events.DialogShownEvent{Dialog: ds.currentDialog})
	}
}

func (ds *DialogSystem) getHandleHideDialog() func(msg eventsystem.BaseEvent) {
	return func(msg eventsystem.BaseEvent) {
		ds.EnsureCurrentDialogClosed()
		ds.currentDialog = nil
	}
}

func (ds *DialogSystem) getHandleDialogClick() func(msg eventsystem.BaseEvent) {
	return func(msg eventsystem.BaseEvent) {
		dce, ok := msg.(events.DialogClickEvent)
		if !ok {
			return
		}

		for _, elem := range ds.currentDialog.Elements {
			if clicker, ok := elem.(ui.Clicker); ok && elem.GetSpaceComponent().Contains(dce.Point) {
				clicker.HandleClick()
				return
			}
		}

	}
}

func (ds *DialogSystem) HasDialogAtPosition(point engo.Point) bool {
	// TODO does this work when the camera has moved and the mouse & window positions no longer match up?
	return ds.currentDialog != nil && ds.currentDialog.Background.SpaceComponent.Contains(point)
}

func (ds *DialogSystem) ShowCurrentDialog() {
	for _, render := range ds.currentDialog.Elements {
		AddToRenderSystem(ds.world, render)
	}
}

func (ds *DialogSystem) EnsureCurrentDialogClosed() {
	if ds.currentDialog == nil {
		return
	}
	for _, render := range ds.currentDialog.Elements {
		RemoveFromRenderSystem(ds.world, render)
	}
}

func (ds *DialogSystem) Update(dt float32) {}

func (ds *DialogSystem) Remove(e ecs.BasicEntity) {}
