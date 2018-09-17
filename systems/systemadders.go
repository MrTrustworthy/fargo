package systems

import (
	"engo.io/ecs"
	"engo.io/engo/common"
)

func AddToRenderSystem(world *ecs.World, renderable *common.Renderable) {
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add((*renderable).GetBasicEntity(), (*renderable).GetRenderComponent(), (*renderable).GetSpaceComponent())
		}
	}
}
