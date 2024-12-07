package rpg

import (
	"fmt"
	"io/ioutil"
	"math/rand"
)

var Nbsoin int

func SupprimerMobAuxCoordonnees(currentMap *Map, x, y int) {
	for i, mob := range currentMap.Mobs {
		if mob.X == x && mob.Y == y {
			// Utilisation de la technique du swap pour retirer le mob de la liste.
			currentMap.Mobs[i] = currentMap.Mobs[len(currentMap.Mobs)-1]
			currentMap.Mobs = currentMap.Mobs[:len(currentMap.Mobs)-1]
			return
		}
	}
}

var globalInventory *Inventory

func CombatLoop(currentMap *Map, ennemi *Mob) bool {
	ennemi.HpMax = ennemi.Hp
	var Ordreatt bool
	var jmiss bool
	var emiss bool
	var success bool
	if rand.Intn(2) == 1 {
		Ordreatt = true
	} else {
		Ordreatt = false
	}
	for {
		if ennemi.Hp <= 0 {
			fmt.Println("Vous avez vaincu l'ennemi !")
			gold += ennemi.dropGold
			SupprimerMobAuxCoordonnees(currentMap, ennemi.X, ennemi.Y)
			if ennemi.Name == "Infect Père des Loup" {
				Win()
			}

			return true
		}
		if globalPlayer.Hp <= 0 {
			GameOver()
			return false
		}
		ClearScreen()
		fmt.Println(ennemi.Name + "\n")
		Affichemob(ennemi)
		affichecombatlog(jmiss, emiss)
		afficherInformations(globalPlayer, ennemi, Ordreatt)
		fmt.Println("Que voulez-vous faire ?")
		fmt.Println("1. Attaquer")
		fmt.Printf("2. Utiliser un soin (%d) \n", Nbsoin)
		fmt.Println("3. Fuir")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if Ordreatt {
				attaquer(globalPlayer, ennemi)
				if ennemi.Hp > 0 {
					ennemiAttaque(ennemi, globalPlayer)
				}
			} else {
				ennemiAttaque(ennemi, globalPlayer)
				attaquer(globalPlayer, ennemi)
			}
		case 2:
			if Nbsoin != 0 {
				success = utiliserSoin(globalPlayer, globalInventory, &Nbsoin) // Utilise un soin avec l'inventaire global
				if !success {
					fmt.Println("L'utilisation de la potion a échoué.")
				}
			} else {
				fmt.Println("Vous n'avez pas de potions de soin.")
			}
		case 3:
			fmt.Println("Vous avez tenté de fuir mais vos petites pattes sont vraiment trop courtes, donc vous êtes lents.")
			ennemiAttaque(ennemi, globalPlayer)
		default:
			fmt.Println("Choix invalide. Veuillez choisir une option valide.")
		}
		ClearScreen()
	}
}

func attaquer(attacker *Player, target *Mob) {
	if rand.Intn(20) >= 1 {
		//fmt.Printf("%s attaque et inflige %d points de dégâts à %s.\n", attacker.Name, damage, target.Name)
		target.Hp -= attacker.Attack
	}
}

func ennemiAttaque(attacker *Mob, target *Player) {
	if rand.Intn(20) >= 1 {
		//fmt.Printf("%s attaque et inflige %d points de dégâts à %s.\n", attacker.Name, attacker.Attack, target.Name)
		target.Hp -= attacker.Attack
	}
}

func utiliserSoin(joueur *Player, inv *Inventory, NbSoins *int) bool {
	if *NbSoins >= 0 {
		// Définir le soin (20 points de vie dans cet exemple)
		soin := 20
		if joueur.Hp+soin > joueur.HpMax {
			joueur.Hp = joueur.HpMax
		} else {
			// Appliquer l'effet de guérison au joueur
			joueur.Hp += soin
		}
		// Décrémenter le nombre de soins disponibles
		*NbSoins--

		fmt.Printf("%s utilise un soin et récupère %d points de vie. Points de vie actuels : %d/%d\n", joueur.Name, soin, joueur.Hp, joueur.HpMax)

		return true
	} else {
		fmt.Println("Vous n'avez plus de soins disponibles.")
		return false
	}
}

func afficherInformations(joueur *Player, ennemi *Mob, Ordreatt bool) {
	// Affiche les informations du joueur
	if Ordreatt {
		fmt.Printf(Green + "Vous etes le premier a attaquer \n" + Reset)
	} else {
		fmt.Printf(Red+"%s vous portera le premier coup \n"+Reset, ennemi.Name)
	}
	afficherBarreDeVie(joueur)
	afficherBarreDeVieEnnemi(ennemi)
}
func affichecombatlog(jmiss bool, emiss bool) {
	if jmiss {
		fmt.Printf("Vous avez rattez votre attaque")
	}
	if emiss {
		fmt.Printf("Votre adversaire a ratter son attaque")
	}
}

func afficherBarreDeVie(p *Player) {
	barre := ""
	pourcentage := float32(p.Hp) / float32(p.HpMax) * 100
	nombreDeBarres := int(pourcentage / 10)

	for i := 0; i < nombreDeBarres; i++ {
		barre += "#"
	}

	fmt.Printf(" Barre de vie: %s[%s]%s (%d/%d)\n", Red, barre, Reset, p.Hp, p.HpMax)
}
func afficherBarreDeVieEnnemi(ennemi *Mob) {
	barre := ""
	pourcentage := float32(ennemi.Hp) / float32(ennemi.HpMax) * 100
	nombreDeBarres := int(pourcentage / 10)

	for i := 0; i < nombreDeBarres; i++ {
		barre += "#"
	}
	fmt.Printf(" Barre de vie de l'ennemi: %s[%s]%s (%d/%d)\n", Red, barre, Reset, ennemi.Hp, ennemi.HpMax)
}
func Affichemob(ennemi *Mob) {
	//fmt.Println("\x1b[7m") //affichaage en blanc
	// Lire le contenu du fichier ascii.txt
	strWin11 := "../image/" + ennemi.Name + ".txt"
	asciiBytes, err := ioutil.ReadFile(strWin11)

	if err != nil {
		//fmt.Println("Erreur lors de la lecture du fichier :", err)
		strWin10 := "./image/" + ennemi.Name + ".txt"
		asciiBytes, erro := ioutil.ReadFile(strWin10)
		fmt.Print(erro)

		asciiArt := string(asciiBytes)
		fmt.Println(asciiArt)
	} else {

		asciiArt := string(asciiBytes)
		fmt.Println(asciiArt)
	}
	fmt.Println("\x1b[0m")
}
