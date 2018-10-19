package ui

import (
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"strconv"
)

func NewInventoryDialog(inventory *entities.Inventory) *Dialog {
	dialogPosition := engo.AABB{Min: engo.Point{X: 100, Y: 100}, Max: engo.Point{X: 400, Y: 400}}
	d := NewDialog(dialogPosition)

	offset := 0
	for item, amount := range *inventory {

		btnPosition := engo.AABB{
			Min: engo.Point{X: 120, Y: float32(120+offset)},
			Max: engo.Point{X: 380, Y: float32(220+offset)},
		}
		btn := NewButton(btnPosition, item.Name + ": " + strconv.Itoa(amount), func(btn *Button) { fmt.Println("Clicked on " + item.Name) })
		d.AddElement(btn)
		offset += 20
	}


	return d
}
