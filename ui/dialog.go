package ui

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"image"
	"image/color"
)

type Dialog struct {
	Background *DialogBackground
	Elements   []common.Renderable
}



type Clicker interface {
	HandleClick()
}

type Button struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (ht *Button) HandleClick() {
	fmt.Println("WHOOOO CLICKED BUTTON")
}

func (ht *Button) SetText(text string) {
	textElem := ht.RenderComponent.Drawable.(common.Text)
	textElem.Text = text
	ht.RenderComponent.Drawable = textElem
}

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

	// This shader is responsible for always keeping the element at the same point in the window
	dBackground.RenderComponent.SetShader(common.HUDShader)
	dBackground.RenderComponent.SetZIndex(200)

	return &dBackground
}
