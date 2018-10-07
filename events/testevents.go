package events

const (
	TEST_SIMPLE_ATTACK = "TestSimpleAttackEvent"
	TEST_KILL_AND_LOOT = "TestKillAndLootEvent"
	TEST_COLLISON_EVENT = "TestCollisionEvent"
	TEST_NOAP_ATTACK_EVENT = "TestNoAPAttackEvent"
	TEST_LOOT_TOO_FAR_EVENT = "TestLootTooFarEvent"

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

type TestNoAPAttackEvent struct{}

func (se TestNoAPAttackEvent) Type() string { return TEST_NOAP_ATTACK_EVENT }

func (se TestNoAPAttackEvent) AsLogMessage() string {
	return "No AP attack test"
}

type TestLootTooFarEvent struct{}

func (se TestLootTooFarEvent) Type() string { return TEST_LOOT_TOO_FAR_EVENT }

func (se TestLootTooFarEvent) AsLogMessage() string {
	return "Loot too far test"
}