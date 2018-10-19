package ui

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image"
	"image/color"
)

type Clicker interface {
	HandleClick()
}

// DIALOG
type Dialog struct {
	Background *DialogBackground
	Elements   []common.Renderable
}

func NewDialog(aabb engo.AABB) *Dialog {
	dialog := &Dialog{}
	bg := NewDialogBackground(aabb)
	dialog.SetBackground(bg)
	return dialog
}

func (d *Dialog) AddElement(elem common.Renderable) {
	d.Elements = append(d.Elements, elem)
}

func (d *Dialog) SetBackground(bg *DialogBackground) {
	d.Elements = append(d.Elements, bg)
	d.Background = bg
}

// BUTTON
type Button struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	clickCallback func(*Button)
}

func NewButton(position engo.AABB, text string, callback func(*Button)) *Button {

	height, width := position.Max.X-position.Min.X, position.Max.Y-position.Min.Y

	button := Button{BasicEntity: ecs.NewBasic()}
	button.SpaceComponent = common.SpaceComponent{
		Position: position.Min,
		Width:    height,
		Height:   width,
	}

	button.RenderComponent = common.RenderComponent{
		Drawable: common.Text{
			Font: getFont(color.Gray16{0x00ff00}),
			Text: text,
		},
	}
	button.RenderComponent.SetShader(common.HUDShader)
	button.RenderComponent.SetZIndex(201)
	button.clickCallback = callback
	return &button
}

func (ht *Button) HandleClick() {
	ht.clickCallback(ht)
}

func (ht *Button) SetText(text string) {
	textElem := ht.RenderComponent.Drawable.(common.Text)
	textElem.Text = text
	ht.RenderComponent.Drawable = textElem
}

// BACKGROUND
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

// UTILS

func getFont(bg color.Color) *common.Font {

	fnt := &common.Font{
		URL:  "fonts/Roboto-Regular.ttf",
		FG:   color.Black,
		Size: 16,
		BG: bg,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}
	return fnt
}
