package rpg

import (
	"math/rand"
)

type Mob struct {
	X, Y          int
	Appearance    rune
	Direction     int
	Name          string
	Hp            int
	HpMax         int
	Attack        int
	LootableItems []*LootableItem
	dropGold      int
}
type LootableItem struct {
	Name        string
	Description string
	DropChance  float64
	Quantity    int
}

func MakeAndPlaceMob(currentMap *Map, MobX, MobY int, Appearance rune, Name string, Hp int, HpMax, Attack int, dropGold int) {
	direction := rand.Intn(4)
	mob := Mob{
		X:          MobX,
		Y:          MobY,
		Appearance: Appearance,
		Direction:  direction,
		Name:       Name,
		Hp:         Hp,
		Attack:     Attack,
		dropGold:   dropGold,
	}

	currentMap.Mobs = append(currentMap.Mobs, mob)
}

func CountMobsOnMap(currentMap *Map) int {
	return len(currentMap.Mobs)
}

func UpdateMobDirections(currentMap *Map, PlayerX, PlayerY int) {
	for i := range currentMap.Mobs {
		mobX, mobY := currentMap.Mobs[i].X, currentMap.Mobs[i].Y
		direction := rand.Intn(4)
		newX, newY := mobX, mobY
		switch direction {
		case 0:
			newY--
		case 1:
			newX++
		case 2:
			newY++
		case 3:
			newX--
		}
		if newX >= 0 && newX < currentMap.Width && newY >= 0 && newY < currentMap.Height {
			if currentMap.Tiles[newY][newX] != '#' &&
				currentMap.Tiles[newY][newX] != 'T' &&
				currentMap.Tiles[newY][newX] != 'S' &&
				currentMap.Tiles[newY][newX] != 'B' &&
				!(newX == PlayerX && newY == PlayerY) { // Vérifie que le mob ne se déplace pas vers le joueur
				currentMap.Mobs[i].Direction = direction
				currentMap.Mobs[i].X = newX
				currentMap.Mobs[i].Y = newY
			}
		}
	}
}
