package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/athom/namepicker"
)

type Unit struct {
	Name string
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	common.AnimationComponent
}

const UNITSIZE = 64

func NewUnit(point *engo.Point) *Unit {

	spriteSheet := common.NewSpritesheetFromFile("models/sheet_hero_idle.png", UNITSIZE, UNITSIZE)
	idleAnimation := &common.Animation{Name: "idle", Frames: []int{0, 1, 2, 3, 4, 5, 6, 7}}

	animationComponent := common.NewAnimationComponent(spriteSheet.Drawables(), 0.1)
	animationComponent.AddDefaultAnimation(idleAnimation)
	return &Unit{
		Name:        namepicker.RandomName(),
		BasicEntity: ecs.NewBasic(),
		SpaceComponent: common.SpaceComponent{
			Position: *point,
			Width:    UNITSIZE,
			Height:   UNITSIZE,
		},
		RenderComponent: common.RenderComponent{
			Drawable: spriteSheet.Cell(0),
			Scale:    engo.Point{1, 1},
		},
		AnimationComponent: animationComponent,
	}
}
