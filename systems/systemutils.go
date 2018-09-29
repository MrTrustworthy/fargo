package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

func AddToRenderSystem(world *ecs.World, renderable common.Renderable) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(renderable.GetBasicEntity(), renderable.GetRenderComponent(), renderable.GetSpaceComponent())
		}
	}
}

func RemoveFromRenderSystem(world *ecs.World, renderable common.Renderable) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Remove(*renderable.GetBasicEntity())
		}
	}
}

func AddToAnimationSystem(world *ecs.World, anim common.Animationable) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.AnimationSystem:
			sys.Add(anim.GetBasicEntity(), anim.GetAnimationComponent(), anim.GetRenderComponent())
		}
	}
}

func RemoveFromAnimationSystem(world *ecs.World, anim common.Animationable) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.AnimationSystem:
			sys.Remove(*anim.GetBasicEntity())
		}
	}
}

func AddToMouseSystem(world *ecs.World, tracker *events.MouseTracker) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(&tracker.BasicEntity, &tracker.MouseComponent, nil, nil)
		}
	}
}

func AddToSelectionSystem(world *ecs.World, unit *entities.Unit) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *UnitTrackingSystem:
			sys.AddUnit(unit)
		}
	}
}

func RemoveFromSelectionSystem(world *ecs.World, unit *entities.Unit) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *UnitTrackingSystem:
			sys.RemoveUnit(unit)
		}
	}
}

// Those functions are used to get global unit state like all units, selected units, units under mouse etc
func GetUnitTrackingSystem(world *ecs.World) *UnitTrackingSystem {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *UnitTrackingSystem:
			return sys
		}
	}
	return nil
}

func GetUnitMovementSystem(world *ecs.World) *UnitMovementSystem {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *UnitMovementSystem:
			return sys
		}
	}
	return nil
}

func GetAllExistingUnits(world *ecs.World) []*entities.Unit {
	return GetUnitTrackingSystem(world).Units
}

func GetCurrentlySelectedUnit(world *ecs.World) *entities.Unit {
	return GetUnitTrackingSystem(world).SelectedUnit
}

func FindUnitUnderMouse(world *ecs.World, tracker *events.MouseTracker) (*entities.Unit, error) {
	return GetUnitTrackingSystem(world).findUnitUnderMouse(tracker)
}

func MovementIsCurrentlyProcessing(world *ecs.World) bool {
	return !GetUnitMovementSystem(world).IsIdle()
}
