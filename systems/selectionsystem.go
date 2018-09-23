package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"errors"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type SelectionSystem struct {
	Units        []*entities.Unit
	SelectedUnit *entities.Unit
}

func (ss *SelectionSystem) Add(unit *entities.Unit) {
	ss.Units = append(ss.Units, unit)
}

func (ss *SelectionSystem) New(world *ecs.World) {

	engo.Mailbox.Listen(events.INPUT_EVENT_NAME, ss.getHandleInputEvent())
	engo.Mailbox.Listen(events.SELECT_EVENT_NAME, ss.getHandleSelectEvent())

}

func (ss *SelectionSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.InputEvent)
		if !ok || imsg.Action != events.INPUT_EVENT_ACTION_SELECT {
			return
		}

		if ss.SelectedUnit != nil {
			deselectedUnit := ss.SelectedUnit
			ss.SelectedUnit = nil
			engo.Mailbox.Dispatch(events.SelectionEvent{Action: events.SELECT_EVENT_ACTION_DESELECT, Unit: deselectedUnit})
		}

		if unit, err := ss.findUnitUnderMouse(&imsg.MouseTracker); err == nil {
			ss.SelectedUnit = unit
			engo.Mailbox.Dispatch(events.SelectionEvent{Action: events.SELECT_EVENT_ACTION_SELECTED, Unit: unit})
		}
	}
}

func (ss *SelectionSystem) getHandleSelectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		smsg, ok := msg.(events.SelectionEvent)
		if !ok {
			return
		}
		if smsg.Action == events.SELECT_EVENT_ACTION_SELECTED {
			smsg.Unit.AnimationComponent.SelectAnimationByName("jump")
		} else if smsg.Action == events.SELECT_EVENT_ACTION_DESELECT {
			smsg.Unit.AnimationComponent.SelectAnimationByName("duck")
		}
	}
}

func (ss *SelectionSystem) Update(dt float32) {}

func (ss *SelectionSystem) findUnitUnderMouse(tracker *events.MouseTracker) (*entities.Unit, error) {
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
