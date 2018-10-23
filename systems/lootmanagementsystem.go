package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"errors"
	"fmt"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
	"github.com/MrTrustworthy/fargo/eventsystem"
	"github.com/MrTrustworthy/fargo/ui"
	"strconv"
)

const LOOT_PICKUP_DISTANCE = 20.0

type LootManagementSystem struct {
	*ecs.World
	ActiveLootPacks []*entities.Lootpack
}

func (lss *LootManagementSystem) New(world *ecs.World) {
	lss.World = world
	eventsystem.Mailbox.Listen(events.LOOT_REQUEST_SPAWN_EVENT, lss.getHandleRequestLootSpawn())
	eventsystem.Mailbox.Listen(events.LOOT_REQUEST_PICKUP_EVENT, lss.getHandleRequestLootPickup())
	eventsystem.Mailbox.Listen(events.LOOT_REQUEST_ITEM_PICKUP_EVENT, lss.getHandleItemPickup())

}

func (lss *LootManagementSystem) getHandleRequestLootSpawn() func(msg engo.Message) {
	return func(msg engo.Message) {
		udmsg, ok := msg.(events.RequestLootSpawn)
		if !ok {
			return
		}

		lootpack := entities.NewLootpack(&udmsg.Point)
		lss.ActiveLootPacks = append(lss.ActiveLootPacks, lootpack)
		AddToRenderSystem(lss.World, lootpack)
		eventsystem.Mailbox.Dispatch(events.LootHasSpawnedEvent{Lootpack: lootpack})
	}
}

func (lss *LootManagementSystem) getHandleRequestLootPickup() func(msg engo.Message) {
	return func(msg engo.Message) {
		rlpe, ok := msg.(events.RequestLootPickupEvent)
		if !ok {
			return
		}
		if WorldIsCurrentlyBusy(lss.World) {
			// Can't start pickup as long as movement is still ongoing
			// TODO check if this is dead code
			fmt.Println("Can't start looting since movement is still in progress")
			return
		}
		unitPosititon, packPosition := rlpe.Unit.Center(), rlpe.Lootpack.Center()
		currentDistance := unitPosititon.PointDistance(packPosition)
		if currentDistance <= LOOT_PICKUP_DISTANCE {
			lss.ensurePickupDialog(rlpe.Unit, rlpe.Lootpack)
		} else {
			moveCloserAndRetryPickup(&rlpe)
		}
	}
}

func (lss *LootManagementSystem) ensurePickupDialog(unit *entities.Unit, lootpack *entities.Lootpack) {
	dialog := CreateLootboxPickupDialog(unit, lootpack)
	eventsystem.Mailbox.Dispatch(events.DialogShowEvent{Dialog: dialog})
}

func (lss *LootManagementSystem) getHandleItemPickup() func(msg engo.Message) {
	return func(msg engo.Message) {
		rlipe, ok := msg.(events.RequestLootItemPickupEvent)
		if !ok {
			return
		}

		amount, err := rlipe.Lootpack.RemoveAll(*rlipe.Item)
		if err != nil {
			panic("This should not happen!!!")
		}
		rlipe.Unit.Inventory.Add(*rlipe.Item, amount)

		// refresh lootpack dialog
		lss.ensurePickupDialog(rlipe.Unit, rlipe.Lootpack)
		// refresh unit item display
		eventsystem.Mailbox.Dispatch(events.UnitAttributesChangedEvent{Unit: rlipe.Unit})

		eventsystem.Mailbox.Dispatch(events.LootPickupItemCompletedEvent{
			Unit:       rlipe.Unit,
			Lootpack:   rlipe.Lootpack,
			Item:       rlipe.Item,
			Successful: true,
		})

		// clean up lootpack after picking up the last item
		if rlipe.Lootpack.IsEmpty() {
			lss.RemoveLootbox(rlipe.Lootpack)
			eventsystem.Mailbox.Dispatch(events.DialogHideEvent{})
			eventsystem.Mailbox.Dispatch(events.LootPickupCompletedEvent{Lootpack: rlipe.Lootpack, Successful: true})
		}
	}
}

func moveCloserAndRetryPickup(pickupEvent *events.RequestLootPickupEvent) {
	eventsystem.Mailbox.ListenOnce(events.MOVEMENT_COMPLETED_EVENT_NAME, func(msg engo.Message) {
		if cmsg, ok := msg.(events.MovementCompletedEvent); ok && cmsg.Successful {
			eventsystem.Mailbox.Dispatch(*pickupEvent)
		} else {
			eventsystem.Mailbox.Dispatch(events.LootPickupCompletedEvent{Lootpack: pickupEvent.Lootpack, Successful: false})
		}
	})
	eventsystem.Mailbox.Dispatch(events.MovementRequestEvent{
		Target:         pickupEvent.Lootpack.Center(),
		StopAtDistance: LOOT_PICKUP_DISTANCE,
		Unit:           pickupEvent.Unit,
	})
}

func (lss *LootManagementSystem) FindLootUnderMouse(point engo.Point) (*entities.Lootpack, error) {
	for _, pack := range lss.ActiveLootPacks {
		xDelta := point.X - pack.GetSpaceComponent().Position.X
		yDelta := point.Y - pack.GetSpaceComponent().Position.Y
		if xDelta > 0 && xDelta < entities.LOOTPACKSIZE && yDelta > 0 && yDelta < entities.LOOTPACKSIZE {
			return pack, nil
		}
	}
	return nil, errors.New("No lootpack Found")
}

func (lss *LootManagementSystem) RemoveLootbox(lootpack *entities.Lootpack) {
	index := -1
	for i, val := range lss.ActiveLootPacks {
		if val == lootpack {
			index = i
			break
		}
	}
	if index == -1 {
		panic("Trying to remove non-existing unit!")
	}
	lss.ActiveLootPacks = append(lss.ActiveLootPacks[:index], lss.ActiveLootPacks[index+1:]...)
	RemoveFromRenderSystem(lss.World, lootpack)
}

func CreateLootboxPickupDialog(unit *entities.Unit, lootpack *entities.Lootpack) *ui.Dialog {
	dialogPosition := engo.AABB{Min: engo.Point{X: 100, Y: 100}, Max: engo.Point{X: 400, Y: 400}}
	d := ui.NewDialog(dialogPosition)

	offset := 0
	for item, amount := range *lootpack.Inventory {
		item := item  // To avoid re-using the memory
		btnPosition := engo.AABB{
			Min: engo.Point{X: 120, Y: float32(120 + offset)},
			Max: engo.Point{X: 380, Y: float32(139 + offset)},
		}
		event := events.RequestLootItemPickupEvent{
			Unit:     unit,
			Lootpack: lootpack,
			Item:     &item,
		}
		btn := ui.NewButton(btnPosition, item.Name+": "+strconv.Itoa(amount), event)
		d.AddElement(btn)
		offset += 40
	}
	return d
}

func (lss *LootManagementSystem) Update(dt float32) {}

func (lss *LootManagementSystem) Remove(e ecs.BasicEntity) {}
