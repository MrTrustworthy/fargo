package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/athom/namepicker"
	"math/rand"
)

type UnitCreationSystem struct {
	*ecs.World
}

func (ucs *UnitCreationSystem) New(world *ecs.World) {
	ucs.World = world
}

func (ucs *UnitCreationSystem) Update(dt float32) {

	if engo.Input.Button("CreateUnit").JustPressed() {
		ucs.AddUnit()
	}
}

func (ucs *UnitCreationSystem) Remove(e ecs.BasicEntity) {}

func (ucs *UnitCreationSystem) AddUnit() {
	unit := entities.NewUnit(&engo.Point{rand.Float32() * 500, rand.Float32() * 500})
	fmt.Println("name of the new unit is", namepicker.RandomName())
	AddToRenderSystem(ucs.World, unit)
	AddToAnimationSystem(ucs.World, unit)
}
