package rpg

import (
	"fmt"
	"os"
)

var (
	gold            int = 0
	playerInventory *Inventory
)

type Player struct {
	X, Y       int
	CurrentMap *Map
	Name       string
	Hp         int
	HpMax      int
	Attack     int
}

func NewPlayer(name string, hp, hpMax, attack int) *Player {
	fmt.Print(player)
	return &Player{
		Name:   name,
		Hp:     hp,
		HpMax:  hpMax,
		Attack: attack,
	}
}

func AddItemStatsToPlayer(itemShop *ItemShop) {
	player.Hp += itemShop.Effectarmure
	player.Attack += itemShop.Effectattack
}

func (p *Player) SetCurrentMap(newMap *Map) {
	p.CurrentMap = newMap
}

// Définir la carte 0 (forestMap) comme la carte par défaut
var Defaut = NewMap("0", 6, 7)

var player = Player{
	X:          1,
	Y:          1,
	CurrentMap: Defaut, // Utiliser la carte 0 comme carte par défaut
}

func (p *Player) Move(dx, dy int) bool {
	newX := p.X + dx
	newY := p.Y + dy
	currentMap := p.CurrentMap

	if newX >= 0 && newX < currentMap.Width && newY >= 0 && newY < currentMap.Height {
		if currentMap.Tiles[newY][newX] == '#' {
			return false
		} else {
			for _, teleporter := range currentMap.Teleporters {
				if teleporter.X == newX && teleporter.Y == newY && CountMobsOnMap(currentMap) == 0 {
					p.X = teleporter.DestinationX
					p.Y = teleporter.DestinationY
					p.SetCurrentMap(teleporter.DestinationMap) // Changer la currentMap du joueur
					return true
				}
			}

			for _, shop := range currentMap.Shop {
				if shop.X == newX && shop.Y == newY {

					if shop.num == 1 {
						ShopChoice()
					}
					return true
				}
			}

			for _, mob := range currentMap.Mobs {
				if mob.X == newX && mob.Y == newY {
					CombatLoop(currentMap, &mob)
					return true
				}
			}
			p.X = newX
			p.Y = newY
		}
	}
	return false
}

func ReadInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

func ProcessInput(input string, playerMap *Map) bool {
	switch input {
	case "z", "Z":
		return player.Move(0, -1)
	case "q", "Q":
		return player.Move(-1, 0)
	case "s", "S":
		return player.Move(0, 1)
	case "d", "D":
		return player.Move(1, 0)
	case "l", "L":
		os.Exit(0)
	case "p", "pause", "Pause":
		Pause()
	case "inv", "Inv":
		playerInventory = NewInventory()
		playerInventory.ShowInventory(&player)
	}
	return false
}
