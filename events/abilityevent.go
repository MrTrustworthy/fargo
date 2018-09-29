package events

import (
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	ABILITY_REQUESTUSE_EVENT_NAME = "RequestAbilityUseEvent"
	ABILITY_COMPLETED_EVENT_NAME  = "AbilityCompletedEvent"
	ABILITY_ABORT_EVENT_NAME      = "AbilityAbortedEvent"
)

type RequestAbilityUseEvent struct {
	Ability *entities.Ability
}

func (raue RequestAbilityUseEvent) Type() string { return ABILITY_REQUESTUSE_EVENT_NAME }

func (raue RequestAbilityUseEvent) AsLogMessage() string {
	return (*raue.Ability).Name() + "between " + (*raue.Ability).Source().Name + " and " + (*raue.Ability).Target().Name
}

type AbilityCompletedEvent struct {
	Ability *entities.Ability
}

func (raue AbilityCompletedEvent) Type() string { return ABILITY_COMPLETED_EVENT_NAME }

func (raue AbilityCompletedEvent) AsLogMessage() string {
	return (*raue.Ability).Name() + "between " + (*raue.Ability).Source().Name + " and " + (*raue.Ability).Target().Name
}

type AbilityAbortedEvent struct {
	Ability *entities.Ability
}

func (raue AbilityAbortedEvent) Type() string { return ABILITY_ABORT_EVENT_NAME }

func (raue AbilityAbortedEvent) AsLogMessage() string {
	return (*raue.Ability).Name() + "between " + (*raue.Ability).Source().Name + " and " + (*raue.Ability).Target().Name
}
