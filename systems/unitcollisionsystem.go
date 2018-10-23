package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
)

type UnitCollisionSystem struct {
	world *ecs.World
}

func (ucs *UnitCollisionSystem) New(world *ecs.World) {
	ucs.world = world
	eventsystem.Mailbox.Listen(events.MOVEMENT_STEP_EVENT_NAME, ucs.getHandleMoveStepEvent())
}

// The collision systems works as follows: Each step of a movement, the moving unit is checked against all other units.
// If a collision is detected, a CollisionEvent is sent. In that case, the MovementSystem is responsible for handling
// the collision by cancelling the movement and resetting the unit to its last known good position.
func (ucs *UnitCollisionSystem) getHandleMoveStepEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		mmsg, ok := msg.(events.MovementStepEvent)
		if !ok {
			return
		}

		unitPosition := mmsg.Unit.Center()

		for _, other := range GetAllExistingUnits(ucs.world) {
			// exclude collisions with the unit itself
			if mmsg.Unit == other {
				continue
			}

			otherPosition := other.Center()
			distance := unitPosition.PointDistance(otherPosition)
			collisionDistance := (float32(mmsg.Unit.HitboxSize) / 2.0) + (float32(other.HitboxSize) / 2.0)
			if distance > collisionDistance {
				continue
			}

			eventsystem.Mailbox.Dispatch(events.CollisionEvent{
				ActiveUnit:  mmsg.Unit,
				PassiveUnit: other,
			})
			return
		}
	}
}

func (ucs *UnitCollisionSystem) Update(dt float32) {}

func (ucs *UnitCollisionSystem) Remove(e ecs.BasicEntity) {}
