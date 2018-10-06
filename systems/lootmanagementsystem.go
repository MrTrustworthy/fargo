package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"errors"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

const LOOT_PICKUP_DISTANCE = 20.0

type LootManagementSystem struct {
	*ecs.World
	ActiveLootPacks []*entities.Lootpack
}

func (lss *LootManagementSystem) New(world *ecs.World) {
	lss.World = world
	events.Mailbox.Listen(events.LOOT_REQUEST_SPAWN_EVENT, lss.getHandleRequestLootSpawn())
	events.Mailbox.Listen(events.LOOT_REQUEST_PICKUP_EVENT, lss.getHandleRequestLootPickup())

}

func (lss *LootManagementSystem) getHandleRequestLootSpawn() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		udmsg, ok := msg.(events.RequestLootSpawn)
		if !ok {
			return
		}

		lootpack := entities.NewLootpack(&udmsg.Point)
		lss.ActiveLootPacks = append(lss.ActiveLootPacks, lootpack)
		AddToRenderSystem(lss.World, lootpack)
	}
}

func (lss *LootManagementSystem) getHandleRequestLootPickup() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		udmsg, ok := msg.(events.RequestLootPickupEvent)
		if !ok {
			return
		}
		if MovementIsCurrentlyProcessing(lss.World) {
			// Can't start pickup as long as movement is still ongoing
			fmt.Println("Can't start attack since movement is still in progress")
			return
		}
		unitPosititon, packPosition := udmsg.Unit.Center(), udmsg.Lootpack.Center()
		currentDistance := unitPosititon.PointDistance(packPosition)
		if currentDistance <= LOOT_PICKUP_DISTANCE {
			lss.executePickup(&udmsg)
		} else {
			moveCloserAndRetryPickup(&udmsg)
		}

	}
}

func (lss *LootManagementSystem) executePickup(pickupEvent *events.RequestLootPickupEvent) {
	pickupEvent.Unit.Inventory.FillFrom(*pickupEvent.Lootpack.Inventory)
	lss.RemoveLootbox(pickupEvent.Lootpack)
	events.Mailbox.Dispatch(events.UnitAttributesChangedEvent{Unit: pickupEvent.Unit})
}

func moveCloserAndRetryPickup(pickupEvent *events.RequestLootPickupEvent) {
	events.Mailbox.ListenOnce(events.MOVEMENT_COMPLETED_EVENT_NAME, func(msg events.BaseEvent) {
		events.Mailbox.Dispatch(*pickupEvent)
	})
	events.Mailbox.Dispatch(events.MovementRequestEvent{
		Target:         pickupEvent.Lootpack.Center(),
		StopAtDistance: LOOT_PICKUP_DISTANCE,
		Unit:           pickupEvent.Unit,
	})
}

func (lss *LootManagementSystem) FindLootUnderMouse(point engo.Point) (*entities.Lootpack, error) {
	for _, pack := range lss.ActiveLootPacks {
		xDelta := point.X - pack.GetSpaceComponent().Position.X
		yDelta := point.Y - pack.GetSpaceComponent().Position.Y
		if xDelta > 0 && xDelta < entities.LOOTPACKSIZE && yDelta > 0 && yDelta < entities.LOOTPACKSIZE {
			return pack, nil
		}
	}
	return nil, errors.New("No lootpack Found")
}

func (lss *LootManagementSystem) RemoveLootbox(lootpack *entities.Lootpack) {
	index := -1
	for i, val := range lss.ActiveLootPacks {
		if val == lootpack {
			index = i
			break
		}
	}
	if index == -1 {
		panic("Trying to remove non-existing unit!")
	}
	lss.ActiveLootPacks = append(lss.ActiveLootPacks[:index], lss.ActiveLootPacks[index+1:]...)
	RemoveFromRenderSystem(lss.World, lootpack)
}

func (lss *LootManagementSystem) Update(dt float32) {}

func (lss *LootManagementSystem) Remove(e ecs.BasicEntity) {}
