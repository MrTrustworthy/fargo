package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
)

type SimulationTestSystem struct {
	*ecs.World
}

func (sts *SimulationTestSystem) New(world *ecs.World) {
	sts.World = world
	events.Mailbox.Listen(events.TEST_SIMPLE_ATTACK, sts.getHandleSimpleAttackEvent())
}

func (sts *SimulationTestSystem) getHandleSimpleAttackEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		if _, ok := msg.(events.TestBasicAttackEvent); !ok {
			return
		}

		posA, posB := engo.Point{X: 200, Y: 200}, engo.Point{X: 700, Y: 500}

		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: posA})
		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: posB})

		unitA, _ := FindUnitUnderMouse(sts.World, posA)
		unitB, _ := FindUnitUnderMouse(sts.World, posB)

		events.Mailbox.Dispatch(events.InputSelectEvent{Point: posA})
		events.Mailbox.Dispatch(events.InputInteractEvent{Point: posB})

		events.Mailbox.ListenOnce(events.ABILITY_COMPLETED_EVENT_NAME, func(msg events.BaseEvent) {
			Assert(unitA.AP == unitB.AP-unitA.SelectedAbility.Cost(), "Unit A should have lost AP")
			Assert(unitB.Health == unitA.Health-unitA.SelectedAbility.Damage(), "Unit B should have lost Health")
			centerA, centerB := unitA.Center(), unitB.Center()
			Assert(centerA.PointDistance(centerB) <= unitA.SelectedAbility.Maxrange(), "Unit A should be in Range")
			fmt.Println("Test 1: getHandleSimpleAttackEvent passed")
		})

	}
}

func Assert(testable bool, message string) {
	if !testable {
		panic(message)
	}
}

func (sts *SimulationTestSystem) Update(dt float32) {}

func (sts *SimulationTestSystem) Remove(e ecs.BasicEntity) {}
