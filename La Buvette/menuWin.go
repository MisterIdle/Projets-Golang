package rpg

import (
	"fmt"
	"os"
)

func Win() {
	for {
		fmt.Println("                                                                                         __     ")
		fmt.Println("__     __                                                                               /_/   _ ")
		fmt.Println("\\ \\   / /  ___   _   _  ___     __ _ __   __  ___  ____    __ _   __ _   __ _  _ __    ___   | |")
		fmt.Println(" \\ \\ / /  / _ \\ | | | |/ __|   / _` |\\ \\ / / / _ \\|_  /   / _` | / _` | / _` || '_ \\  / _ \\  | |")
		fmt.Println("  \\ V /  | (_) || |_| |\\__ \\  | (_| | \\ V / |  __/ / /   | (_| || (_| || (_| || | | ||  __/  |_|")
		fmt.Println("   \\_/    \\___/  \\__,_||___/   \\__,_|  \\_/   \\___|/___|   \\__, | \\__,_| \\__, ||_| |_| \\___|  (_)")
		fmt.Println("                                                          |___/         |___/                   ")
		fmt.Println("\n\n")
		fmt.Println("                      ░▓▓  ▓▓▓     ")
		fmt.Println("                ░░▒  ░▓▓ ▒▓▓▓▓▓    ")
		fmt.Println("               ░▒███░▓▓▒ ░▓▓▓█▓▓▓  ")
		fmt.Println("                ░▒█░▓▓▒░░▒▓▓███▓▓  ")
		fmt.Println("                  ░▓▒▓░▒▓▓████▓▓▓▓ ")
		fmt.Println("                 ░▓▒▓▒▓▓▓████▓▓▓▓  ")
		fmt.Println("                ░▓▒ ▓▓▓█████▓▓▓    ")
		fmt.Println("               ░▓▒ ▒▓▓▓███▓▓▓      ")
		fmt.Println("              ░▓▒ ▒▓▓▓▓▓▓          ")
		fmt.Println("             ░▓▒                   ")
		fmt.Println("            ░▓▒      ▒▓▒           ")
		fmt.Println("           ░▓▒     ▒▓▒▒▓▓▓▒        ")
		fmt.Println("          ░▓▒    ▒▒▒▒░░░▒▓▓▓▒      ")
		fmt.Println("         ░▓▒    ▒▓▓▒░░▒▒▓▓▓█▓▒     ")
		fmt.Println("        ▒▓▒    ▓▓▓▓▒░▒▒▓█████▓▒    ")
		fmt.Println("       ▒▓▓▓    ▓█▓▓▓███▓▓███▓▓█▓   ")
		fmt.Println("       ░▓██    ▒▓▓██▓░▓▓██░██▓▓▓   ")
		fmt.Println("      ▒▓ ███  ▓▓▓▓▓▓▓████▓▓▓▓▓▓▓   ")
		fmt.Println("     ▒▓▒ ▓█████▓▓▓▓▓▓▓▓▓█▓▓▓▓▓▓    ")
		fmt.Println("      ░   ▓██▓██▓▓▓▓▓▓▓██▓▓▓▓▓     ")
		fmt.Println("            ▒▓███▓▓▓▓█████▓▓▓█▓█   ")
		fmt.Println("             █▓█████████████▓ ▓▓█  ")
		fmt.Println("              ░███████████▓▓▓ █▓█  ")
		fmt.Println("              ▒█████▓▓███████  ██  ")
		fmt.Println("              ▓███▓██▓█▓▓█▓▒   ▓▓▓ ")
		fmt.Println("              ▒████▓▓▓████▒        ")
		fmt.Println("              ▒▓▓▓▓▒ ▒██▓█▒        ")
		fmt.Println("               ▒▒▒     ▒▒▒         ")
		fmt.Println("\n\n\n            1. Recommencer une partie")
		fmt.Println("            2. Quitter le jeu")

		var choice int
		fmt.Println("Choisissez une option : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Recommencer une partie")
		case 2:
			fmt.Println("Quitter le jeu.")
			os.Exit(1)
		default:
			fmt.Println("Option Invalide")
		}
	}
}
