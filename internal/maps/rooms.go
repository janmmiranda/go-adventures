package maps

import (
	"github.com/janmmiranda/go_adventures/internal/entities"
)

func CreateRoom(npcs []entities.Entity, description string, treasures []entities.Item,
	combatants []entities.Entity, forcedCB bool, start bool, end bool) RoomStruct {
	return RoomStruct{
		NPCs:         npcs,
		Description:  description,
		Treasures:    treasures,
		Combatants:   combatants,
		ForcedCombat: forcedCB,
		Start:        start,
		End:          end,
	}
}
