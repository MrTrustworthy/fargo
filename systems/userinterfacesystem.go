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
}

func (uis *UserInterfaceSystem) New(world *ecs.World) {
	uis.MainHUD = entities.NewHUD()
	AddToRenderSystem(world, uis.MainHUD)
	uis.SelectText = entities.NewHUDText()
	AddToRenderSystem(world, uis.SelectText)
	engo.Mailbox.Listen(events.SELECT_EVENT_NAME, uis.getHandleSelectEvent())

}

func (uis *UserInterfaceSystem) getHandleSelectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.SelectionEvent)
		if !ok {
			return
		}
		if imsg.Action == events.SELECT_EVENT_ACTION_SELECTED {
			uis.SelectText.SetText("Unit:" + imsg.Unit.Name)
		} else if imsg.Action == events.SELECT_EVENT_ACTION_DESELECT {
			uis.SelectText.SetText("Unit: None")
		}
	}
}

func (uis *UserInterfaceSystem) Update(dt float32) {
}

func (uis *UserInterfaceSystem) Remove(e ecs.BasicEntity) {}
