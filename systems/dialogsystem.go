

package systems

import (
	"engo.io/ecs"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type DialogSystem struct {
	world *ecs.World
	currentDialog *entities.Dialog
}

func (ds *DialogSystem) New(world *ecs.World) {
	ds.world = world
	events.Mailbox.Listen(events.DIALOG_SHOW_EVENT, ds.getHandleShowDialog())
	events.Mailbox.Listen(events.DIALOG_HIDE_EVENT, ds.getHandleHideDialog())
}

func (ds *DialogSystem) getHandleShowDialog() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		_, ok := msg.(events.DialogShowEvent)
		if !ok {
			return
		}
		ds.currentDialog = entities.NewDialog()
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
