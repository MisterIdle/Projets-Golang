package rpg

import (
	"fmt"
	"os"
)

func Pause() {
	ClearScreen()
	for {
		fmt.Println("___  ___                          _                                            ")
		fmt.Println("|  \\/  |                         | |                                         _ ")
		fmt.Println("| .  . |  ___  _ __   _   _    __| |  ___   _ __    __ _  _   _  ___   ___  (_)")
		fmt.Println("| |\\/| | / _ \\| '_ \\ | | | |  / _` | / _ \\ | '_ \\  / _` || | | |/ __| / _ \\    ")
		fmt.Println("| |  | ||  __/| | | || |_| | | (_| ||  __/ | |_) || (_| || |_| |\\__ \\|  __/  _ ")
		fmt.Println("\\_|  |_/ \\___||_| |_| \\__,_|  \\__,_| \\___| | .__/  \\__,_| \\__,_||___/ \\___| (_)")
		fmt.Println("                                           | |                                 ")
		fmt.Println("                                           |_|                                 ")
		fmt.Println("\n\n       1. Reprendre la partie")
		fmt.Println("       2. Quitter la partie")

		var choix string
		fmt.Print("\nQue voulez-vous faire ? (1/2)\n")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			fmt.Println("\nReprendre la partie")
			return
		case "2":
			fmt.Println("\nQuitter la partie")
			os.Exit(1)
		case "pause":
			Pause()
		default:
			fmt.Println("\nChoix invalide. Veuillez choisir une option valide.")
		}
	}
}
