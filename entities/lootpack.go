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
}

const LOOTPACKSIZE = 64

func NewLootpack(point *engo.Point) *Lootpack {

	sprite, err := common.LoadedSprite("models/backpack.png")
	if err != nil {
		panic("Can't load backpack texture")
	}

	lootpack := &Lootpack{
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: *point,
			Width:    LOOTPACKSIZE,
			Height:   LOOTPACKSIZE,
		},
		RenderComponent: common.RenderComponent{
			Drawable: sprite,
			Scale:    engo.Point{1, 1},
		},
	}

	return lootpack
}
