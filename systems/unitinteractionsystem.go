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
	engo.Mailbox.Listen(events.INPUT_INTERACT_EVENT_NAME, uis.getHandleInputEvent())
}

func (uis *UnitInteractionSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.InputInteractEvent)
		if !ok {
			return
		}
		if clickedUnit, err := FindUnitUnderMouse(uis.World, &imsg.MouseTracker); err == nil {
			selectedUnit := GetCurrentlySelectedUnit(uis.World)
			if clickedUnit == selectedUnit {
				return
			}
			dispatchAttackUnit(selectedUnit, clickedUnit)
		} else {
			dispatchMoveTo(imsg.MouseX, imsg.MouseY, 0)
		}
	}
}

func (uis *UnitInteractionSystem) Update(dt float32) {}

func (uis *UnitInteractionSystem) Remove(e ecs.BasicEntity) {}
