package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
	"strconv"
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
	engo.Mailbox.Listen(events.SELECTION_SELECTED_EVENT_NAME, uis.getHandleSelectEvent())
	engo.Mailbox.Listen(events.SELECTION_DESELECTED_EVENT_NAME, uis.getHandleDeselectEvent())
}

func (uis *UserInterfaceSystem) getHandleSelectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.SelectionSelectedEvent)
		if !ok {
			return
		}
		unitText := "Unit: " + imsg.Unit.Name + "\nSelected Attack: " + imsg.Unit.SelectedAbility.Name() + "\nSpeed:" +
			strconv.Itoa(int(imsg.Unit.Speed)) + " HP: " + strconv.Itoa(imsg.Unit.Health) + " AP: " + strconv.Itoa(imsg.Unit.AP)
		uis.SelectText.SetText(unitText)
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

func (uis *UserInterfaceSystem) Update(dt float32) {
}

func (uis *UserInterfaceSystem) Remove(e ecs.BasicEntity) {}
