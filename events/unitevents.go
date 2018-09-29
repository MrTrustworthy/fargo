package events

import (
	"github.com/MrTrustworthy/fargo/entities"
	"strconv"
)

const (
	UNIT_REQUEST_DAMAGE_EVENT = "RequestUnitDamageEvent"
	UNIT_DEATH_EVENT          = "UnitDeathEvent"
)

type RequestUnitDamageEvent struct {
	*entities.Unit
	Damage int
}

func (se RequestUnitDamageEvent) Type() string { return UNIT_REQUEST_DAMAGE_EVENT }

func (se RequestUnitDamageEvent) AsLogMessage() string {
	return strconv.Itoa(se.Damage) + "for unit " + se.Unit.Name
}

type UnitDeathEvent struct {
	*entities.Unit
}

func (de UnitDeathEvent) Type() string { return UNIT_DEATH_EVENT }

func (de UnitDeathEvent) AsLogMessage() string {
	return "Unit died:" + de.Unit.Name
}
