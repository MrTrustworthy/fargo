package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"errors"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
)

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type SelectionSystem struct {
	*ecs.World
	MouseTracker
	Units        []*entities.Unit
	SelectedUnit *entities.Unit
}

func (ss *SelectionSystem) New(world *ecs.World) {
	ss.MouseTracker = MouseTracker{
		BasicEntity:    ecs.NewBasic(),
		MouseComponent: common.MouseComponent{Track: true},
	}
	AddToMouseSystem(world, &ss.MouseTracker)
}

func (ss *SelectionSystem) Add(unit *entities.Unit) {
	ss.Units = append(ss.Units, unit)
}

func (ss *SelectionSystem) Update(dt float32) {

	if engo.Input.Mouse.Action != engo.Press {
		return
	}
	if engo.Input.Mouse.Button == engo.MouseButtonLeft {
		if unit, err := ss.findUnitUnderMouse(); err == nil {
			fmt.Println("Selecting unit", unit.Name)
			ss.SelectedUnit = unit
		}

	} else if engo.Input.Mouse.Button == engo.MouseButtonRight {
		if ss.SelectedUnit != nil {
			ss.SelectedUnit.AnimationComponent.SelectAnimationByName("stab")
		}
	}
}

func (ss *SelectionSystem) findUnitUnderMouse() (*entities.Unit, error) {
	for _, unit := range ss.Units {
		xDelta := ss.MouseTracker.MouseX - unit.GetSpaceComponent().Position.X
		yDelta := ss.MouseTracker.MouseY - unit.GetSpaceComponent().Position.Y
		if xDelta > 0 && xDelta < entities.UNITSIZE && yDelta > 0 && yDelta < entities.UNITSIZE {
			return unit, nil
		}
	}
	return nil, errors.New("No unit Found")
}

func (ss *SelectionSystem) Remove(e ecs.BasicEntity) {}
