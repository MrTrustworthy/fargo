package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"github.com/MrTrustworthy/fargo/entities"
	"github.com/MrTrustworthy/fargo/events"
)

type LootSpawnSystem struct {
	*ecs.World
	dyingUnit *entities.Unit
}

func (lss *LootSpawnSystem) New(world *ecs.World) {
	lss.World = world
	events.Mailbox.Listen(events.LOOT_REQUEST_SPAWN_EVENT, lss.getHandleRequestLootSpawn())
}

func (lss *LootSpawnSystem) getHandleRequestLootSpawn() func(msg events.BaseEvent) {
	return func(msg events.BaseEvent) {
		udmsg, ok := msg.(events.RequestLootSpawn)
		if !ok {
			return
		}

		lss.createRandomLoot(&udmsg.Point)
	}
}

func (lss *LootSpawnSystem) createRandomLoot(position *engo.Point) {
	lootpack := entities.NewLootpack(position)
	AddToRenderSystem(lss.World, lootpack)
}

func (lss *LootSpawnSystem) Update(dt float32) {

}

func (lss *LootSpawnSystem) Remove(e ecs.BasicEntity) {}
