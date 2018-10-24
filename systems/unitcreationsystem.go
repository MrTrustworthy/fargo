package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
)

type UnitCreationSystem struct {
	*ecs.World
	*TaskQueue
}

func (ucs *UnitCreationSystem) New(world *ecs.World) {
	ucs.World = world
	ucs.TaskQueue = &TaskQueue{}
	eventsystem.Mailbox.Listen(events.INPUT_CREATEUNIT_EVENT_NAME, ucs.getHandleInputEvent())
}

func (ucs *UnitCreationSystem) Update(dt float32) {
	ucs.TaskQueue.RunNext()
}

func (ucs *UnitCreationSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.InputCreateunitEvent)
		if !ok {
			return
		}
		ucs.AddTask(NewTask(createRandomUnit, ucs, imsg.Point))
	}
}

func (ucs *UnitCreationSystem) Remove(e ecs.BasicEntity) {}

func createRandomUnit(args ...interface{}) {
	ucs := args[0].(*UnitCreationSystem)
	point := args[1].(engo.Point)
	unit := entities.NewUnit()
	unit.SetCenter(point)
	fmt.Println("name of the new unit is", unit.Name)
	AddToRenderSystem(ucs.World, unit)
	AddToAnimationSystem(ucs.World, unit)
	AddToSelectionSystem(ucs.World, unit)
}
