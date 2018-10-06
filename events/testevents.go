package events

const (
	TEST_SIMPLE_ATTACK = "TestSimpleAttackEvent"
	TEST_KILL_AND_LOOT = "TestKillAndLootEvent"
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