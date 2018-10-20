package systems

import (
	"engo.io/ecs"
	"fmt"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/ui"
)

type UnitInventorySystem struct {
	*ecs.World
}

func (is *UnitInventorySystem) New(world *ecs.World) {
	is.World = world
	events.Mailbox.Listen(events.INVENTORY_SHOW_EVENT, is.getHandleShowInventory())
	events.Mailbox.Listen(events.INVENTORY_ELEMENT_CLICKED, is.getHandleInventoryItemClicked())

}

func (is *UnitInventorySystem) getHandleShowInventory() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		simsg, ok := msg.(events.ShowInventoryEvent)
		if !ok || simsg.Unit == nil {
			return
		}
		inventoryDialog := ui.NewInventoryDialog(simsg.Unit.Inventory)
		events.Mailbox.Dispatch(events.DialogShowEvent{Dialog: inventoryDialog})
	}
}
func (is *UnitInventorySystem) getHandleInventoryItemClicked() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		simsg, ok := msg.(events.InventoryElementClickedEvent)
		if !ok || simsg.Unit == nil {
			return
		}
		fmt.Println("clicked item " + simsg.Item.Name)
	}
}



func (is *UnitInventorySystem) Update(dt float32)        {}
func (is *UnitInventorySystem) Remove(e ecs.BasicEntity) {}
