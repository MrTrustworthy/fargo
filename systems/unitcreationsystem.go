package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitCreationSystem struct {
	*ecs.World
}

func (ucs *UnitCreationSystem) New(world *ecs.World) {
	ucs.World = world
	engo.Mailbox.Listen(events.INPUT_EVENT_NAME, ucs.getHandleInputEvent())
}

func (ucs *UnitCreationSystem) Update(dt float32) {}

func (ucs *UnitCreationSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.InputEvent)
		if !ok || imsg.Action != "CreateUnit" {
			return
		}
		ucs.createRandomUnit(imsg.MouseTracker)
	}
}

func (ucs *UnitCreationSystem) Remove(e ecs.BasicEntity) {}

func (ucs *UnitCreationSystem) createRandomUnit(tracker events.MouseTracker) {
	unit := entities.NewUnit(&engo.Point{
		X: tracker.MouseX + entities.UNIT_CENTER_OFFSET.X,
		Y: tracker.MouseY + entities.UNIT_CENTER_OFFSET.Y,
	})
	fmt.Println("name of the new unit is", unit.Name)
	AddToRenderSystem(ucs.World, unit)
	AddToAnimationSystem(ucs.World, unit)
	AddToSelectionSystem(ucs.World, unit)
}
