package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"
	"errors"
	"github.com/MrTrustworthy/fargo/entities"
)

func AddToRenderSystem(world *ecs.World, renderable common.Renderable) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(renderable.GetBasicEntity(), renderable.GetRenderComponent(), renderable.GetSpaceComponent())
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
		case *SelectionSystem:
			sys.Add(unit)
		}
	}
}

func GetCurrentlySelectedUnit(world *ecs.World) *entities.Unit {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *SelectionSystem:
			return sys.SelectedUnit
		}
	}
	return nil
}

func FindUnitUnderMouse(world *ecs.World, tracker *MouseTracker) (*entities.Unit, error) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *SelectionSystem:
			return sys.findUnitUnderMouse(tracker)
		}
	}
	return nil, errors.New("No SelectionSystem found")
}
