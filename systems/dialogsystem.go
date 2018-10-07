package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/ui"
)

type DialogSystem struct {
	world         *ecs.World
	currentDialog *ui.Dialog
}

func (ds *DialogSystem) New(world *ecs.World) {
	ds.world = world
	events.Mailbox.Listen(events.DIALOG_SHOW_INVENTORY_EVENT, ds.getHandleShowInventoryDialog())
	events.Mailbox.Listen(events.DIALOG_HIDE_EVENT, ds.getHandleHideDialog())
	events.Mailbox.Listen(events.DIALOG_CLICK_EVENT, ds.getHandleDialogClick())
}

func (ds *DialogSystem) getHandleShowInventoryDialog() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		_, ok := msg.(events.DialogShowInventoryEvent)
		if !ok {
			return
		}
		if ds.currentDialog != nil {
			ds.HideCurrentDialog()
		}
		ds.currentDialog = ui.NewInventoryDialog()
		ds.ShowCurrentDialog()
	}
}

func (ds *DialogSystem) getHandleHideDialog() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		_, ok := msg.(events.DialogHideEvent)
		if !ok {
			return
		}
		ds.HideCurrentDialog()
		ds.currentDialog = nil
	}
}

func (ds *DialogSystem) getHandleDialogClick() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		dce, ok := msg.(events.DialogClickEvent)
		if !ok {
			return
		}

		fmt.Println("CLICKED DIALOG")

		for _, elem := range ds.currentDialog.Elements {
			if clicker, ok := elem.(ui.Clicker); ok && elem.GetSpaceComponent().Contains(dce.Point){
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

func (ds *DialogSystem) HideCurrentDialog() {
	for _, render := range ds.currentDialog.Elements {
		RemoveFromRenderSystem(ds.world, render)
	}
}

func (ds *DialogSystem) Update(dt float32) {}

func (ds *DialogSystem) Remove(e ecs.BasicEntity) {}
