package rpg

import (
	"fmt"
	"strconv"
	"strings"
)

var selectedItem *ItemShop
var Nbheal int

type ItemShop struct {
	Name         string
	Price        int
	Effectarmure int
	Effectattack int
	Effectheal   int
	Quantity     int
}

func ShopChoice() {
	Nbheal = Nbsoin
	playerInventory := NewInventory()
	// Supprimez les potions du Shop1
	Shop1 := []ItemShop{
		{"1 - Hache en fer", 30, 0, 2, 0, 1},
		{"2 - Hache en or", 100, 0, 10, 0, 1},
		{"3 - Armure en maille", 30, 15, 0, 0, 1},
		{"4 - Armure en fer", 100, 10, 0, 0, 1},
		{"5 - Potion", 2, 0, 0, 15, 5},
	}
	// Ajoutez un compteur pour les soins
	for {
		Shoprint()

		switch 1 {
		case 1:
			for _, item := range Shop1 {
				fmt.Printf("%s, Price: %d, Armure: %d, Attaque: %d, Soins: %d, Quantity: %d\n", item.Name, item.Price, item.Effectarmure, item.Effectattack, item.Effectheal, item.Quantity)
			}
		}
		// Ajoutez également l'affichage du nombre de soins disponibles

		// Demandez au joueur de choisir un article
		fmt.Println(" ")
		fmt.Println("―――――――――――――――――――――――――――――――")
		fmt.Println("Or:", gold)
		fmt.Println("―――――――――――――――――――――――――――――――")
		fmt.Println("\nQuel article souhaitez-vous acheter ?")
		fmt.Println("(1-5) ou 'quitter' pour quitter le shop: ")

		var choix string
		fmt.Scan(&choix)

		// Vérifiez si le joueur veut quitter le shop
		choix = strings.ToLower(choix)
		if choix == "quitter" || choix == "q" || choix == "Q" {
			fmt.Println("Vous avez quitté le shop.")
			break
		}

		// Convertissez le choix en un entier
		numChoix, err := strconv.Atoi(choix)
		if err != nil || numChoix < 1 || numChoix > 5 {
			fmt.Println("Cette action est interdite ici.")
			continue
		}

		// Récupérez l'article sélectionné
		var selectedItem *ItemShop
		switch numChoix {
		case 1:
			selectedItem = &Shop1[0]
		case 2:
			selectedItem = &Shop1[1]
		case 3:
			selectedItem = &Shop1[2]
		case 4:
			selectedItem = &Shop1[3]
		case 5:
			// Option pour ajouter 1 au compteur de soins (Nbsoin) avec un coût de 2
			if gold >= 2 {
				Nbsoin++
				gold -= 2 // Soustraire 2 au gold du joueur
				fmt.Printf("\n\n\nVous avez obtenu un soin supplémentaire pour 2 gold. Nombre total de soins : %d\n", Nbsoin)
			} else {
				fmt.Println("\n\n\nVous n'avez pas assez de gold pour acheter une potion de soin.")
			}
			continue
		}

		// Vérifiez si le joueur a assez d'or pour acheter l'article
		if gold >= selectedItem.Price && selectedItem.Quantity >= 1 {
			// Mettez à jour la quantité d'or du joueur et la quantité de l'article acheté
			gold -= selectedItem.Price
			selectedItem.Quantity--
			playerInventory.AddItemFromShop(selectedItem)
			AddItemStatsToPlayer(selectedItem)

			// Si l'article acheté est un soin, incrémentez le compteur de soins
			if selectedItem.Effectheal > 0 {
				Nbsoin++
			}

			//clear
			fmt.Printf("\n\n\nFélicitations, vous avez réussi votre achat ! Il vous reste maintenant %d gold.\n", gold)
		} else if gold < selectedItem.Price {
			fmt.Println("\n\n\nVous n'avez pas assez de gold.")
		} else {
			fmt.Println("\n\n\nRupture de stock !")
		}
	}
}

func Shoprint() {
	fmt.Println("                                ____   _                   ")
	fmt.Println("                               / ___| | |__    ___   _ __  ")
	fmt.Println("                               \\___ \\ | '_ \\  / _ \\ | '_ \\ ")
	fmt.Println("                                ___) || | | || (_) || |_) |")
	fmt.Println("                               |____/ |_| |_| \\___/ | .__/ ")
	fmt.Println("                                                    |_|\n")
	fmt.Println("                                   █     █▓▓▓▓▓      █")
	fmt.Println("                                  ▓░▓▓ ▓▓█████▓▓█▓ ▓▓░▓")
	fmt.Println("                                 ▓▒▒▒▓█▓▓▓▒▒▓▓▓▓▓▓▓▒▒▒▒▒▓█")
	fmt.Println("                               ▓▓▒▒▒░░░▒▓▓▓▓▓▓▓▓▒░░░░▒▒▒▒▓▓█")
	fmt.Println("                             ▓█░░░░░░░░░░░░░░░░▒▒▒▒▒░░░░░░▒█▓")
	fmt.Println("                           ▓▒▒▒▒▒▒▒░░░░░▒▒░░░░▒▓▓▒▓▓▓░▒▒▒▒▒▒▓█")
	fmt.Println("                             ▓███▓▓░░░▒▒▒▒░░░▒▒▓▓▒░▒▒▒▒▒▒██▓▓▓▓█")
	fmt.Println("                            ▓▒▒▒▒▒▒░░▒▒▓▒▒░░▒▒▓▒▓▓▒▒▒▒▒▓▓▓▓▓")
	fmt.Println("                           ▓▒▒▓▓▓▓▓░▒▒▓▒▒░▒▒▓▓▒░░▓▓▓▓░░▒▒▓█")
	fmt.Println("                          █▒████▒▒░▒▓▓▓▓▓▓▓▒▒░░▒▒▓▒▓▓▓▒▒▓▓██")
	fmt.Println("                           ▓███▓▓▓▓▓▒▒██▓▒░░░░░░▒██▓▒▒▒▓▓▓▓███")
	fmt.Println("                             ▓▓░▓▒▒▒░░██▒░░░░░░░░██▒░░▒▒▓▓▒▓█▓█")
	fmt.Println("                           ▓▒▒▓▒░▓▒▒▒░██▒░░░█░░░░██▒░░▒▓░░▒▓            █▓▓▓██")
	fmt.Println("                            █▒  ▓▒▒▒░░▒▒░░░░░░░░░▒▒░░░░▒▓               ▓▓▓███")
	fmt.Println("                                  ▓▓▒▒░░░░█████░░░░░▒▓▓                   ▓█")
	fmt.Println("                                    ██▓▒░░▒░░░▒░░▒▓▓██       █▓▓█       █▓▓▒▒▓▓")
	fmt.Println("      ▓███▓         ██     ██      ▓███▒▒▒▒▓▓▓▓▓▒▒▒███▓      ▒▓▓▓      ██▓▒▒▓▒▓▓")
	fmt.Println("      ▓███▓       ▓▓▓▓▓▓ ▓▓▓▓▓▓  ▒▓█▒▒░░▒█▓█▒▒███▒░░▒▓█▓      ▓▓      ▓▓▓░░░░░░▓▓")
	fmt.Println("       ▓█▓▓▓      ▓█▓▒░▓▒▒▓▒█▓▓ ▓▒░░▓██▒░▒██████▓░░███░░▓  ▓▒░░░▒▒▒▓  █▒░▒▒▒▒░▒▓▓")
	fmt.Println("      ▓█▒░░░██    ▓▓▒▒░▓▓▒░█▓▓▓ █▓▒▒▓██▓░░▓██▓█▓░░▒██▓▒▒▓▓ ░▒▒░░▓▒▓▓  ▓▒▒▒▒▒░░▒▓▓")
	fmt.Println("    ▓▓▒▒▒▒▒▒▒▒▓█  ▓█▓▒▒▓ ▒░█▒▓█ ▓▓▒▒███▒░░▓█▓▓█▓░░░███▓▓▓▒ ░░▓▓▓█▓██  █▓▓▒▒▒▒▒░▒█")
	fmt.Println("    ▓▒░░░░░░░░▓█   ▓▓▓█   ░▓░▓  █▒▒█▓▓░▒█░▓█▓▓█▓░█░░▓█▓▒▓▒ ▓▒███████  ▓▓██▒▒▒░▒█")
	fmt.Println(" ▒▓▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▒▓▓▓▓▓▒▒████▓▓▒▒█▒▓████▓▒█▒▒█████▓  ███████    █▓▓▓▓▓▓█")
	fmt.Println(" ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓")
	fmt.Println(" ▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓███")
	fmt.Println("   ▓█▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██\n")
}
