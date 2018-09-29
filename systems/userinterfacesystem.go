package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type UserInterfaceSystem struct {
	MainHUD    *entities.HUD
	SelectText *entities.HUDText
	*ecs.World
}

func (uis *UserInterfaceSystem) New(world *ecs.World) {
	uis.World = world
	uis.MainHUD = entities.NewHUD()
	AddToRenderSystem(uis.World, uis.MainHUD)
	uis.SelectText = entities.NewHUDText()
	AddToRenderSystem(uis.World, uis.SelectText)
	engo.Mailbox.Listen(events.SELECTION_SELECTED_EVENT_NAME, uis.getHandleSelectEvent())
	engo.Mailbox.Listen(events.SELECTION_DESELECTED_EVENT_NAME, uis.getHandleDeselectEvent())
	engo.Mailbox.Listen(events.UNIT_ATTRIBUTE_CHANGE_EVENT, uis.getHandleAttributeChangeEvent())

}

func (uis *UserInterfaceSystem) getHandleSelectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.SelectionSelectedEvent)
		if !ok {
			return
		}
		uis.SelectText.SetTextForUnit(imsg.Unit)
	}
}

func (uis *UserInterfaceSystem) getHandleDeselectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		_, ok := msg.(events.SelectionDeselectedEvent)
		if !ok {
			return
		}
		uis.SelectText.SetText("Unit: None")
	}
}

func (uis *UserInterfaceSystem) getHandleAttributeChangeEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.UnitAttributesChangedEvent)
		if !ok {
			return
		}
		if imsg.Unit != GetCurrentlySelectedUnit(uis.World) {
			return
		}
		uis.SelectText.SetTextForUnit(imsg.Unit)
	}
}

func (uis *UserInterfaceSystem) Update(dt float32) {
}

func (uis *UserInterfaceSystem) Remove(e ecs.BasicEntity) {}
