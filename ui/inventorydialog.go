package ui

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
)

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