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

func NewInventoryDialog() *Dialog {
	dialogPosition := engo.AABB{Min: engo.Point{X: 100, Y: 100}, Max: engo.Point{X: 400, Y: 400}}
	btnPosition := engo.AABB{Min: engo.Point{X: 120, Y: 120}, Max: engo.Point{X: 380, Y: 220}}

	bg := NewDialogBackground(dialogPosition)
	btn := NewInventoryButton(btnPosition, "Hello")
	d := Dialog{
		Background: bg,
	}
	d.Elements = append(d.Elements, bg, btn)
	return &d
}

type Clicker interface {
	HandleClick()
}

type Button struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func NewInventoryButton(outline engo.AABB, text string) *Button {
	height, width := outline.Max.X-outline.Min.X, outline.Max.Y-outline.Min.Y

	button := Button{BasicEntity: ecs.NewBasic()}
	button.SpaceComponent = common.SpaceComponent{
		Position: outline.Min,
		Width:    height,
		Height:   width,
	}

	fnt := &common.Font{
		URL:  "fonts/Roboto-Regular.ttf",
		FG:   color.Black,
		Size: 16,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	button.RenderComponent = common.RenderComponent{
		Drawable: common.Text{
			Font: fnt,
			Text: "##",
		},
	}
	button.RenderComponent.SetShader(common.HUDShader)
	button.RenderComponent.SetZIndex(201)
	button.SetText(text)
	return &button
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
