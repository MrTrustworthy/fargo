package systems

import (
	"engo.io/ecs"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitDeathSystem struct {
	*ecs.World
	dyingUnit *entities.Unit
}

func (uds *UnitDeathSystem) New(world *ecs.World) {
	uds.World = world
	events.Mailbox.Listen(events.UNIT_DEATH_EVENT, uds.getHandleUnitDeathEvent())
}

func (uds *UnitDeathSystem) getHandleUnitDeathEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		udmsg, ok := msg.(events.UnitDeathEvent)
		if !ok {
			return
		}

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
		lootPosition := uds.dyingUnit.GetSpaceComponent().Center()
		uds.removeDyingUnit()
		uds.dyingUnit = nil
		events.Mailbox.Dispatch(events.RequestLootSpawn{Point: lootPosition})
	}
}

func (uds *UnitDeathSystem) Remove(e ecs.BasicEntity) {}
