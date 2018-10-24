package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
	"github.com/MrTrustworthy/fargo/ui"
	"strconv"
)

type UnitInventorySystem struct {
	*ecs.World
	currentDamageEvent *events.DialogShowEvent
}

func (is *UnitInventorySystem) New(world *ecs.World) {
	is.World = world
	eventsystem.Mailbox.Listen(events.INVENTORY_SHOW_EVENT, is.getHandleShowInventory())
	eventsystem.Mailbox.Listen(events.INVENTORY_ELEMENT_CLICKED, is.getHandleInventoryItemClicked())

}

func (is *UnitInventorySystem) Update(dt float32) {
	if is.currentDamageEvent != nil {
		is.handleShowInventory(is.currentDamageEvent)
		is.currentDamageEvent = nil
	}
}

func (is *UnitInventorySystem) getHandleShowInventory() func(msg engo.Message) {
	return func(msg engo.Message) {
		simsg, ok := msg.(events.ShowInventoryEvent)
		if !ok || simsg.Unit == nil {
			return
		}

		if is.currentDamageEvent != nil {
			fmt.Println("WARNING: Trying to add ShowInventoryEvent even though there is already one pending")
			return
		}
		inventoryDialog := CreateInventoryDialog(simsg.Unit)
		is.currentDamageEvent = &events.DialogShowEvent{Dialog: inventoryDialog}
	}
}

func (is *UnitInventorySystem) handleShowInventory(msg *events.DialogShowEvent) {
	engo.Mailbox.Dispatch(msg)
}

func (is *UnitInventorySystem) getHandleInventoryItemClicked() func(msg engo.Message) {
	return func(msg engo.Message) {
		simsg, ok := msg.(events.InventoryElementClickedEvent)
		if !ok || simsg.Unit == nil {
			return
		}
		fmt.Println("clicked item " + simsg.Item.Name)
	}
}

func CreateInventoryDialog(unit *entities.Unit) *ui.Dialog {
	dialogPosition := engo.AABB{Min: engo.Point{X: 100, Y: 100}, Max: engo.Point{X: 400, Y: 400}}
	d := ui.NewDialog(dialogPosition)

	offset := 0
	for item, amount := range *unit.Inventory {

		btnPosition := engo.AABB{
			Min: engo.Point{X: 120, Y: float32(120 + offset)},
			Max: engo.Point{X: 380, Y: float32(220 + offset)},
		}
		event := events.InventoryElementClickedEvent{
			Unit: unit,
			Item: &item,
		}
		btn := ui.NewButton(btnPosition, item.Name+": "+strconv.Itoa(amount), event)
		d.AddElement(btn)
		offset += 20
	}

	return d
}

func (is *UnitInventorySystem) Remove(e ecs.BasicEntity) {}
