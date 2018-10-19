package systems

import (
	"engo.io/ecs"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitInteractionSystem struct {
	*ecs.World
}

func (uis *UnitInteractionSystem) New(world *ecs.World) {
	uis.World = world
	events.Mailbox.Listen(events.INPUT_INTERACT_EVENT_NAME, uis.getHandleInputEvent())
}

func (uis *UnitInteractionSystem) getHandleInputEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		imsg, ok := msg.(events.InputInteractEvent)
		if !ok {
			return
		}
		selectedUnit := imsg.Unit
		if selectedUnit == nil {
			// there is nothing to do for this system when we don't have a selected unit
			return
		}

		if clickedUnit, err := FindUnitUnderMouse(uis.World, imsg.Point); err == nil {
			if clickedUnit == selectedUnit {
				return
			}
			selectedUnit.SelectedAbility.SetTarget(clickedUnit)
			events.Mailbox.Dispatch(events.RequestAbilityUseEvent{
				Ability: &selectedUnit.SelectedAbility,
			})
		} else if clickedLootpack, err := FindLootUnderMouse(uis.World, imsg.Point); err == nil {
			events.Mailbox.Dispatch(events.RequestLootPickupEvent{
				Unit:     selectedUnit,
				Lootpack: clickedLootpack,
			})
		} else {
			events.Mailbox.Dispatch(events.MovementRequestEvent{
				Target:         imsg.Point,
				StopAtDistance: 0,
				Unit:           selectedUnit,
			})
		}
	}
}

func (uis *UnitInteractionSystem) Update(dt float32) {}

func (uis *UnitInteractionSystem) Remove(e ecs.BasicEntity) {}
