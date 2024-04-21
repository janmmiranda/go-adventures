package maps

import (
	"github.com/janmmiranda/go_adventures/internal/entities"
)

type Map struct {
	Rooms        []room
	VisitedRooms []room
}

type room interface {
	hasCombat() (bool, bool) //has combat, has forced combat
	getNpcs() []entities.Entity
	getTreasure() []entities.Item
}

type RoomStruct struct {
	NPCs         []entities.Entity
	Description  string
	Treasures    []entities.Item
	Combatants   []entities.Entity
	LinkedRooms  []room
	ForcedCombat bool
	Start        bool
	End          bool
}

func (rs *RoomStruct) hasCombat() (bool, bool) {
	return len(rs.Combatants) > 0, rs.ForcedCombat
}

func (rs *RoomStruct) getNpcs() []entities.Entity {
	return rs.NPCs
}

func (rs *RoomStruct) getTreasure() []entities.Item {
	return rs.Treasures
}
