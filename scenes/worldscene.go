package scenes

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
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
	world.AddSystem(&systems.UnitTrackingSystem{})
	world.AddSystem(&systems.InputSystem{})
	world.AddSystem(&systems.UserInterfaceSystem{})
	world.AddSystem(&systems.UnitMovementSystem{})
	world.AddSystem(&systems.UnitInteractionSystem{})
	world.AddSystem(&systems.UnitCollisionSystem{})

	engo.Input.RegisterButton(events.INPUT_EVENT_ACTION_CREATEUNIT, engo.KeyC)

	events.InitEventLogging(fmt.Println)

}
