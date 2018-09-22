package scenes

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/systems"
)

type WorldScene struct{}

func (*WorldScene) Type() string { return "myGame" }

func (*WorldScene) Preload() {
	engo.Files.Load("models/hero_sprite.png", "fonts/Roboto-Regular.ttf")
}

func (scene *WorldScene) Setup(updater engo.Updater) {
	world, _ := updater.(*ecs.World)
	scene.LoadSystems(world)
}

func (*WorldScene) LoadSystems(world *ecs.World) {
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(common.NewKeyboardScroller(400, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	world.AddSystem(&common.EdgeScroller{400, 20})
	world.AddSystem(&common.MouseZoomer{-0.125})
	world.AddSystem(&common.AnimationSystem{})

	world.AddSystem(&systems.UnitCreationSystem{})
	world.AddSystem(&systems.SelectionSystem{})
	world.AddSystem(&systems.InputSystem{})
	world.AddSystem(&systems.UserInterfaceSystem{})
	world.AddSystem(&systems.UnitMovementSystem{})

	engo.Input.RegisterButton("CreateUnit", engo.KeyC)

}
