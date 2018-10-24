package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
)

type UnitInteractionSystem struct {
	*ecs.World
	currentInputEvent *events.InputInteractEvent
}

func (uis *UnitInteractionSystem) New(world *ecs.World) {
	uis.World = world
	eventsystem.Mailbox.Listen(events.INPUT_INTERACT_EVENT_NAME, uis.getHandleInputEvent())
}
func (uis *UnitInteractionSystem) Update(dt float32) {
	if uis.currentInputEvent != nil {
		uis.handleInput(uis.currentInputEvent)
		uis.currentInputEvent = nil
	}
}
func (uis *UnitInteractionSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.InputInteractEvent)
		if !ok {
			return
		}
		if uis.currentInputEvent != nil {
			fmt.Println("WARNING: Trying to add InputInteractEvent even though there is already one pending")
			return
		}
		uis.currentInputEvent = &imsg
	}
}

func (uis *UnitInteractionSystem) handleInput(msg *events.InputInteractEvent) {

	selectedUnit := msg.Unit
	if selectedUnit == nil {
		// there is nothing to do for this system when we don't have a selected unit
		return
	}

	if clickedUnit, err := FindUnitUnderMouse(uis.World, msg.Point); err == nil {
		if clickedUnit == selectedUnit {
			return
		}
		selectedUnit.SelectedAbility.SetTarget(clickedUnit)
		eventsystem.Mailbox.Dispatch(events.RequestAbilityUseEvent{
			Ability: &selectedUnit.SelectedAbility,
		})
	} else if clickedLootpack, err := FindLootUnderMouse(uis.World, msg.Point); err == nil {
		eventsystem.Mailbox.Dispatch(events.RequestLootPickupEvent{
			Unit:     selectedUnit,
			Lootpack: clickedLootpack,
		})
	} else {
		eventsystem.Mailbox.Dispatch(events.MovementRequestEvent{
			Target:         msg.Point,
			StopAtDistance: 0,
			Unit:           selectedUnit,
		})
	}

}

func (uis *UnitInteractionSystem) Remove(e ecs.BasicEntity) {}
