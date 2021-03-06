package systems

import (
	"engo.io/ecs"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
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
	eventsystem.Mailbox.Listen(events.SELECTION_SELECTED_EVENT_NAME, uis.getHandleSelectEvent())
	eventsystem.Mailbox.Listen(events.SELECTION_DESELECTED_EVENT_NAME, uis.getHandleDeselectEvent())
	eventsystem.Mailbox.Listen(events.UNIT_ATTRIBUTE_CHANGE_EVENT, uis.getHandleAttributeChangeEvent())

}

func (uis *UserInterfaceSystem) getHandleSelectEvent() func(msg eventsystem.BaseEvent) {
	return func(msg eventsystem.BaseEvent) {
		imsg, ok := msg.(events.SelectionSelectedEvent)
		if !ok {
			return
		}
		uis.SelectText.SetDisplayeddUnit(imsg.Unit)
	}
}

func (uis *UserInterfaceSystem) getHandleDeselectEvent() func(msg eventsystem.BaseEvent) {
	return func(msg eventsystem.BaseEvent) {
		_, ok := msg.(events.SelectionDeselectedEvent)
		if !ok {
			return
		}
		uis.SelectText.SetText("Unit: None")
	}
}

func (uis *UserInterfaceSystem) getHandleAttributeChangeEvent() func(msg eventsystem.BaseEvent) {
	return func(msg eventsystem.BaseEvent) {
		imsg, ok := msg.(events.UnitAttributesChangedEvent)
		if !ok {
			return
		}
		uis.SelectText.UpdateTextForUnitIfDisplayed(imsg.Unit)
	}
}

func (uis *UserInterfaceSystem) Update(dt float32) {
}

func (uis *UserInterfaceSystem) Remove(e ecs.BasicEntity) {}
