package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
)

type UnitMovementSystem struct {
	world         *ecs.World
	CurrentUnit   *entities.Unit
	CurrentTarget *engo.Point
	CurrentPath   []engo.Point
}

func (ums *UnitMovementSystem) New(world *ecs.World) {
	ums.world = world
	engo.Mailbox.Listen(INPUT_EVENT_NAME, ums.getHandleInputEvent())
}

func (ums *UnitMovementSystem) getHandleInputEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		imsg, ok := msg.(InputEvent)
		if !ok || imsg.Action != "Move" {
			return
		}
		if ums.CurrentTarget != nil || ums.CurrentUnit != nil {
			fmt.Println("Can't start new movement until old one is finished")
			return
		}

		if unit := GetCurrentlySelectedUnit(ums.world); unit != nil {
			ums.CurrentUnit = unit
			ums.CurrentTarget = &engo.Point{
				X: imsg.MouseTracker.MouseX + entities.UNIT_CENTER_OFFSET.X,
				Y: imsg.MouseTracker.MouseY + entities.UNIT_CENTER_OFFSET.Y,
			}
			ums.CurrentPath = InterpolateBetween(
				&ums.CurrentUnit.GetSpaceComponent().Position,
				ums.CurrentTarget,
				ums.CurrentUnit.Speed,
			)
			ums.CurrentUnit.AnimationComponent.SelectAnimationByName("walk")
		}

	}
}

func (ums *UnitMovementSystem) Update(dt float32) {
	if ums.CurrentTarget == nil || ums.CurrentUnit == nil || ums.CurrentPath == nil {
		return
	}
	nextPos := ums.CurrentPath[0]
	ums.CurrentPath = ums.CurrentPath[1:]
	ums.CurrentUnit.Position = nextPos

	if len(ums.CurrentPath) == 0 {
		ums.CurrentUnit.AnimationComponent.SelectAnimationByName("idle")
		ums.CurrentUnit = nil
		ums.CurrentTarget = nil
		ums.CurrentPath = nil
	}

}

func InterpolateBetween(a, b *engo.Point, stepSize float32) []engo.Point {
	dist := a.PointDistance(*b)

	var points []engo.Point

	for i := float32(0.0); i < dist; i += stepSize {
		p := engo.Point{
			X: a.X + (i/dist)*(b.X-a.X),
			Y: a.Y + (i/dist)*(b.Y-a.Y),
		}
		points = append(points, p)
	}
	points = append(points, *b)

	return points
}

func (ums *UnitMovementSystem) Remove(e ecs.BasicEntity) {}
