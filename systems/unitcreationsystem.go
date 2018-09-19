package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/MrTrustworthy/fargo/entities"
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
	unit := entities.NewUnit(&engo.Point{40, 50})
	entity := common.Renderable(unit)
	aentity := common.Animationable(unit)
	AddToRenderSystem(ucs.World, &entity)
	AddToAnimationSystem(ucs.World, &aentity)
}
