package events

const (
	TEST_SIMPLE_ATTACK = "TestSimpleAttackEvent"
	TEST_KILL_AND_LOOT = "TestKillAndLootEvent"
	TEST_COLLISON_EVENT = "TestCollisionEvent"
)

type TestBasicAttackEvent struct{}

func (se TestBasicAttackEvent) Type() string { return TEST_SIMPLE_ATTACK }

func (se TestBasicAttackEvent) AsLogMessage() string {
	return "Simple attack test"
}

type TestKillAndLootEvent struct{}

func (se TestKillAndLootEvent) Type() string { return TEST_KILL_AND_LOOT }

func (se TestKillAndLootEvent) AsLogMessage() string {
	return "Kill and loot test"
}

type TestCollisionEvent struct{}

func (se TestCollisionEvent) Type() string { return TEST_COLLISON_EVENT }

func (se TestCollisionEvent) AsLogMessage() string {
	return "Collision test"
}