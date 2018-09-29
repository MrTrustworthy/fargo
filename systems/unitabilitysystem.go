package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
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
		if MovementIsCurrentlyProcessing(uas.world) {
			// Can't start attack as long as movement is still ongoing
			fmt.Println("Can't start attack since movement is still in progress")
			return
		}

		source, target := (*ramsg.Ability).Source(), (*ramsg.Ability).Target()
		sourcePosition := source.GetSpaceComponent().Center()
		currentDistance := sourcePosition.PointDistance(target.GetSpaceComponent().Center())

		if currentDistance <= source.SelectedAbility.Maxrange() {
			executeAbility(*ramsg.Ability)
		} else {
			fmt.Println("Can't attack, distance too great:", currentDistance, "trying again")
			moveCloserAndRetry(source, target)
		}
	}
}

func executeAbility(ability entities.Ability) {
	if !ability.CanExecute() {
		engo.Mailbox.Dispatch(events.AbilityAbortedEvent{Ability: &ability})
		return
	}

	ability.Source().AP -= ability.Cost()
	ability.Source().AnimationComponent.SelectAnimationByName(ability.AnimationName())

	engo.Mailbox.Dispatch(events.RequestUnitDamageEvent{
		Unit:   ability.Target(),
		Damage: ability.Damage(),
	})

	if ability.Target().Health <= 0 {
		engo.Mailbox.Dispatch(events.UnitDeathEvent{
			Unit: ability.Target(),
		})
	}
	engo.Mailbox.Dispatch(events.AbilityCompletedEvent{Ability: &ability})
}

func moveCloserAndRetry(originUnit, targetUnit *entities.Unit) {
	// TODO handle cases where a current movement is ongoing and no new movement is started,
	// TODO but the ability use is still queued
	events.ListenOnce(events.MOVEMENT_COMPLETED_EVENT_NAME, func(msg engo.Message) {
		dispatchAttackUnit(originUnit, targetUnit)
	})
	dispatchMoveTo(targetUnit.Center().X, targetUnit.Center().Y, originUnit.SelectedAbility.Maxrange())
}

func dispatchAttackUnit(originUnit, targetUnit *entities.Unit) {
	originUnit.SelectedAbility.SetTarget(targetUnit)
	engo.Mailbox.Dispatch(events.RequestAbilityUseEvent{
		Ability: &originUnit.SelectedAbility,
	})
}

func (uas *UnitAbilitySystem) Update(dt float32) {}

func (uas *UnitAbilitySystem) Remove(e ecs.BasicEntity) {}
