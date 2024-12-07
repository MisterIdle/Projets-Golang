package rpg

var globalPlayer *Player

func InitializeMaps() map[string]*Map {
	maps := make(map[string]*Map)
	globalPlayer = NewPlayer(player.Name, 50, 50, 5)
	Nbsoin = 1
	//floor 1
	FloorOne := NewMap("1", 7, 7)
	MakeAndPlaceTeleporter(Defaut, 3, 2, 1, 1, FloorOne, 'T')
	MakeAndPlaceMob(FloorOne, 4, 2, 's', "Slime", 4, 4, 1, 4)

	//floor 2
	FloorTwo := NewMap("2", 10, 7)
	MakeAndPlaceTeleporter(FloorOne, 5, 5, 1, 1, FloorTwo, 'T')
	MakeAndPlaceTeleporter(FloorTwo, 5, 4, 1, 1, FloorOne, 't')
	MakeAndPlaceMob(FloorTwo, 4, 2, 's', "Slime", 4, 4, 1, 4)
	MakeAndPlaceMob(FloorTwo, 3, 5, 's', "Slime", 4, 4, 1, 4)
	MakeAndPlaceMob(FloorTwo, 6, 4, 's', "Slime", 4, 4, 1, 4)
	MakeAndPlaceMob(FloorTwo, 6, 3, 'g', "Goblin", 6, 6, 2, 10)

	//floor 3
	FloorThree := NewMap("3", 11, 11)
	MakeAndPlaceTeleporter(FloorTwo, 7, 5, 1, 1, FloorThree, 'T')
	MakeAndPlaceTeleporter(FloorThree, 7, 4, 1, 1, FloorTwo, 't')
	MakeAndPlaceMob(FloorThree, 2, 5, 's', "Slime", 4, 4, 1, 4)
	MakeAndPlaceMob(FloorThree, 4, 5, 's', "Slime", 4, 4, 1, 4)
	MakeAndPlaceMob(FloorThree, 6, 4, 'g', "Goblin", 6, 6, 2, 10)
	MakeAndPlaceMob(FloorThree, 5, 3, 'S', "Giga Slime", 10, 10, 3, 15)

	//floor 4
	FloorFour := NewMap("4", 12, 12)
	MakeAndPlaceTeleporter(FloorThree, 9, 6, 1, 1, FloorFour, 'T')
	MakeAndPlaceTeleporter(FloorFour, 9, 5, 1, 1, FloorThree, 't')
	MakeAndPlaceShop(FloorFour, 6, 5, 'M', 1)
	MakeAndPlaceMob(FloorFour, 3, 7, 's', "Slime", 4, 4, 1, 2)
	MakeAndPlaceMob(FloorFour, 5, 5, 'g', "Goblin", 6, 6, 2, 10)
	MakeAndPlaceMob(FloorFour, 6, 6, 'g', "Goblin", 6, 6, 2, 10)
	MakeAndPlaceMob(FloorFour, 7, 5, 'g', "Goblin", 6, 6, 2, 10)
	MakeAndPlaceMob(FloorFour, 9, 3, 'G', "Gros Goblin", 17, 17, 5, 25)

	//floor 5
	FloorFive := NewMap("5", 13, 13)
	MakeAndPlaceTeleporter(FloorFour, 10, 7, 1, 1, FloorFive, 'T')
	MakeAndPlaceTeleporter(FloorFive, 10, 6, 1, 1, FloorFour, 't')
	MakeAndPlaceMob(FloorFive, 3, 5, 'S', "Giga Slime", 10, 10, 3, 5)
	MakeAndPlaceMob(FloorFive, 5, 4, 'S', "Giga Slime", 10, 10, 3, 5)
	MakeAndPlaceMob(FloorFive, 7, 4, 'S', "Giga Slime", 10, 10, 3, 5)
	MakeAndPlaceMob(FloorFive, 9, 6, 'b', "Bandi", 14, 14, 4, 20)
	MakeAndPlaceMob(FloorFive, 10, 3, 'b', "Bandi", 14, 14, 4, 20)

	//floor 6
	FloorSix := NewMap("6", 10, 8)
	MakeAndPlaceTeleporter(FloorFive, 11, 5, 1, 1, FloorSix, 'T')
	MakeAndPlaceTeleporter(FloorSix, 11, 4, 1, 1, FloorFive, 't')
	MakeAndPlaceMob(FloorSix, 3, 3, 'S', "Giga Slime", 10, 10, 3, 5)
	MakeAndPlaceMob(FloorSix, 5, 5, 'b', "Bandi", 14, 14, 4, 20)
	MakeAndPlaceMob(FloorSix, 7, 4, 'b', "Bandi", 14, 14, 4, 20)
	MakeAndPlaceMob(FloorSix, 4, 4, 'B', "Chef Bandi", 20, 20, 10, 10)

	//floor 7
	FloorSeven := NewMap("7", 13, 12)
	MakeAndPlaceTeleporter(FloorSix, 7, 5, 1, 1, FloorSeven, 'T')
	MakeAndPlaceTeleporter(FloorSeven, 7, 4, 1, 1, FloorSix, 't')
	MakeAndPlaceMob(FloorSeven, 3, 3, 'S', "Giga Slime", 10, 10, 3, 5)
	MakeAndPlaceMob(FloorSeven, 6, 5, 'b', "Bandi", 14, 14, 4, 20)
	MakeAndPlaceMob(FloorSeven, 4, 6, 'B', "Chef Bandi", 20, 20, 10, 10)
	MakeAndPlaceMob(FloorSeven, 7, 6, 'l', "Loup", 18, 18, 8, 30)

	//floor 8
	FloorEight := NewMap("8", 11, 12)
	MakeAndPlaceTeleporter(FloorSeven, 8, 7, 1, 1, FloorEight, 'T')
	MakeAndPlaceTeleporter(FloorEight, 8, 6, 1, 1, FloorSeven, 't')
	MakeAndPlaceShop(FloorEight, 4, 4, 'M', 1)
	MakeAndPlaceMob(FloorEight, 3, 2, 'B', "Chef Bandi", 20, 20, 10, 10)
	MakeAndPlaceMob(FloorEight, 4, 4, 'l', "Loup", 18, 18, 8, 30)
	MakeAndPlaceMob(FloorEight, 5, 3, 'l', "Loup", 18, 18, 8, 30)

	//floor 9
	FloorNine := NewMap("9", 13, 10)
	MakeAndPlaceTeleporter(FloorEight, 6, 4, 1, 1, FloorNine, 'T')
	MakeAndPlaceTeleporter(FloorNine, 6, 3, 1, 1, FloorEight, 't')
	MakeAndPlaceMob(FloorNine, 3, 6, 'l', "Loup", 18, 18, 8, 30)
	MakeAndPlaceMob(FloorNine, 5, 5, 'l', "Loup", 18, 18, 8, 30)
	MakeAndPlaceMob(FloorNine, 7, 3, 'l', "Loup", 18, 18, 8, 30)

	//floor 10
	FloorTen := NewMap("Boss", 14, 14)
	MakeAndPlaceTeleporter(FloorNine, 8, 6, 1, 1, FloorTen, 'T')
	MakeAndPlaceTeleporter(FloorTen, 8, 5, 1, 1, FloorNine, 't')
	MakeAndPlaceMob(FloorTen, 2, 7, 'L', "Infect Père des Loup", 100, 100, 5, 1000)

	maps[FloorOne.Name] = FloorOne
	maps[FloorTwo.Name] = FloorTwo
	maps[FloorThree.Name] = FloorThree
	maps[FloorFour.Name] = FloorFour
	maps[FloorFive.Name] = FloorFive
	maps[FloorSix.Name] = FloorSix
	maps[FloorSeven.Name] = FloorSeven
	maps[FloorEight.Name] = FloorEight
	maps[FloorNine.Name] = FloorNine
	maps[FloorTen.Name] = FloorTen
	return maps
}
