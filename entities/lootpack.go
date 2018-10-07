package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Lootpack struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	*Inventory
}

const LOOTPACKSIZE = 32

func NewLootpack(point *engo.Point) *Lootpack {

	sprite, err := common.LoadedSprite("models/backpack.png")
	if err != nil {
		panic("Can't load backpack texture")
	}

	lootpack := &Lootpack{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Width:    LOOTPACKSIZE,
			Height:   LOOTPACKSIZE,
		},
		RenderComponent: common.RenderComponent{
			Drawable: sprite,
			Scale:    engo.Point{1, 1},
		},
		Inventory: NewSampleInventory(),
	}
	lootpack.SetCenter(*point)

	return lootpack
}
