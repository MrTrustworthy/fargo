package entities

type Ability interface {
	Source() *Unit
	Target() *Unit
	SetTarget(*Unit)
	Maxrange() float32
	Name() string
	AnimationName() string
	Execute()
}

type StabAbility struct {
	source   *Unit
	target   *Unit
	maxrange float32
}

func NewStabAbility(unit *Unit) *StabAbility {
	return &StabAbility{
		source:   unit,
		maxrange: 140,
	}
}

func (sa *StabAbility) Source() *Unit {
	return sa.source
}

func (sa *StabAbility) Target() *Unit {
	return sa.target
}
func (sa *StabAbility) SetTarget(unit *Unit) {
	sa.target = unit
}

func (sa *StabAbility) Maxrange() float32 {
	return sa.maxrange
}

func (sa *StabAbility) Name() string {
	return "StabAbility"
}

func (sa *StabAbility) AnimationName() string {
	return "stab"
}

func (sa *StabAbility) Execute() {
	sa.source.AnimationComponent.SelectAnimationByName("stab")
}
