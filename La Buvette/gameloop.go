package rpg

func gameLoop() {

	for {
		ClearScreen()
		player.CurrentMap.PrintMap(player)
		playerInput := ReadInput()
		ProcessInput(playerInput, player.CurrentMap)
		UpdateMobDirections(player.CurrentMap, player.X, player.Y)
	}
}
