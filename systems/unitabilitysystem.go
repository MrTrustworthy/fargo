package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitAbilitySystem struct {
	world *ecs.World
}

func (uas *UnitAbilitySystem) New(world *ecs.World) {
	uas.world = world
	engo.Mailbox.Listen(events.ABILITY_REQUESTUSE_EVENT_NAME, uas.getHandleRequestAbilityEvent())

}

func (uas *UnitAbilitySystem) getHandleRequestAbilityEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		ramsg, ok := msg.(events.RequestAbilityUseEvent)
		if !ok {
			return
		}
		// TODO check if actually in range
		// TODO select animation name dynamically or dispatch via Ability.Execute()
		ramsg.Source.AnimationComponent.SelectAnimationByName("stab")
	}
}

func (uas *UnitAbilitySystem) Update(dt float32) {}

func (uas *UnitAbilitySystem) Remove(e ecs.BasicEntity) {}
