package events

import (
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	ABILITY_REQUESTUSE_EVENT_NAME = "RequestAbilityUseEvent"
)

type RequestAbilityUseEvent struct {
	Source  *entities.Unit
	Target  *entities.Unit
	Ability *entities.Ability
}

func (raue RequestAbilityUseEvent) Type() string { return ABILITY_REQUESTUSE_EVENT_NAME }

func (raue RequestAbilityUseEvent) AsLogMessage() string {
	return "between " + raue.Source.Name + " and " + raue.Target.Name
}
