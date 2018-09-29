package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/events"
)

type DamageSystem struct {
}

func (ds *DamageSystem) New(world *ecs.World) {
	engo.Mailbox.Listen(events.UNIT_REQUEST_DAMAGE_EVENT, ds.getHandleDamageEvent())
}

func (ds *DamageSystem) getHandleDamageEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.RequestUnitDamageEvent)
		if !ok {
			return
		}

		imsg.Unit.Health -= imsg.Damage
		if imsg.Unit.Health <= 0 {
			engo.Mailbox.Dispatch(events.UnitDeathEvent{
				Unit: imsg.Unit,
			})
		}
	}
}

func (ds *DamageSystem) Update(dt float32) {}

func (ds *DamageSystem) Remove(e ecs.BasicEntity) {}
