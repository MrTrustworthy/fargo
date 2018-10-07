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
	events.Mailbox.Listen(events.TEST_COLLISON_EVENT, sts.getHandleCollisionEvent())

}

func (sts *SimulationTestSystem) getHandleSimpleAttackEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		if _, ok := msg.(events.TestBasicAttackEvent); !ok {
			return
		}

		posA, posB := engo.Point{X: 200, Y: 100}, engo.Point{X: 500, Y: 100}

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

		posA, posB := engo.Point{X: 200, Y: 250}, engo.Point{X: 500, Y: 250}

		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: posA})
		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: posB})

		unitA, _ := FindUnitUnderMouse(sts.World, posA)

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
						events.Mailbox.Dispatch(events.TestCollisionEvent{})
					})

					events.Mailbox.Dispatch(events.RequestLootPickupEvent{ // Loot pickup request
						Unit:     unitA,
						Lootpack: lootMsg.Lootpack,
					})
				})
				events.Mailbox.Dispatch(events.InputInteractEvent{Point: posB}) // Third Attack
			})
			events.Mailbox.Dispatch(events.InputInteractEvent{Point: posB}) // Second attack
		})

	}
}

func (sts *SimulationTestSystem) getHandleCollisionEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		if _, ok := msg.(events.TestCollisionEvent); !ok {
			return
		}

		posA, posB, goal := engo.Point{X: 200, Y: 400}, engo.Point{X: 500, Y: 400}, engo.Point{X: 800, Y: 400}

		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: posA})
		events.Mailbox.Dispatch(events.InputCreateunitEvent{Point: posB})

		unitA, _ := FindUnitUnderMouse(sts.World, posA)
		unitB, _ := FindUnitUnderMouse(sts.World, posB)

		events.Mailbox.Dispatch(events.InputSelectEvent{Point: posA})
		events.Mailbox.Dispatch(events.InputInteractEvent{Point: goal})

		events.Mailbox.ListenOnce(events.MOVEMENT_COMPLETED_EVENT_NAME, func(msg events.BaseEvent) {
			centerA, centerB := unitA.Center(), unitB.Center()
			distance := centerA.PointDistance(centerB)
			Assert(distance < 70 && distance >= 64, "Unit A should be stuck")
			fmt.Println("Test 3: Unit is stuck now")

			// now that we're stuck, let's try to move back to the origin
			events.Mailbox.ListenOnce(events.MOVEMENT_COMPLETED_EVENT_NAME, func(msg events.BaseEvent) {
				Assert(unitA.Center() == posA, "Unit A should have made it back to its origin")
				fmt.Println("Test 3: getHandleCollisionEvent passed")
			})
			events.Mailbox.Dispatch(events.InputInteractEvent{Point: posA})

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
