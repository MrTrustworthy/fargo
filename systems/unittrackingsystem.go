package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"errors"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitTrackingSystem struct {
	Units        []*entities.Unit
	SelectedUnit *entities.Unit
}

func (ss *UnitTrackingSystem) AddUnit(unit *entities.Unit) {
	ss.Units = append(ss.Units, unit)
}

func (ss *UnitTrackingSystem) New(world *ecs.World) {

	engo.Mailbox.Listen(events.INPUT_SELECT_EVENT_NAME, ss.getHandleInputEvent())
	engo.Mailbox.Listen(events.SELECTION_SELECTED_EVENT_NAME, ss.getHandleSelectEvent())
	engo.Mailbox.Listen(events.SELECTION_DESELECTED_EVENT_NAME, ss.getHandleDeselectEvent())

}

func (ss *UnitTrackingSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(events.InputSelectEvent)
		if !ok {
			return
		}

		if ss.SelectedUnit != nil {
			deselectedUnit := ss.SelectedUnit
			ss.SelectedUnit = nil
			engo.Mailbox.Dispatch(events.SelectionDeselectedEvent{Unit: deselectedUnit})
		}

		if unit, err := ss.findUnitUnderMouse(&imsg.MouseTracker); err == nil {
			ss.SelectedUnit = unit
			engo.Mailbox.Dispatch(events.SelectionSelectedEvent{Unit: unit})
		}
	}
}

func (ss *UnitTrackingSystem) getHandleSelectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		smsg, ok := msg.(events.SelectionSelectedEvent)
		if !ok {
			return
		}
		smsg.Unit.AnimationComponent.SelectAnimationByName("jump")
	}
}

func (ss *UnitTrackingSystem) getHandleDeselectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		smsg, ok := msg.(events.SelectionDeselectedEvent)
		if !ok {
			return
		}
		smsg.Unit.AnimationComponent.SelectAnimationByName("duck")
	}
}

func (ss *UnitTrackingSystem) Update(dt float32) {}

func (ss *UnitTrackingSystem) findUnitUnderMouse(tracker *events.MouseTracker) (*entities.Unit, error) {
	for _, unit := range ss.Units {
		xDelta := tracker.MouseX - unit.GetSpaceComponent().Position.X
		yDelta := tracker.MouseY - unit.GetSpaceComponent().Position.Y
		if xDelta > 0 && xDelta < entities.UNITSIZE && yDelta > 0 && yDelta < entities.UNITSIZE {
			return unit, nil
		}
	}
	return nil, errors.New("No unit Found")
}

func (ss *UnitTrackingSystem) RemoveUnit(unit *entities.Unit) {
	index := -1
	for i, val := range ss.Units {
		if val == unit {
			index = i
			break
		}
	}
	if index == -1 {
		panic("Trying to remove non-existing unit!")
	}
	fmt.Println("Length before:", len(ss.Units))
	ss.Units = append(ss.Units[:index], ss.Units[index+1:]...)
	fmt.Println("Length after:", len(ss.Units))
}

func (ss *UnitTrackingSystem) Remove(e ecs.BasicEntity) {}
