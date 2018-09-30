package scenes

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/systems"
)

type WorldScene struct {
	EventChannel chan<- events.BaseEvent
}

func (*WorldScene) Type() string { return "Fargo" }

func (*WorldScene) Preload() {
	engo.Files.Load("models/hero_sprite.png", "models/backpack.png", "fonts/Roboto-Regular.ttf")
}

func (scene *WorldScene) Setup(updater engo.Updater) {
	world, _ := updater.(*ecs.World)
	scene.LoadSystems(world)
}

func (scene *WorldScene) LoadSystems(world *ecs.World) {
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
	world.AddSystem(&systems.UnitAbilitySystem{})
	world.AddSystem(&systems.DamageSystem{})
	world.AddSystem(&systems.UnitDeathSystem{})
	world.AddSystem(&systems.LootSpawnSystem{})

	engo.Input.RegisterButton(systems.INPUT_CREATE_UNIT_KEY_BIND, engo.KeyC)

	events.InitEventCapturing(scene.EventChannel)

}
