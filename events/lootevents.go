package events

import (
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	LOOT_REQUEST_SPAWN_EVENT    = "RequestUnitDamageEvent"
	LOOT_REQUEST_PICKUP_EVENT   = "RequestLootPickupEvent"
	LOOT_PICKUP_COMPLETED_EVENT = "LootPickupCompletedEvent"
	LOOT_HAS_SPAWNED_EVENT      = "LootHasSpawnedEvent"
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

type LootHasSpawnedEvent struct {
	*entities.Lootpack
}

func (se LootHasSpawnedEvent) Type() string { return LOOT_HAS_SPAWNED_EVENT }

func (se LootHasSpawnedEvent) AsLogMessage() string {
	return "Loot has spawned"
}

type LootPickupCompletedEvent struct {
	*entities.Lootpack
	Successful bool
}

func (se LootPickupCompletedEvent) Type() string { return LOOT_PICKUP_COMPLETED_EVENT }

func (se LootPickupCompletedEvent) AsLogMessage() string {
	s := "Aborted"
	if se.Successful {
		s = "Successful"
	}
	return "Loot pickup completed " + s
}
