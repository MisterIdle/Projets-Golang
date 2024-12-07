package rpg

import (
	"fmt"
	"os"
	"os/exec"
)

var choice int

func GameOver() {
	ClearScreen()
	fmt.Println("  ___   _                 ____                         _ ")
	fmt.Println(" / _ \\ | |__             / ___|  _ __    __ _  _ __   | |")
	fmt.Println("| | | || '_ \\            \\___ \\ | '_ \\  / _` || '_ \\  | |")
	fmt.Println("| |_| || | | |  _  _  _   ___) || | | || (_| || |_) | |_|")
	fmt.Println(" \\___/ |_| |_| (_)(_)(_) |____/ |_| |_| \\__,_|| .__/  (_)")
	fmt.Println("                                              |_|        ")
	fmt.Println("\n")
	fmt.Println(" __     __                           _                                        _   ")
	fmt.Println(" \\ \\   / /  ___   _   _  ___    ___ | |_   ___  ___   _ __ ___    ___   _ __ | |_ ")
	fmt.Println("  \\ \\ / /  / _ \\ | | | |/ __|  / _ \\| __| / _ \\/ __| | '_ ` _ \\  / _ \\ | '__|| __|")
	fmt.Println("   \\ V /  | (_) || |_| |\\__ \\ |  __/| |_ |  __/\\__ \\ | | | | | || (_) || |   | |_ ")
	fmt.Println("    \\_/    \\___/  \\__,_||___/  \\___| \\__| \\___||___/ |_| |_| |_| \\___/ |_|    \\__|\n")
	fmt.Println("\n          1. Reessayer !")
	fmt.Println("          2. Quitter le jeu")
	fmt.Println("")
	for {
		fmt.Print("Choisissez une option (1/2): ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("\n Choix invalide. Veuillez choisir 1 ou 2.")
			continue
		}
		switch choice {
		case 1:
			RestartCMD()
		case 2:
			fmt.Println("Vous avez choisi l'Option 2. Fermeture de la fenêtre en cours...")
			os.Exit(0)
		default:
			fmt.Println("\n Choix invalide. Veuillez choisir 1 ou 2.")
		}
	}
}

// temporaire pour relancer le jeu sauf si il manque du temps
func RestartCMD() {
	cmd := exec.Command("cmd", "/c", "start", "cmd", "/c", "go", "run", "main.go")
	cmd.Start()
	os.Exit(0)
}
