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
		selectedUnit := GetCurrentlySelectedUnit(uis.World)
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
			// TODO remove the "dispatchXXX" calls here, also directly send event with selected unit
			dispatchMoveTo(imsg.Point.X, imsg.Point.Y, 0)
		}
	}
}

func (uis *UnitInteractionSystem) Update(dt float32) {}

func (uis *UnitInteractionSystem) Remove(e ecs.BasicEntity) {}
