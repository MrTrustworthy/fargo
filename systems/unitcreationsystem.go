package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
)

type UnitCreationSystem struct {
	*ecs.World
}

func (ucs *UnitCreationSystem) New(world *ecs.World) {
	ucs.World = world
	engo.Mailbox.Listen(INPUT_EVENT_NAME, ucs.getHandleInputEvent())
}

func (ucs *UnitCreationSystem) Update(dt float32) {}

func (ucs *UnitCreationSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(InputEvent)
		if !ok || imsg.Action != "CreateUnit" {
			return
		}
		ucs.createRandomUnit(imsg.MouseTracker)
	}
}

func (ucs *UnitCreationSystem) Remove(e ecs.BasicEntity) {}

func (ucs *UnitCreationSystem) createRandomUnit(tracker MouseTracker) {
	unit := entities.NewUnit(&engo.Point{
		X: tracker.MouseX - entities.UNITSIZE/2,
		Y: tracker.MouseY - entities.UNITSIZE/2,
	})
	fmt.Println("name of the new unit is", unit.Name)
	AddToRenderSystem(ucs.World, unit)
	AddToAnimationSystem(ucs.World, unit)
	AddToSelectionSystem(ucs.World, unit)
}
