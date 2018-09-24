package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type UnitMovementSystem struct {
	world         *ecs.World
	CurrentUnit   *entities.Unit
	CurrentTarget *engo.Point
	CurrentPath   []engo.Point
	LastPosition  engo.Point // used to reset the position after a collision
}

func (ums *UnitMovementSystem) New(world *ecs.World) {
	ums.world = world
	engo.Mailbox.Listen(events.MOVEMENT_REQUESTMOVE_EVENT_NAME, ums.getHandleInteractionEvent())
	engo.Mailbox.Listen(events.COLLISON_EVENT_NAME, ums.getHandleCollisionEvent())

}

func (ums *UnitMovementSystem) getHandleInteractionEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		amsg, ok := msg.(events.MovementRequestEvent)
		if !ok {
			return
		}
		if !ums.IsIdle() {
			fmt.Println("Can't start new movement until old one is finished")
			return
		}

		unit := GetCurrentlySelectedUnit(ums.world)
		if unit == nil {
			panic("This shouldn't happen: No unit is selected, can't perform a movement!")
		}

		ums.InitiateMovement(unit, &amsg.Target, amsg.StopAtDistance)
	}
}

func (ums *UnitMovementSystem) getHandleCollisionEvent() func(msg engo.Message) {
	return func(msg engo.Message) {
		cmsg, ok := msg.(events.CollisionEvent)
		if !ok {
			return
		}
		if ums.IsIdle() {
			panic("Collision causing movement abort, but there is no moving unit!" + cmsg.AsLogMessage())
		}
		ums.ResetToLastPosition()
		ums.StopMovement()
	}
}

func (ums *UnitMovementSystem) InitiateMovement(unit *entities.Unit, point *engo.Point, stopDistance float32) {
	ums.CurrentUnit = unit
	ums.CurrentTarget = &engo.Point{
		X: point.X + entities.UNIT_CENTER_OFFSET.X,
		Y: point.Y + entities.UNIT_CENTER_OFFSET.Y,
	}
	plannedPath := InterpolateBetween(
		&ums.CurrentUnit.GetSpaceComponent().Position,
		ums.CurrentTarget,
		ums.CurrentUnit.Speed,
	)
	ums.CurrentPath = ShortenPathToStopDistance(plannedPath, ums.CurrentTarget, stopDistance)
	ums.CurrentUnit.AnimationComponent.SelectAnimationByName("walk")
}

func (ums *UnitMovementSystem) Update(dt float32) {
	if ums.IsIdle() {
		return
	}
	nextPos := ums.CurrentPath[0]
	ums.CurrentPath = ums.CurrentPath[1:]
	ums.LastPosition = ums.CurrentUnit.Position
	ums.CurrentUnit.Position = nextPos

	// stop here if movement is not yet finished
	if len(ums.CurrentPath) > 0 {
		engo.Mailbox.Dispatch(events.MovementStepEvent{Unit: ums.CurrentUnit})
	} else {
		ums.StopMovement()
	}
}

func (ums *UnitMovementSystem) ResetToLastPosition() {
	fmt.Println("Resetting position from", ums.CurrentUnit.Position, "to", ums.LastPosition)
	ums.CurrentUnit.Position = ums.LastPosition
}

func (ums *UnitMovementSystem) IsIdle() bool {
	return ums.CurrentTarget == nil && ums.CurrentUnit == nil && ums.CurrentPath == nil && ums.LastPosition == engo.Point{}
}

func (ums *UnitMovementSystem) StopMovement() {
	if ums.IsIdle() {
		panic("Attempting to stop movemement, but there is no moving unit!")
	}
	oldUnit := ums.CurrentUnit
	ums.CurrentUnit = nil
	ums.CurrentTarget = nil
	ums.CurrentPath = nil
	ums.LastPosition = engo.Point{}
	oldUnit.AnimationComponent.SelectAnimationByName("idle")
	engo.Mailbox.Dispatch(events.MovementCompletedEvent{Unit: oldUnit})
}

func InterpolateBetween(a, b *engo.Point, stepSize float32) []engo.Point {
	dist := a.PointDistance(*b)

	var points []engo.Point
	// we always start the movement at the origin position itself so that when moving and encountering an issue,
	// we'll always have a valid position (the origin) to go back to
	points = append(points, *a)
	for i := stepSize; i < dist; i += stepSize {
		p := engo.Point{
			X: a.X + (i/dist)*(b.X-a.X),
			Y: a.Y + (i/dist)*(b.Y-a.Y),
		}
		points = append(points, p)
	}
	points = append(points, *b)
	return points
}

// Can be used to shorten a given path so that the last point is at least stopDistance from the target
func ShortenPathToStopDistance(plannedPath []engo.Point, target *engo.Point, stopDistance float32) []engo.Point {

	maxOvershoot := float32(0.01) // in order to balance out floating point comparison
	var points []engo.Point
	for _, pathPoint := range plannedPath {
		pointDistance := pathPoint.PointDistance(*target)
		if pointDistance+maxOvershoot < stopDistance {
			fmt.Println("Filtered out at least one point in path due to stop distance!")
			break
		}
		points = append(points, pathPoint)
	}

	// even a shortened path must always contain at least one point - the origin of the movement itself
	// see InterpolateBetween() for details
	if len(points) == 0 {
		points = []engo.Point{plannedPath[0]}
	}

	return points
}

func (ums *UnitMovementSystem) Remove(e ecs.BasicEntity) {}
