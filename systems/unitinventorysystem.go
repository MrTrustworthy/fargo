package systems

import (
	"engo.io/ecs"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/ui"
)

type UnitInventorySystem struct {
	*ecs.World
}

func (is *UnitInventorySystem) New(world *ecs.World) {
	is.World = world
	events.Mailbox.Listen(events.INVENTORY_SHOW_EVENT, is.getHandleShowInventory())

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



func (is *UnitInventorySystem) Update(dt float32)        {}
func (is *UnitInventorySystem) Remove(e ecs.BasicEntity) {}
