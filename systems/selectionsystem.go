package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"errors"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
)

type SelectionSystem struct {
	Units        []*entities.Unit
	SelectedUnit *entities.Unit
}

func (ss *SelectionSystem) Add(unit *entities.Unit) {
	ss.Units = append(ss.Units, unit)
}

func (ss *SelectionSystem) New(world *ecs.World) {

	engo.Mailbox.Listen(INPUT_EVENT_NAME, ss.getHandleSelectEvent())
}

func (ss *SelectionSystem) getHandleSelectEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(InputEvent)
		if !ok || imsg.Action != "Select" {
			return
		}
		if unit, err := ss.findUnitUnderMouse(&imsg.MouseTracker); err == nil {
			ss.SelectedUnit = unit
			fmt.Println("Selected unit with name", unit.Name)
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
