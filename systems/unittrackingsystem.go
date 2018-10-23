package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"errors"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
)

type UnitTrackingSystem struct {
	Units        []*entities.Unit
	SelectedUnit *entities.Unit
	*ecs.World
}

func (ss *UnitTrackingSystem) AddUnit(unit *entities.Unit) {
	ss.Units = append(ss.Units, unit)
}

func (ss *UnitTrackingSystem) New(world *ecs.World) {
	ss.World = world
	eventsystem.Mailbox.Listen(events.INPUT_SELECT_EVENT_NAME, ss.getHandleInputEvent())
	eventsystem.Mailbox.Listen(events.SELECTION_SELECTED_EVENT_NAME, ss.getHandleSelectEvent())
	eventsystem.Mailbox.Listen(events.SELECTION_DESELECTED_EVENT_NAME, ss.getHandleDeselectEvent())
	eventsystem.Mailbox.Listen(events.UNIT_REMOVAL_EVENT, ss.getHandleRemoveUnitEvent())

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
			eventsystem.Mailbox.Dispatch(events.SelectionDeselectedEvent{Unit: deselectedUnit})
		}

		if unit, err := ss.findUnitUnderMouse(imsg.Point); err == nil {
			ss.SelectedUnit = unit
			eventsystem.Mailbox.Dispatch(events.SelectionSelectedEvent{Unit: unit})
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

func (ss *UnitTrackingSystem) getHandleRemoveUnitEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		smsg, ok := msg.(events.UnitRemovalEvent)
		if !ok {
			return
		}
		ss.removeUnit(smsg.Unit)
	}
}

func (ss *UnitTrackingSystem) Update(dt float32) {}

func (ss *UnitTrackingSystem) findUnitUnderMouse(point engo.Point) (*entities.Unit, error) {
	for _, unit := range ss.Units {
		xDelta := point.X - unit.GetSpaceComponent().Position.X
		yDelta := point.Y - unit.GetSpaceComponent().Position.Y
		if xDelta > 0 && xDelta < entities.UNITSIZE && yDelta > 0 && yDelta < entities.UNITSIZE {
			return unit, nil
		}
	}
	return nil, errors.New("No unit Found")
}

func (ss *UnitTrackingSystem) removeUnit(unit *entities.Unit) {
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
	ss.Units = append(ss.Units[:index], ss.Units[index+1:]...)
	RemoveFromRenderSystem(ss.World, unit)
	RemoveFromAnimationSystem(ss.World, unit)
}

func (ss *UnitTrackingSystem) Remove(e ecs.BasicEntity) {}
