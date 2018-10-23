package scenes

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
	"github.com/MrTrustworthy/fargo/systems"
)

type WorldScene struct{}

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
	world.AddSystem(&systems.LootManagementSystem{})
	world.AddSystem(&systems.UnitInventorySystem{})
	world.AddSystem(&systems.DialogSystem{})
	world.AddSystem(&systems.TickSystem{})

	world.AddSystem(&systems.SimulationTestSystem{})

	engo.Input.RegisterButton(systems.INPUT_CREATE_UNIT_KEY_BIND, engo.KeyC)
	engo.Input.RegisterButton(systems.INPUT_RUN_TESTS_KEY_BIND, engo.KeyT)
	engo.Input.RegisterButton(systems.INPUT_SHOW_INVENTORY_DIALOG, engo.KeyF)
	engo.Input.RegisterButton(systems.INPUT_HIDE_DIALOG, engo.KeyG)

}


func logMsg(eventChan chan eventsystem.BaseEvent) {
	for true {
		msg := <-eventChan
		if msg.Type() != events.MOVEMENT_STEP_EVENT_NAME && msg.Type() != events.TICK_EVENT {
			fmt.Println(msg.Type(), msg.AsLogMessage())
		}
	}
}
