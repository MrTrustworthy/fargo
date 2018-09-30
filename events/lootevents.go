package events

import (
	"engo.io/engo"
	"fmt"
)

const (
	LOOT_REQUEST_SPAWN_EVENT = "RequestUnitDamageEvent"
)

type RequestLootSpawn struct {
	engo.Point
}

func (se RequestLootSpawn) Type() string { return LOOT_REQUEST_SPAWN_EVENT }

func (se RequestLootSpawn) AsLogMessage() string {
	return fmt.Sprint("at point %.2f : %.2f", se.Point.X, se.Point.Y)
}
