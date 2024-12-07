package rpg

import "fmt"

func Header(m Map) {
	fmt.Println(" ")
	fmt.Printf(" Position: [X: %d, Y: %d]\n", player.X, player.Y)
	afficherBarreDeVie(globalPlayer)
	fmt.Println("―――――――――――――――――――――――――――――――")
	fmt.Println(" Monstres:", CountMobsOnMap(player.CurrentMap))
	fmt.Println(" Or:", gold)
	fmt.Println("―――――――――――――――――――――――――――――――")
	fmt.Println(" Étage:", m.Name)
	fmt.Println("―――――――――――――――――――――――――――――――")
}

func Footer() {
	fmt.Println("―――――――――――――――――――――――――――――――")
	fmt.Println(" Légende:")
	fmt.Println("- @ (Vous)")
	fmt.Println("- # (Bordures)")
	fmt.Println("- ⍐ (Téléporteur)")
	fmt.Println("- ⌂ (Boutique)")
	fmt.Println("- s/S/g/G/b/B/l/L (Monstres)")
	fmt.Println("―――――――――――――――――――――――――――――――")
	fmt.Println(" Touches:")
	fmt.Println("- Z (↑)")
	fmt.Println("- Q (←)")
	fmt.Println("- S (↓)")
	fmt.Println("- D (→)")
	fmt.Println("- P (Pause)")
	fmt.Println("- L (Quitter)")
	fmt.Println("- Inv (Inventaire)")
	fmt.Println(" ")
	fmt.Print("Taper la touche puis appuyer sur Entrée: ")
}

func (m Map) PrintMap(player Player) {
	Header(m)
	for y := 0; y < m.Height; y++ {
		for i := 0; i < 10; i++ {
			fmt.Print(" ")
		}

		for x := 0; x < m.Width; x++ {
			if x == player.X && y == player.Y {
				fmt.Print(Blackground + "@" + Reset)
			} else {
				var tile rune
				for _, teleporter := range m.Teleporters {
					if teleporter.X == x && teleporter.Y == y {
						tile = teleporter.Appearance
						break
					}
				}

				for _, shop := range m.Shop {
					if shop.X == x && shop.Y == y {
						tile = shop.Appearance
						break
					}
				}

				if tile == 0 {
					for _, mob := range m.Mobs {
						if mob.X == x && mob.Y == y {
							tile = mob.Appearance
							break
						}
					}
				}

				if tile == 0 {
					tile = m.Tiles[y][x]
				}

				switch tile {

				//Carte
				case '#':
					fmt.Print(Gray + "#" + Reset) // Bordures maps
				case 'T':
					fmt.Print(Green + "⍐" + Reset) // Téléporteur
				case 't':
					fmt.Print(Red + "⍗" + Reset) // Téléporteur
				case 'M':
					fmt.Print(Bold + Yellow + "⌂" + Reset) // Boutique

				//Mobs
				case 's':
					fmt.Print(LightGreen + "s" + Reset)
				case 'S':
					fmt.Print(Green + "S" + Reset)
				case 'g':
					fmt.Print(Bold + LightGreen + "g" + Reset)
				case 'G':
					fmt.Print(Bold + Green + "G" + Reset)
				case 'b':
					fmt.Print(Yellow + "b" + Reset)
				case 'B':
					fmt.Print(Orange + "B" + Reset)
				case 'l':
					fmt.Print(Gray + "l" + Reset)
				case 'L':
					fmt.Print(Red + "L" + Reset)

				default:
					fmt.Print(" ") // Espace vide
				}
			}
		}
		fmt.Println()
	}
	Footer()
}
