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
	CurrentTarget engo.Point
	CurrentPath   []engo.Point
	LastPosition  engo.Point // used to reset the position after a collision
}

func (ums *UnitMovementSystem) New(world *ecs.World) {
	ums.world = world
	events.Mailbox.Listen(events.MOVEMENT_REQUESTMOVE_EVENT_NAME, ums.getHandleInteractionEvent())
	events.Mailbox.Listen(events.COLLISON_EVENT_NAME, ums.getHandleCollisionEvent())

}

func (ums *UnitMovementSystem) getHandleInteractionEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
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

func (ums *UnitMovementSystem) getHandleCollisionEvent() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
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
	ums.CurrentTarget = engo.Point{
		X: point.X,
		Y: point.Y,
	}
	plannedPath := InterpolateBetween(
		ums.CurrentUnit.GetSpaceComponent().Center(),
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
	ums.CurrentUnit.SetCenter(nextPos)

	// stop here if movement is not yet finished
	if len(ums.CurrentPath) > 0 {
		events.Mailbox.Dispatch(events.MovementStepEvent{Unit: ums.CurrentUnit})
	} else {
		ums.StopMovement()
	}
}

func (ums *UnitMovementSystem) ResetToLastPosition() {
	fmt.Println("Resetting position from", ums.CurrentUnit.Position, "to", ums.LastPosition)
	ums.CurrentUnit.Position = ums.LastPosition
}

func (ums *UnitMovementSystem) IsIdle() bool {
	return ums.CurrentTarget == engo.Point{} && ums.CurrentUnit == nil && ums.CurrentPath == nil && ums.LastPosition == engo.Point{}
}

func (ums *UnitMovementSystem) StopMovement() {
	if ums.IsIdle() {
		panic("Attempting to stop movemement, but there is no moving unit!")
	}
	oldUnit := ums.CurrentUnit
	ums.CurrentUnit = nil
	ums.CurrentTarget = engo.Point{}
	ums.CurrentPath = nil
	ums.LastPosition = engo.Point{}
	oldUnit.AnimationComponent.SelectAnimationByName("idle")
	events.Mailbox.Dispatch(events.MovementCompletedEvent{Unit: oldUnit})
}

func InterpolateBetween(a, b engo.Point, stepSize float32) []engo.Point {
	dist := a.PointDistance(b)

	var points []engo.Point
	// we always start the movement at the origin position itself so that when moving and encountering an issue,
	// we'll always have a valid position (the origin) to go back to
	points = append(points, a)
	for i := stepSize; i < dist; i += stepSize {
		p := engo.Point{
			X: a.X + (i/dist)*(b.X-a.X),
			Y: a.Y + (i/dist)*(b.Y-a.Y),
		}
		points = append(points, p)
	}
	points = append(points, b)
	return points
}

// Can be used to shorten a given path so that the last point is at least stopDistance from the target
func ShortenPathToStopDistance(plannedPath []engo.Point, target engo.Point, stopDistance float32) []engo.Point {

	maxOvershoot := float32(0.01) // in order to balance out floating point comparison
	var points []engo.Point
	for _, pathPoint := range plannedPath {
		pointDistance := pathPoint.PointDistance(target)
		if pointDistance-maxOvershoot < stopDistance {
			// add the first point that is in range of stopDistance so that dist(a, b) < stopDist after moving
			// this also ensures that there is at least one point in the list, see InterpolateBetween() for details
			points = append(points, pathPoint)
			fmt.Println("Keeping point that is", pointDistance, "away at", pathPoint, "removing later points")
			break
		}
		points = append(points, pathPoint)
	}

	return points
}

func dispatchMoveTo(x, y, dist float32) {
	events.Mailbox.Dispatch(events.MovementRequestEvent{
		Target:         engo.Point{X: x, Y: y},
		StopAtDistance: dist,
	})
}

func (ums *UnitMovementSystem) Remove(e ecs.BasicEntity) {}
