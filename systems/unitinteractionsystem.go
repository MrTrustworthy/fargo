package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitInteractionSystem struct {
	*ecs.World
}

func (uis *UnitInteractionSystem) New(world *ecs.World) {
	uis.World = world
	engo.Mailbox.Listen(events.INPUT_EVENT_NAME, uis.getHandleInputEvent())
}

func (uis *UnitInteractionSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.InputEvent)
		if !ok || imsg.Action != events.INPUT_EVENT_ACTION_INTERACT {
			return
		}

		if clickedUnit, err := FindUnitUnderMouse(uis.World, &imsg.MouseTracker); err == nil {
			selectedUnit := GetCurrentlySelectedUnit(uis.World)
			if clickedUnit == selectedUnit {
				return
			} else {

				// ability use is requested

				// TODO handle cases where a current movement is ongoing and no new movement is started,
				// TODO but the ability use is still queued
				events.ListenOnce(events.MOVEMENT_EVENT_NAME, events.MOVEMENT_EVENT_ACTION_FINISHED, func(msg engo.Message) {
					engo.Mailbox.Dispatch(events.RequestAbilityUseEvent{
						Action:  events.REQUESTABILITYUSE_EVENT_ACTION_REQUEST_ABILITY,
						Source:  selectedUnit,
						Target:  clickedUnit,
						Ability: &selectedUnit.StandardAbility,
					})
				})
				dispatchMoveTo(clickedUnit.Center().X, clickedUnit.Center().Y, selectedUnit.StandardAbility.Maxrange())

			}
		} else {
			dispatchMoveTo(imsg.MouseX, imsg.MouseY, 0)
		}

	}
}

func dispatchMoveTo(x, y, dist float32) {
	engo.Mailbox.Dispatch(events.InteractionEvent{
		Target:         engo.Point{X: x, Y: y},
		Action:         events.INTERACTION_EVENT_ACTION_MOVE_TO,
		StopAtDistance: dist,
	})
}

func (uis *UnitInteractionSystem) Update(dt float32) {}

func (uis *UnitInteractionSystem) Remove(e ecs.BasicEntity) {}
