package ui

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image"
	"image/color"
)

type Dialog struct {
	Background *DialogBackground
	Elements []common.Renderable
}

func NewDialog() *Dialog {
	dialogPosition := engo.AABB{Min: engo.Point{X: 100, Y: 100}, Max: engo.Point{X: 400, Y: 400}}
	bg := NewDialogBackground(dialogPosition)
	d := Dialog{
		Background: bg,
	}
	d.Elements = append(d.Elements, bg)
	return &d
}

type Button struct{}

type DialogBackground struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func NewDialogBackground(outline engo.AABB) *DialogBackground {
	height, width := outline.Max.X-outline.Min.X, outline.Max.Y-outline.Min.Y

	dBackground := DialogBackground{BasicEntity: ecs.NewBasic()}
	dBackground.SpaceComponent = common.SpaceComponent{
		Position: outline.Min,
		Width:    height,
		Height:   width,
	}
	dBackgroundImage := image.NewUniform(color.RGBA{205, 205, 205, 255})
	dBackgroundNRGBA := common.ImageToNRGBA(dBackgroundImage, int(width), int(height))
	dBackgroundImageObj := common.NewImageObject(dBackgroundNRGBA)
	dBackgroundTexture := common.NewTextureSingle(dBackgroundImageObj)

	dBackground.RenderComponent = common.RenderComponent{
		Drawable: dBackgroundTexture,
		Scale:    engo.Point{1, 1},
	}
	dBackground.RenderComponent.SetShader(common.HUDShader)
	dBackground.RenderComponent.SetZIndex(200)

	return &dBackground
}
