package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/athom/namepicker"
)

type UnitAttributes struct {
	Name   string
	Speed  float32
	Health int
	AP     int
}

type Unit struct {
	UnitAttributes
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	common.AnimationComponent
	SelectedAbility Ability
}

const UNITSIZE = 64

// describes the offset we need to add to a given unit position to get the "visual" center of the unit
// TODO check if we can use SpaceComponent.SetCenter() instead!
var UNIT_CENTER_OFFSET = engo.Point{
	X: -UNITSIZE / 2,
	Y: -UNITSIZE / 1.5,
}

var idleAnimation = &common.Animation{Name: "idle", Frames: []int{0, 1, 2, 3, 4, 5, 6, 7}}
var jumpAnimation = &common.Animation{Name: "jump", Frames: []int{8, 9, 10, 11, 12}}
var stabAnimation = &common.Animation{Name: "stab", Frames: []int{13, 14, 15, 16, 17}}
var walkAnimation = &common.Animation{Name: "walk", Frames: []int{18, 19, 20}, Loop: true} // needs to loop until manually turned off
var deadAnimation = &common.Animation{Name: "dead", Frames: []int{21, 22, 23, 24, 25, 26, 27, 28}}
var duckAnimation = &common.Animation{Name: "duck", Frames: []int{29, 30, 31, 32, 33, 34}}
var spawnAnimation = &common.Animation{Name: "spawn", Frames: []int{35, 36, 37, 38, 39, 40}}
var upstabAnimation = &common.Animation{Name: "upstab", Frames: []int{41, 42, 43, 44, 45}}
var allAnimations = []*common.Animation{idleAnimation, jumpAnimation, stabAnimation, walkAnimation, deadAnimation, duckAnimation, spawnAnimation, upstabAnimation}

func NewUnit(point *engo.Point) *Unit {

	spriteSheet := common.NewSpritesheetFromFile("models/hero_sprite.png", UNITSIZE, UNITSIZE)

	animationComponent := common.NewAnimationComponent(spriteSheet.Drawables(), 0.1)
	animationComponent.AddAnimations(allAnimations)
	animationComponent.AddDefaultAnimation(idleAnimation)
	animationComponent.SelectAnimationByName("spawn")

	unit := &Unit{
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
		UnitAttributes: UnitAttributes{
			Name:   namepicker.RandomName(),
			Speed:  5.0,
			Health: 10,
			AP:     7,
		},
	}

	unit.SelectedAbility = NewStabAbility(unit)
	return unit
}
