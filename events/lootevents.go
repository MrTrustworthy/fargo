package events

import (
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	LOOT_REQUEST_SPAWN_EVENT  = "RequestUnitDamageEvent"
	LOOT_REQUEST_PICKUP_EVENT = "RequestLootPickupEvent"
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
