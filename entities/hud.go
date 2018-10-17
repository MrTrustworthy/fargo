package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image"
	"image/color"
	"strconv"
)

type HUDText struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	*Unit
}

func NewHUDText() *HUDText {
	height, width := float32(90), float32(700)

	hudText := HUDText{BasicEntity: ecs.NewBasic()}
	hudText.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: engo.WindowWidth()/2 - width/2,
			Y: engo.WindowHeight() - height},
		Width:  width,
		Height: height,
	}

	fnt := &common.Font{
		URL:  "fonts/Roboto-Regular.ttf",
		FG:   color.Black,
		Size: 22,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	hudText.RenderComponent = common.RenderComponent{
		Drawable: common.Text{
			Font: fnt,
			Text: "Unit: None",
		},
	}
	hudText.RenderComponent.SetShader(common.HUDShader)
	hudText.RenderComponent.SetZIndex(101)
	return &hudText
}

func (ht *HUDText) SetDisplayeddUnit(unit *Unit) {
	ht.Unit = unit
	ht.UpdateTextForUnitIfDisplayed(unit)
}

func (ht *HUDText) UpdateTextForUnitIfDisplayed(unit *Unit) {
	if unit != ht.Unit {
		// No reason to update hud text for unit that's not currently selected
		return
	}
	unitText := "Unit: " + unit.Name + "\nSelected Attack: " + unit.SelectedAbility.Name() + "\nSpeed:" +
		strconv.Itoa(int(unit.Speed)) + " HP: " + strconv.Itoa(unit.Health) + " AP: " + strconv.Itoa(unit.AP) + " | " +
		unit.Inventory.String()
	ht.SetText(unitText)
}

func (ht *HUDText) SetText(text string) {
	textElem := ht.RenderComponent.Drawable.(common.Text)
	textElem.Text = text
	ht.RenderComponent.Drawable = textElem
}

type HUD struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func NewHUD() *HUD {
	height, width := float32(90), float32(700)

	hud := HUD{BasicEntity: ecs.NewBasic()}
	hud.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: engo.WindowWidth()/2 - width/2,
			Y: engo.WindowHeight() - height},
		Width:  width,
		Height: height,
	}
	hudImage := image.NewUniform(color.RGBA{205, 205, 205, 255})
	hudNRGBA := common.ImageToNRGBA(hudImage, int(width), int(height))
	hudImageObj := common.NewImageObject(hudNRGBA)
	hudTexture := common.NewTextureSingle(hudImageObj)

	hud.RenderComponent = common.RenderComponent{
		Drawable: hudTexture,
		Scale:    engo.Point{1, 1},
	}
	hud.RenderComponent.SetShader(common.HUDShader)
	hud.RenderComponent.SetZIndex(100)

	return &hud
}
