package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"errors"
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	SELECT_EVENT_NAME            = "SelectionEvent"
	SELECT_EVENT_ACTION_SELECTED = "Selected"
	SELECT_EVENT_ACTION_DESELECT = "Deselected"
)

type SelectionEvent struct {
	Action string
	*entities.Unit
}

func (se SelectionEvent) Type() string { return SELECT_EVENT_NAME }

type SelectionSystem struct {
	Units        []*entities.Unit
	SelectedUnit *entities.Unit
}

func (ss *SelectionSystem) Add(unit *entities.Unit) {
	ss.Units = append(ss.Units, unit)
}

func (ss *SelectionSystem) New(world *ecs.World) {

	engo.Mailbox.Listen(INPUT_EVENT_NAME, ss.getHandleInputEvent())
	engo.Mailbox.Listen(SELECT_EVENT_NAME, ss.getHandleSelectEvent())

}

func (ss *SelectionSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(InputEvent)
		if !ok || imsg.Action != INPUT_EVENT_ACTION_SELECT {
			return
		}

		if ss.SelectedUnit != nil {
			deselectedUnit := ss.SelectedUnit
			ss.SelectedUnit = nil
			engo.Mailbox.Dispatch(SelectionEvent{Action: SELECT_EVENT_ACTION_DESELECT, Unit: deselectedUnit})
		}

		if unit, err := ss.findUnitUnderMouse(&imsg.MouseTracker); err == nil {
			ss.SelectedUnit = unit
			engo.Mailbox.Dispatch(SelectionEvent{Action: SELECT_EVENT_ACTION_SELECTED, Unit: unit})
		}
	}
}

func (ss *SelectionSystem) getHandleSelectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		smsg, ok := msg.(SelectionEvent)
		if !ok {
			return
		}
		if smsg.Action == SELECT_EVENT_ACTION_SELECTED {
			smsg.Unit.AnimationComponent.SelectAnimationByName("jump")
		} else if smsg.Action == SELECT_EVENT_ACTION_DESELECT {
			smsg.Unit.AnimationComponent.SelectAnimationByName("duck")
		}
	}
}

func (ss *SelectionSystem) Update(dt float32) {}

func (ss *SelectionSystem) findUnitUnderMouse(tracker *MouseTracker) (*entities.Unit, error) {
	for _, unit := range ss.Units {
		xDelta := tracker.MouseX - unit.GetSpaceComponent().Position.X
		yDelta := tracker.MouseY - unit.GetSpaceComponent().Position.Y
		if xDelta > 0 && xDelta < entities.UNITSIZE && yDelta > 0 && yDelta < entities.UNITSIZE {
			return unit, nil
		}
	}
	return nil, errors.New("No unit Found")
}

func (ss *SelectionSystem) Remove(e ecs.BasicEntity) {}
