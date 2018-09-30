package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"errors"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type LootManagementSystem struct {
	*ecs.World
	ActiveLootPacks []*entities.Lootpack
}

func (lss *LootManagementSystem) New(world *ecs.World) {
	lss.World = world
	events.Mailbox.Listen(events.LOOT_REQUEST_SPAWN_EVENT, lss.getHandleRequestLootSpawn())
}

func (lss *LootManagementSystem) getHandleRequestLootSpawn() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		udmsg, ok := msg.(events.RequestLootSpawn)
		if !ok {
			return
		}

		lss.createRandomLoot(&udmsg.Point)
	}
}

func (lss *LootManagementSystem) createRandomLoot(position *engo.Point) {
	lootpack := entities.NewLootpack(position)
	lss.ActiveLootPacks = append(lss.ActiveLootPacks, lootpack)
	AddToRenderSystem(lss.World, lootpack)
}

func (lss *LootManagementSystem) FindLootUnderMouse(tracker *events.MouseTracker) (*entities.Lootpack, error) {
	for _, pack := range lss.ActiveLootPacks {
		xDelta := tracker.MouseX - pack.GetSpaceComponent().Position.X
		yDelta := tracker.MouseY - pack.GetSpaceComponent().Position.Y
		if xDelta > 0 && xDelta < entities.LOOTPACKSIZE && yDelta > 0 && yDelta < entities.LOOTPACKSIZE {
			return pack, nil
		}
	}
	return nil, errors.New("No lootpack Found")
}

func (lss *LootManagementSystem) Update(dt float32) {}

func (lss *LootManagementSystem) Remove(e ecs.BasicEntity) {}
