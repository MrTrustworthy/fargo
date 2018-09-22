package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
)

const (
	MOVEMENT_EVENT_ACTION_FINISHED = "MoveCompleted"
	MOVEMENT_EVENT_NAME            = "InteractionEvent"
)

type MovementEvent struct {
	*entities.Unit
	Action string
}

func (ae MovementEvent) Type() string { return MOVEMENT_EVENT_NAME }

type UnitMovementSystem struct {
	world         *ecs.World
	CurrentUnit   *entities.Unit
	CurrentTarget *engo.Point
	CurrentPath   []engo.Point
}

func (ums *UnitMovementSystem) New(world *ecs.World) {
	ums.world = world
	engo.Mailbox.Listen(INTERACTION_EVENT_NAME, ums.getHandleActionEvent())
}

func (ums *UnitMovementSystem) getHandleActionEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		amsg, ok := msg.(InteractionEvent)
		if !ok || amsg.Action != INTERACTION_EVENT_ACTION_MOVE {
			return
		}
		if ums.CurrentTarget != nil || ums.CurrentUnit != nil {
			fmt.Println("Can't start new movement until old one is finished")
			return
		}

		if unit := GetCurrentlySelectedUnit(ums.world); unit != nil {
			ums.InitiateMovement(unit, &amsg.Target)
		}
	}
}

func (ums *UnitMovementSystem) InitiateMovement(unit *entities.Unit, point *engo.Point) {
	ums.CurrentUnit = unit
	ums.CurrentTarget = &engo.Point{
		X: point.X + entities.UNIT_CENTER_OFFSET.X,
		Y: point.Y + entities.UNIT_CENTER_OFFSET.Y,
	}
	ums.CurrentPath = InterpolateBetween(
		&ums.CurrentUnit.GetSpaceComponent().Position,
		ums.CurrentTarget,
		ums.CurrentUnit.Speed,
	)
	ums.CurrentUnit.AnimationComponent.SelectAnimationByName("walk")

}

func (ums *UnitMovementSystem) Update(dt float32) {
	if ums.IsIdle() {
		return
	}
	nextPos := ums.CurrentPath[0]
	ums.CurrentPath = ums.CurrentPath[1:]
	ums.CurrentUnit.Position = nextPos

	// stop here if movement is not yet finished
	if len(ums.CurrentPath) > 0 {
		return
	}

	// end the movement
	oldUnit := ums.CurrentUnit
	ums.SetIdle()
	oldUnit.AnimationComponent.SelectAnimationByName("idle")
	engo.Mailbox.Dispatch(MovementEvent{oldUnit, MOVEMENT_EVENT_ACTION_FINISHED})
}

func (ums *UnitMovementSystem) IsIdle() bool {
	return ums.CurrentTarget == nil && ums.CurrentUnit == nil && ums.CurrentPath == nil
}

func (ums *UnitMovementSystem) SetIdle() {
	ums.CurrentUnit = nil
	ums.CurrentTarget = nil
	ums.CurrentPath = nil
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
