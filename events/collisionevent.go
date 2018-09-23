package events

import (
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	COLLISON_EVENT_NAME            = "CollisionEvent"
	COLLISON_EVENT_ACTION_COLLIDED = "Collided"
)

type CollisionEvent struct {
	ActiveUnit  *entities.Unit
	PassiveUnit *entities.Unit
	Action      string
}

func (ce CollisionEvent) Type() string { return COLLISON_EVENT_NAME }

func (ce CollisionEvent) AsLogMessage() string {
	x, y := PointToXYStrings(ce.ActiveUnit.Position)
	return "Action[" + ce.Action + "] between " + ce.ActiveUnit.Name + " and " + ce.PassiveUnit.Name + " at (" + x + ":" + y + ")"
}
