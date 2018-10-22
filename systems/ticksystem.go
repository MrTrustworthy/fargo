package systems

import (
	"engo.io/ecs"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
)

type TickSystem struct {
	*ecs.World
	MouseTracker
}


func (is *TickSystem) Update(dt float32) {

	eventsystem.Mailbox.Dispatch(events.TickEvent{})

}

func (is *TickSystem) Remove(e ecs.BasicEntity) {}