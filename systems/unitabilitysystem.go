package systems

import (
	"engo.io/ecs"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitAbilitySystem struct {
	world            *ecs.World
	executingAbility entities.Ability
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
		if WorldIsCurrentlyBusy(uas.world) {
			// Can't start attack as long as movement is still ongoing
			fmt.Println("Can't start attack since movement is still in progress")
			return
		}

		source, target := (*raue.Ability).Source(), (*raue.Ability).Target()
		sourcePosition := source.GetSpaceComponent().Center()
		currentDistance := sourcePosition.PointDistance(target.GetSpaceComponent().Center())

		if currentDistance <= source.SelectedAbility.Maxrange() {
			uas.executeAbility(&raue)
		} else {
			fmt.Println("Can't attack, distance too great:", currentDistance, "trying again")
			moveCloserAndRetryAbility(&raue)
		}
	}
}

func (uas *UnitAbilitySystem) executeAbility(raue *events.RequestAbilityUseEvent) {
	ability := *raue.Ability
	if !ability.CanExecute() {
		events.Mailbox.Dispatch(events.AbilityCompletedEvent{Ability: &ability, Successful: false})
		return
	}
	uas.executingAbility = ability
	ability.Source().AnimationComponent.SelectAnimationByName(ability.AnimationName())

}

func moveCloserAndRetryAbility(raue *events.RequestAbilityUseEvent) {
	source, target := (*raue.Ability).Source(), (*raue.Ability).Target()
	events.Mailbox.ListenOnce(events.MOVEMENT_COMPLETED_EVENT_NAME, func(msg events.BaseEvent) {
		if cmsg, ok := msg.(events.MovementCompletedEvent); ok && cmsg.Successful {
			events.Mailbox.Dispatch(*raue)
		} else {
			events.Mailbox.Dispatch(events.AbilityCompletedEvent{Ability: raue.Ability, Successful: false})
		}
	})
	events.Mailbox.Dispatch(events.MovementRequestEvent{
		Target:         target.Center(),
		StopAtDistance: source.SelectedAbility.Maxrange(),
		Unit:           source,
	})
}

func (uas *UnitAbilitySystem) Update(dt float32) {
	// no ability in progress
	if uas.executingAbility == nil {
		return
	}

	// ability is still in progress, not yet completed
	animation := uas.executingAbility.Source().GetAnimationComponent().CurrentAnimation
	if animation == nil || animation.Name == uas.executingAbility.AnimationName() {
		return
	}

	// actually execute the results of the ability
	ability := uas.executingAbility
	ability.Source().AP -= ability.Cost()
	events.Mailbox.Dispatch(events.UnitAttributesChangedEvent{Unit: ability.Source()})

	events.Mailbox.Dispatch(events.RequestUnitDamageEvent{
		Unit:   ability.Target(),
		Damage: ability.Damage(),
	})

	if ability.Target().Health <= 0 {
		events.Mailbox.Dispatch(events.UnitDeathEvent{
			Unit: ability.Target(),
		})
	}
	uas.executingAbility = nil
	events.Mailbox.Dispatch(events.AbilityCompletedEvent{Ability: &ability, Successful: true})
}

func (uas *UnitAbilitySystem) IsIdle() bool {
	return uas.executingAbility == nil
}

func (uas *UnitAbilitySystem) Remove(e ecs.BasicEntity) {}
