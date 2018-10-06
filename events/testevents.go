package events

const (
	TEST_SIMPLE_ATTACK = "TestSimpleAttackEvent"
)

type TestBasicAttackEvent struct{}

func (se TestBasicAttackEvent) Type() string { return TEST_SIMPLE_ATTACK }

func (se TestBasicAttackEvent) AsLogMessage() string {
	return "Simple attack test"
}
