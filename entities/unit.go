package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/athom/namepicker"
	"log"
)

type Unit struct {
	Name string
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
}

func NewUnit(point *engo.Point) *Unit {

	texture, err := common.LoadedSprite("models/sheet_hero_idle.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	return &Unit{
		Name:        namepicker.RandomName(),
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: *point,
		},
		RenderComponent: common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{1, 1},
		},
	}
}
