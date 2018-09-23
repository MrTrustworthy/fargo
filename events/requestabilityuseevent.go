package events

import (
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	REQUESTABILITYUSE_EVENT_NAME                   = "RequestAbilityUseEvent"
	REQUESTABILITYUSE_EVENT_ACTION_REQUEST_ABILITY = "RequestAbilityUseAction"
)

type RequestAbilityUseEvent struct {
	Action  string
	Source  *entities.Unit
	Target  *entities.Unit
	Ability *entities.Ability
}

func (raue RequestAbilityUseEvent) Type() string      { return REQUESTABILITYUSE_EVENT_NAME }
func (raue RequestAbilityUseEvent) GetAction() string { return raue.Action }

func (raue RequestAbilityUseEvent) AsLogMessage() string {
	return "Action[" + raue.Action + "] between " + raue.Source.Name + " and " + raue.Target.Name
}
