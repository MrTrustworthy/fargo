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
	currentCreateUnitEvent *events.InputCreateunitEvent
}

func (ucs *UnitCreationSystem) New(world *ecs.World) {
	ucs.World = world
	eventsystem.Mailbox.Listen(events.INPUT_CREATEUNIT_EVENT_NAME, ucs.getHandleInputEvent())
}

func (ucs *UnitCreationSystem) Update(dt float32) {
	if ucs.currentCreateUnitEvent != nil {
		ucs.createRandomUnit(*ucs.currentCreateUnitEvent)
		ucs.currentCreateUnitEvent = nil
	}
}

func (ucs *UnitCreationSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.InputCreateunitEvent)
		if !ok {
			return
		}
		if ucs.currentCreateUnitEvent != nil {
			fmt.Println("WARNING: Trying to add CreateUnitEvent even though there is already one pending")
			return
		}
		ucs.currentCreateUnitEvent = &imsg
	}
}

func (ucs *UnitCreationSystem) Remove(e ecs.BasicEntity) {}

func (ucs *UnitCreationSystem) createRandomUnit(msg events.InputCreateunitEvent) {
	unit := entities.NewUnit()
	unit.SetCenter(msg.Point)
	fmt.Println("name of the new unit is", unit.Name)
	AddToRenderSystem(ucs.World, unit)
	AddToAnimationSystem(ucs.World, unit)
	AddToSelectionSystem(ucs.World, unit)
}
