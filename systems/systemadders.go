package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"
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
