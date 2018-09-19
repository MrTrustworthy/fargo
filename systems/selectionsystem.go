package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
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
	entities.Unit
}

func (ss *SelectionSystem) New(world *ecs.World) {
	ss.MouseTracker = MouseTracker{
		BasicEntity:    ecs.NewBasic(),
		MouseComponent: common.MouseComponent{Track: true},
	}
	AddToMouseSystem(world, &ss.MouseTracker)
}

func (ss *SelectionSystem) Update(dt float32) {

	if engo.Input.Mouse.Action != engo.Press {
		return
	}
	if engo.Input.Mouse.Button == engo.MouseButtonLeft {
		fmt.Println("yoooo at", engo.Point{ss.MouseTracker.MouseX, ss.MouseTracker.MouseY})
	} else if engo.Input.Mouse.Button == engo.MouseButtonRight {
		fmt.Println("neeeey at", ss.MouseTracker.MouseX, ss.MouseTracker.MouseY)
	}
}

func (ss *SelectionSystem) Remove(e ecs.BasicEntity) {}
