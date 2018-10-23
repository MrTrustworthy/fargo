package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
)

type UnitDeathSystem struct {
	*ecs.World
	dyingUnit *entities.Unit
}

func (uds *UnitDeathSystem) New(world *ecs.World) {
	uds.World = world
	eventsystem.Mailbox.Listen(events.UNIT_DEATH_EVENT, uds.getHandleUnitDeathEvent())
}

func (uds *UnitDeathSystem) getHandleUnitDeathEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		udmsg, ok := msg.(events.UnitDeathEvent)
		if !ok {
			return
		}

		udmsg.GetAnimationComponent().SelectAnimationByName("dead")
		uds.dyingUnit = udmsg.Unit
	}
}

func (uds *UnitDeathSystem) Update(dt float32) {
	if uds.dyingUnit == nil {
		return
	}

	animation := uds.dyingUnit.GetAnimationComponent().CurrentAnimation
	if animation != nil && animation.Name != "dead" {
		lootPosition := uds.dyingUnit.GetSpaceComponent().Center()
		eventsystem.Mailbox.Dispatch(events.UnitRemovalEvent{Unit: uds.dyingUnit})
		uds.dyingUnit = nil
		eventsystem.Mailbox.Dispatch(events.RequestLootSpawn{Point: lootPosition})
	}
}

func (uds *UnitDeathSystem) Remove(e ecs.BasicEntity) {}
