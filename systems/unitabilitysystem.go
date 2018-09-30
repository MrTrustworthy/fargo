package systems

import (
	"engo.io/ecs"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
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
		ramsg, ok := msg.(events.RequestAbilityUseEvent)
		if !ok {
			return
		}
		if MovementIsCurrentlyProcessing(uas.world) {
			// Can't start attack as long as movement is still ongoing
			fmt.Println("Can't start attack since movement is still in progress")
			return
		}

		source, target := (*ramsg.Ability).Source(), (*ramsg.Ability).Target()
		sourcePosition := source.GetSpaceComponent().Center()
		currentDistance := sourcePosition.PointDistance(target.GetSpaceComponent().Center())
		// TODO maybe only pass the event here in both cases like for lootmanagement as well?
		if currentDistance <= source.SelectedAbility.Maxrange() {
			executeAbility(*ramsg.Ability)
		} else {
			fmt.Println("Can't attack, distance too great:", currentDistance, "trying again")
			moveCloserAndRetryAbility(source, target)
		}
	}
}

func executeAbility(ability entities.Ability) {
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

func moveCloserAndRetryAbility(originUnit, targetUnit *entities.Unit) {
	// TODO handle cases where a current movement is ongoing and no new movement is started,
	// TODO but the ability use is still queued
	events.Mailbox.ListenOnce(events.MOVEMENT_COMPLETED_EVENT_NAME, func(msg events.BaseEvent) {
		dispatchAttackUnit(originUnit, targetUnit)
	})
	dispatchMoveTo(targetUnit.Center().X, targetUnit.Center().Y, originUnit.SelectedAbility.Maxrange())
}

func dispatchAttackUnit(originUnit, targetUnit *entities.Unit) {
	originUnit.SelectedAbility.SetTarget(targetUnit)
	events.Mailbox.Dispatch(events.RequestAbilityUseEvent{
		Ability: &originUnit.SelectedAbility,
	})
}

func (uas *UnitAbilitySystem) Update(dt float32) {}

func (uas *UnitAbilitySystem) Remove(e ecs.BasicEntity) {}
