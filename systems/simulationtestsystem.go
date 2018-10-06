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
	events.Mailbox.Listen(events.TEST_KILL_AND_LOOT, sts.getHandleKillAndLootEvent())

}

func (sts *SimulationTestSystem) getHandleSimpleAttackEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		if _, ok := msg.(events.TestBasicAttackEvent); !ok {
			return
		}

		posA, posB := engo.Point{X: 100, Y: 100}, engo.Point{X: 600, Y: 100}

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

			events.Mailbox.Dispatch(events.TestKillAndLootEvent{})
		})

	}
}

func (sts *SimulationTestSystem) getHandleKillAndLootEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {

		if _, ok := msg.(events.TestKillAndLootEvent); !ok {
			return
		}

		posA, posB := engo.Point{X: 100, Y: 300}, engo.Point{X: 600, Y: 300}

		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: posA})
		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: posB})

		unitA, _ := FindUnitUnderMouse(sts.World, posA)
		//unitB, _ := FindUnitUnderMouse(sts.World, posB)

		events.Mailbox.Dispatch(events.InputSelectEvent{Point: posA})
		events.Mailbox.Dispatch(events.InputInteractEvent{Point: posB}) // First attack

		events.Mailbox.ListenOnce(events.ABILITY_COMPLETED_EVENT_NAME, func(msg events.BaseEvent) {
			fmt.Println("Test 2: First Attack Completed")
			events.Mailbox.ListenOnce(events.ABILITY_COMPLETED_EVENT_NAME, func(msg events.BaseEvent) {
				fmt.Println("Test 2: Second Attack Completed")
				events.Mailbox.ListenOnce(events.LOOT_HAS_SPAWNED_EVENT, func(msg events.BaseEvent) {
					fmt.Println("Test 2: Third Attack, Unit Death and Loot Spawn Completed")
					lootMsg, _ := msg.(events.LootHasSpawnedEvent)
					lootPos := lootMsg.Lootpack.SpaceComponent.Center()

					events.Mailbox.ListenOnce(events.UNIT_ATTRIBUTE_CHANGE_EVENT, func(msg events.BaseEvent) {
						fmt.Println("Test 2: Loot picked up")
						centerA := unitA.Center()
						Assert(centerA.PointDistance(lootPos) <= LOOT_PICKUP_DISTANCE, "Should be in distance for pickup")
						Assert(unitA.AP == 1, "Should have only 1 AP left")
						Assert(unitA.Inventory.Size() == 8, "Should have 8 items now")
						fmt.Println("Test 2: getHandleKillAndLootEvent passed")
					})

					events.Mailbox.Dispatch(events.RequestLootPickupEvent{ // Loot pickup request
						Unit: unitA,
						Lootpack: lootMsg.Lootpack,
					})
				})
				events.Mailbox.Dispatch(events.InputInteractEvent{Point: posB}) // Third Attack
			})
			events.Mailbox.Dispatch(events.InputInteractEvent{Point: posB}) // Second attack
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
