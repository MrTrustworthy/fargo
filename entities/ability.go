package entities

type Ability interface {
	Source() *Unit
	Target() *Unit
	SetTarget(*Unit)
	Maxrange() float32
	Name() string
	Damage() int
	Cost() int
	AnimationName() string
	CanExecute() bool
}

type SimpleTargetAbility struct {
	source    *Unit
	target    *Unit
	cost      int
	name      string
	animation string
	maxrange  float32
	damage    int
}

func NewStabAbility(unit *Unit) *SimpleTargetAbility {
	return &SimpleTargetAbility{
		source:    unit,
		maxrange:  140,
		name:      "SimpleTargetAbility",
		animation: "stab",
		damage:    4,
		cost:      2,
	}
}

func (sa *SimpleTargetAbility) Source() *Unit {
	return sa.source
}

func (sa *SimpleTargetAbility) Target() *Unit {
	return sa.target
}
func (sa *SimpleTargetAbility) SetTarget(unit *Unit) {
	sa.target = unit
}

func (sa *SimpleTargetAbility) Maxrange() float32 {
	return sa.maxrange
}

func (sa *SimpleTargetAbility) Name() string {
	return sa.name
}

func (sa *SimpleTargetAbility) Cost() int {
	return sa.cost
}

func (sa *SimpleTargetAbility) Damage() int {
	return sa.damage
}

func (sa *SimpleTargetAbility) AnimationName() string {
	return sa.animation
}

func (sa *SimpleTargetAbility) CanExecute() bool {
	origin, target := sa.Source().Center(), sa.Target().Center()
	isInRange := origin.PointDistance(target) <= sa.Maxrange()
	hasEnoughAP := sa.Source().AP >= sa.cost
	return isInRange && hasEnoughAP
}
