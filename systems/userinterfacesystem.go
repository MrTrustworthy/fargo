package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/entities"
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
	engo.Mailbox.Listen(SELECT_EVENT_NAME, uis.getHandleSelectEvent())

}

func (uis *UserInterfaceSystem) getHandleSelectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(SelectionEvent)
		if !ok {
			return
		}
		if imsg.Action == "Select" {
			uis.SelectText.SetText("Unit:" + imsg.Unit.Name)
		} else if imsg.Action == "Deselect" {
			uis.SelectText.SetText("Unit: None")
		}
	}
}

func (uis *UserInterfaceSystem) Update(dt float32) {
}

func (uis *UserInterfaceSystem) Remove(e ecs.BasicEntity) {}
