package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
)

type CollisionSystem struct {
	world *ecs.World
}

func (cs *CollisionSystem) New(world *ecs.World) {
	cs.world = world
	engo.Mailbox.Listen(events.MOVEMENT_EVENT_NAME, cs.getHandleMoveStepEvent())
}


func (cs *CollisionSystem) getHandleMoveStepEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		mmsg, ok := msg.(events.MovementEvent)
		if !ok || mmsg.Action != events.MOVEMENT_EVENT_ACTION_STEP {
			return
		}
		fmt.Println("TODO: check for collisions")

	}
}


func (cs *CollisionSystem) Update(dt float32) {}

func (cs *CollisionSystem) Remove(e ecs.BasicEntity) {}
