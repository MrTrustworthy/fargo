package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/entities"
)

// Blocking systems are systems that must be idle before some actions (movement, attacking, ...) can be started
type BlockingSystem interface {
	IsIdle() bool
}

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

func AddToMouseSystem(world *ecs.World, tracker *MouseTracker) {
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

func GetUnitAbilitySystem(world *ecs.World) *UnitAbilitySystem {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *UnitAbilitySystem:
			return sys
		}
	}
	return nil
}

func GetLootmanagementSystem(world *ecs.World) *LootManagementSystem {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *LootManagementSystem:
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

func FindUnitUnderMouse(world *ecs.World, point engo.Point) (*entities.Unit, error) {
	return GetUnitTrackingSystem(world).findUnitUnderMouse(point)
}

func FindLootUnderMouse(world *ecs.World, point engo.Point) (*entities.Lootpack, error) {
	return GetLootmanagementSystem(world).FindLootUnderMouse(point)
}

func WorldIsCurrentlyBusy(world *ecs.World) bool {
	for _, system := range world.Systems() {
		if sys, ok := system.(BlockingSystem); ok && !sys.IsIdle(){
			return true
		}
	}
	return false
}
