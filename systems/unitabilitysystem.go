package systems

import (
	"engo.io/ecs"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitAbilitySystem struct {
	world *ecs.World
}

func (uas *UnitAbilitySystem) New(world *ecs.World) {
	uas.world = world
	events.Mailbox.Listen(events.ABILITY_REQUESTUSE_EVENT_NAME, uas.getHandleRequestAbilityEvent())

}

func (uas *UnitAbilitySystem) getHandleRequestAbilityEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		raue, ok := msg.(events.RequestAbilityUseEvent)
		if !ok {
			return
		}
		if MovementIsCurrentlyProcessing(uas.world) {
			// Can't start attack as long as movement is still ongoing
			fmt.Println("Can't start attack since movement is still in progress")
			return
		}

		source, target := (*raue.Ability).Source(), (*raue.Ability).Target()
		sourcePosition := source.GetSpaceComponent().Center()
		currentDistance := sourcePosition.PointDistance(target.GetSpaceComponent().Center())

		if currentDistance <= source.SelectedAbility.Maxrange() {
			executeAbility(&raue)
		} else {
			fmt.Println("Can't attack, distance too great:", currentDistance, "trying again")
			moveCloserAndRetryAbility(&raue)
		}
	}
}

func executeAbility(raue *events.RequestAbilityUseEvent) {
	ability := *raue.Ability
	if !ability.CanExecute() {
		events.Mailbox.Dispatch(events.AbilityAbortedEvent{Ability: &ability})
		return
	}

	ability.Source().AP -= ability.Cost()
	events.Mailbox.Dispatch(events.UnitAttributesChangedEvent{Unit: ability.Source()})
	ability.Source().AnimationComponent.SelectAnimationByName(ability.AnimationName())

	events.Mailbox.Dispatch(events.RequestUnitDamageEvent{
		Unit:   ability.Target(),
		Damage: ability.Damage(),
	})

	if ability.Target().Health <= 0 {
		events.Mailbox.Dispatch(events.UnitDeathEvent{
			Unit: ability.Target(),
		})
	}
	events.Mailbox.Dispatch(events.AbilityCompletedEvent{Ability: &ability})
}

func moveCloserAndRetryAbility(raue *events.RequestAbilityUseEvent) {
	source, target := (*raue.Ability).Source(), (*raue.Ability).Target()
	events.Mailbox.ListenOnce(events.MOVEMENT_COMPLETED_EVENT_NAME, func(msg events.BaseEvent) {
		events.Mailbox.Dispatch(raue)
	})
	events.Mailbox.Dispatch(events.MovementRequestEvent{
		Target:         target.Center(),
		StopAtDistance: source.SelectedAbility.Maxrange(),
		Unit:           source,
	})
}

func (uas *UnitAbilitySystem) Update(dt float32) {}

func (uas *UnitAbilitySystem) Remove(e ecs.BasicEntity) {}
