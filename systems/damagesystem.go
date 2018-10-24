package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
)

type DamageSystem struct {
	currentDamageEvent *events.RequestUnitDamageEvent
}

func (ds *DamageSystem) New(world *ecs.World) {
	eventsystem.Mailbox.Listen(events.UNIT_REQUEST_DAMAGE_EVENT, ds.getHandleDamageEvent())
}
func (ds *DamageSystem) Update(dt float32) {
	if ds.currentDamageEvent != nil {
		ds.handleDamage(ds.currentDamageEvent)
		ds.currentDamageEvent = nil
	}
}

func (ds *DamageSystem) getHandleDamageEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.RequestUnitDamageEvent)
		if !ok {
			return
		}
		if ds.currentDamageEvent != nil {
			fmt.Println("WARNING: Trying to add RequestUnitDamageEvent even though there is already one pending")
			return
		}
		ds.currentDamageEvent = &imsg

	}
}

func (ds *DamageSystem) handleDamage(msg *events.RequestUnitDamageEvent) {
	msg.Unit.Health -= msg.Damage
	eventsystem.Mailbox.Dispatch(events.UnitAttributesChangedEvent{Unit: msg.Unit})
	if msg.Unit.Health <= 0 {
		eventsystem.Mailbox.Dispatch(events.UnitDeathEvent{
			Unit: msg.Unit,
		})
	}
}



func (ds *DamageSystem) Remove(e ecs.BasicEntity) {}
