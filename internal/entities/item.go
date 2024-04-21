package entities

type Item interface {
	inspect()
	use() (bool, string)
	count() int
}

type Weapon interface {
	attack(target Entity, user Entity) (bool, error)
	specialAttack(target Entity, user Entity) (bool, error)
}

type Armor interface {
	specialSkill(arget Entity, user Entity) (bool, error)
}
