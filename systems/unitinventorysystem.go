package systems

import (
	"engo.io/ecs"
)



type UnitInventorySystem struct {
	*ecs.World
}

func (is *UnitInventorySystem) New(world *ecs.World) {
	is.World = world
}

func (is *UnitInventorySystem) Update(dt float32) {}
func (is *UnitInventorySystem) Remove(e ecs.BasicEntity) {}
