package rpg

import (
	"fmt"
	"os"
)

var choix int

func MainMenu() {
	ClearScreen()
	fmt.Println("                                                      ")
	fmt.Println("                   ..:.                                ")
	fmt.Println("                  .==+*#+==-.                          ")
	fmt.Println("                 -*%%%%%*+====:                        ")
	fmt.Println("               .+#%%##%%%#*+=---:                      ")
	fmt.Println("               #%%#=.=%%%%%+:**+*+                 	")
	fmt.Println("               #%.  :#@@%%%%=+**+*++++=-:.       		")
	fmt.Println("                =   +%%##***#%%@@@@@@@@%%%+      		")
	fmt.Println("              .  .:=**#%%@@@@@%%%%@@@@@%%%=    		")
	fmt.Println("        .=-.......+@@@@@@%@@%*=-=#%%@%#%-          	")
	fmt.Println("       .+**+:==:.:=%%%%%#++=**+==*+#*+..      			")
	fmt.Println("        :###-=++=++%%@@@@@%%%#%%@##%#         			")
	fmt.Println("      .-*###*+++--=@@@@@%######%%#*##*-     			")
	fmt.Println("     +++####*++*--=@@@%@%%@%@%%##%%%%%%-     :#%#*-    ")
	fmt.Println("     +++####*+=*:=-%@@@%#%%###*#***####%+*###%##+==    ")
	fmt.Println("      +*#*#***++=--+@@@@%%%%##%####%##%%%%*%%%@#*=-    ")
	fmt.Println("       .##*****#####%@@@@@%%%@@##%%##%@#@%##%#%%#*:    ")
	fmt.Println("         =%%%%###%%%@@%%@@@@@@@%#%%%%@@*@%%#*%%#*=+*   ")
	fmt.Println("          .-+++++==@@@%###%%@@%%##%%%%%==+==-%%#*=-+.  ")
	fmt.Println("                  :*##%%%##%%@@%%%%##%@%++- .%%#*=-==  ")
	fmt.Println("                -+###*%%%%@%%#*#+**@%%%%@**=+%%#*+-=*- ")
	fmt.Println("               :%%@%%%##%%@@@+#%##-@%@%@@%#=%@@%%#**+: ")
	fmt.Println("               *@##%%%#@%%%%%#*##+*##%@@%%#-     		")
	fmt.Println("               *@@@%%%#@@@**+*%*#*#%%%%@@##*   	 	")
	fmt.Println("                 -@#*%%%@@%#+*#++*%%%%*-*+-.  			")
	fmt.Println("                 =*:*#%%@@@%****%@@%%%#=-  			")
	fmt.Println("                  +%%%%%%%@= .:.%@@@%%%#%+  			")
	fmt.Println("                  -=+**+:.        .:=+==-:				")
	fmt.Println("																				")
	fmt.Println("																				")
	fmt.Println("██╗      █████╗     ██████╗ ██╗   ██╗██╗   ██╗███████╗████████╗████████╗███████╗")
	fmt.Println("██║     ██╔══██╗    ██╔══██╗██║   ██║██║   ██║██╔════╝╚══██╔══╝╚══██╔══╝██╔════╝")
	fmt.Println("██║     ███████║    ██████╔╝██║   ██║██║   ██║█████╗     ██║      ██║   █████╗  ")
	fmt.Println("██║     ██╔══██║    ██╔══██╗██║   ██║╚██╗ ██╔╝██╔══╝     ██║      ██║   ██╔══╝  ")
	fmt.Println("███████╗██║  ██║    ██████╔╝╚██████╔╝ ╚████╔╝ ███████╗   ██║      ██║   ███████╗")
	fmt.Println("╚══════╝╚═╝  ╚═╝    ╚═════╝  ╚═════╝   ╚═══╝  ╚══════╝   ╚═╝      ╚═╝   ╚══════╝")
	fmt.Println("						" + Red + "CLOSED OPEN BETA v1.0.0 " + Reset + "	")
	fmt.Println("			Projet RPG en Go réalisé par: Alexy / Mathieu / Nicolas / Gabriel")
	fmt.Println("																				")
	fmt.Println(" ")
	fmt.Println("Menu Principal: ")
	fmt.Println(" - 1. Play")
	fmt.Println(" - 2. Quit")
	fmt.Println("")

	for {
		fmt.Print("Choisissez une option (1/2): ")
		_, err := fmt.Scan(&choix)

		if err != nil {
			fmt.Println("\n Choix invalide. Veuillez choisir 1 ou 2.")
			continue
		}

		switch choix {
		case 1:
			Cinestart(&player)
			return
		case 2:
			fmt.Println("Vous avez choisi l'Option 2. Fermeture de la fenêtre en cours...")
			os.Exit(0)
		default:
			fmt.Println("\n Choix invalide. Veuillez choisir 1 ou 2.")
		}
	}
}
