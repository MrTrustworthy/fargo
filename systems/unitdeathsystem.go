package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitDeathSystem struct {
	*ecs.World
	dyingUnit *entities.Unit
}

func (uds *UnitDeathSystem) New(world *ecs.World) {
	uds.World = world
	engo.Mailbox.Listen(events.UNIT_DEATH_EVENT, uds.getHandleUnitDeathEvent())
}

func (uds *UnitDeathSystem) getHandleUnitDeathEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		udmsg, ok := msg.(events.UnitDeathEvent)
		if !ok {
			return
		}

		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
		udmsg.GetAnimationComponent().SelectAnimationByName("dead")
		uds.dyingUnit = udmsg.Unit
	}
}

func (uds *UnitDeathSystem) removeDyingUnit() {
	RemoveFromRenderSystem(uds.World, uds.dyingUnit)
	RemoveFromAnimationSystem(uds.World, uds.dyingUnit)
	RemoveFromSelectionSystem(uds.World, uds.dyingUnit)
}

func (uds *UnitDeathSystem) Update(dt float32) {
	if uds.dyingUnit == nil {
		return
	}

	animation := uds.dyingUnit.GetAnimationComponent().CurrentAnimation
	if animation != nil && animation.Name != "dead" {
		uds.removeDyingUnit()
		uds.dyingUnit = nil
	}
}

func (uds *UnitDeathSystem) Remove(e ecs.BasicEntity) {}
