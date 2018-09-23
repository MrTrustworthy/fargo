package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitCollisionSystem struct {
	world *ecs.World
}

func (ucs *UnitCollisionSystem) New(world *ecs.World) {
	ucs.world = world
	engo.Mailbox.Listen(events.MOVEMENT_EVENT_NAME, ucs.getHandleMoveStepEvent())
}

func (ucs *UnitCollisionSystem) getHandleMoveStepEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		mmsg, ok := msg.(events.MovementEvent)
		if !ok || mmsg.Action != events.MOVEMENT_EVENT_ACTION_STEP {
			return
		}

		unitHitbox := mmsg.Unit.SpaceComponent.AABB()

		for _, other := range GetAllExistingUnits(ucs.world) {
			// exclude collisions with the unit itself
			if mmsg.Unit == other {
				continue
			}

			otherHitbox := other.SpaceComponent.AABB()
			if !common.IsIntersecting(unitHitbox, otherHitbox) {
				continue
			}
			engo.Mailbox.Dispatch(events.CollisionEvent{
				ActiveUnit:  mmsg.Unit,
				PassiveUnit: other,
				Action:      events.COLLISON_EVENT_ACTION_COLLIDED,
			})
			return
		}
	}
}

func (ucs *UnitCollisionSystem) Update(dt float32) {}

func (ucs *UnitCollisionSystem) Remove(e ecs.BasicEntity) {}
