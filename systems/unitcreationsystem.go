package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
)

type UnitCreationSystem struct {
	*ecs.World
}

func (ucs *UnitCreationSystem) New(world *ecs.World) {
	ucs.World = world
}

func (ucs *UnitCreationSystem) Update(dt float32) {

	if engo.Input.Button("AddCity").JustPressed() {
		ucs.AddUnit()
	}
}

func (ucs *UnitCreationSystem) Remove(e ecs.BasicEntity) {}

func (ucs *UnitCreationSystem) AddUnit() {
	fmt.Println("yyyeeesccc")
	unit := entities.NewUnit(&engo.Point{40, 50})
	entity := common.Renderable(unit)
	AddToRenderSystem(ucs.World, &entity)
}
