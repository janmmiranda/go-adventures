package entities

type Entity interface {
	getInventory() []Item
	// useItem() (bool, error)
	getInitiative() int
	getModifier(ability int) int
	talk()
	getHealth() int
	//skills() []Skill
	//spells() []Spell
}
