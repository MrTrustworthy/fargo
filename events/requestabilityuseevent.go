package events

import (
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	ABILITY_REQUESTUSE_EVENT_NAME = "RequestAbilityUseEvent"
)

type RequestAbilityUseEvent struct {
	Ability *entities.Ability
}

func (raue RequestAbilityUseEvent) Type() string { return ABILITY_REQUESTUSE_EVENT_NAME }

func (raue RequestAbilityUseEvent) AsLogMessage() string {
	return (*raue.Ability).Name() + "between " + (*raue.Ability).Source().Name + " and " + (*raue.Ability).Target().Name
}
