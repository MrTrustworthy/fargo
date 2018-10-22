package events

import (
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	LOOT_REQUEST_SPAWN_EVENT         = "RequestUnitDamageEvent"
	LOOT_REQUEST_PICKUP_EVENT        = "RequestLootPickupEvent"
	LOOT_REQUEST_ITEM_PICKUP_EVENT   = "RequestLootItemPickupEvent"
	LOOT_PICKUP_COMPLETED_EVENT      = "LootPickupCompletedEvent"
	LOOT_PICKUP_ITEM_COMPLETED_EVENT = "LootPickupItemCompletedEvent"
	LOOT_HAS_SPAWNED_EVENT           = "LootHasSpawnedEvent"
)

type RequestLootSpawn struct {
	engo.Point
}

func (se RequestLootSpawn) Type() string { return LOOT_REQUEST_SPAWN_EVENT }

func (se RequestLootSpawn) AsLogMessage() string {
	return fmt.Sprint("at point %.2f : %.2f", se.Point.X, se.Point.Y)
}

type RequestLootPickupEvent struct {
	*entities.Unit
	*entities.Lootpack
}

func (se RequestLootPickupEvent) Type() string { return LOOT_REQUEST_PICKUP_EVENT }

func (se RequestLootPickupEvent) AsLogMessage() string {
	return "Unit " + se.Unit.Name + " should pickup lootpack"
}

type LootPickupCompletedEvent struct {
	*entities.Lootpack
	Successful bool
}

func (se LootPickupCompletedEvent) Type() string { return LOOT_PICKUP_COMPLETED_EVENT }

func (se LootPickupCompletedEvent) AsLogMessage() string {
	s := "with a failure"
	if se.Successful {
		s = "Successfully"
	}
	return "Loot pickup completed " + s
}

type RequestLootItemPickupEvent struct {
	*entities.Unit
	*entities.Item
	*entities.Lootpack
}

func (se RequestLootItemPickupEvent) Type() string { return LOOT_REQUEST_ITEM_PICKUP_EVENT }

func (se RequestLootItemPickupEvent) AsLogMessage() string {
	return "Unit " + se.Unit.Name + " should pickup item " + se.Item.Name + " from lootpack"
}

type LootPickupItemCompletedEvent struct {
	*entities.Unit
	*entities.Item
	*entities.Lootpack
	Successful bool
}

func (se LootPickupItemCompletedEvent) Type() string { return LOOT_PICKUP_ITEM_COMPLETED_EVENT }

func (se LootPickupItemCompletedEvent) AsLogMessage() string {
	s := "with a failure"
	if se.Successful {
		s = "Successfully"
	}
	return "Loot item pickup completed " + s
}

type LootHasSpawnedEvent struct {
	*entities.Lootpack
}

func (se LootHasSpawnedEvent) Type() string { return LOOT_HAS_SPAWNED_EVENT }

func (se LootHasSpawnedEvent) AsLogMessage() string {
	return "Loot has spawned"
}
