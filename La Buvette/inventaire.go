package rpg

import (
	"fmt"
	"strings"
	"sync"
)

// Structure pour représenter un objet dans l'inventaire
type Item struct {
	Name        string
	Description string
	Quantity    int
	Stats       int
	Typestats   int
}

// Structure pour représenter l'inventaire
type Inventory struct {
	Items []*Item
}

var (
	inventory     *Inventory
	inventoryOnce sync.Once
)

// Fonction pour créer un nouvel inventaire vide
func NewInventory() *Inventory {
	inventoryOnce.Do(func() {
		inventory = &Inventory{
			Items: make([]*Item, 0),
		}
	})
	return inventory
}

// Fonction pour ajouter un objet à l'inventaire
func (inv *Inventory) AddItem(item *Item) {
	// Vérifie si l'objet est déjà dans l'inventaire
	for _, existingItem := range inv.Items {
		if existingItem.Name == item.Name {
			existingItem.Quantity += item.Quantity
			fmt.Printf("%s a été ajouté à l'inventaire (x%d).\n", item.Name, item.Quantity)
			return
		}
	}

	// Si l'objet n'est pas déjà dans l'inventaire, ajoutez-le
	inv.Items = append(inv.Items, item)
	fmt.Printf("%s a été ajouté à l'inventaire.\n", item.Name)
}

// Fonction pour retirer un objet de l'inventaire
func (inv *Inventory) RemoveItem(itemName string, quantity int) {
	for i, item := range inv.Items {
		if item.Name == itemName {
			if item.Quantity > quantity {
				item.Quantity -= quantity
				fmt.Printf("%dx %s a été retiré de l'inventaire.\n", quantity, itemName)
			} else if item.Quantity == quantity {
				// Supprime complètement l'objet de l'inventaire
				inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
				fmt.Printf("%s a été retiré de l'inventaire.\n", itemName)
			} else {
				fmt.Printf("L'inventaire ne contient pas suffisamment de %s.\n", itemName)
			}
			return
		}
	}
	fmt.Printf("L'inventaire ne contient pas de %s.\n", itemName)
}

// Fonction pour afficher l'inventaire
func (inv *Inventory) ShowInventory(player *Player) {
	for {
		if len(inv.Items) == 0 {
			fmt.Println("L'inventaire est vide.")
			break
		}

		ClearScreen()
		fmt.Println("Inventaire:")
		for _, item := range inv.Items {
			fmt.Printf("%s (x%d) - %s\n", item.Name, item.Quantity, item.Description)
		}

		fmt.Printf("Position (X, Y): (%d, %d)\n", player.X, player.Y)
		fmt.Printf("Carte actuelle: %s\n", player.CurrentMap.Name)
		fmt.Printf("Nom du joueur: %s\n", player.Name)
		fmt.Printf("Points de vie: %d/%d\n", player.Hp, player.HpMax)
		fmt.Printf("Attaque: %d\n", player.Attack)

		fmt.Print("Appuyez sur 'quit', 'Q', ou 'q' pour sortir de l'inventaire: ")
		var input string
		fmt.Scanln(&input)

		input = strings.ToLower(input)
		if input == "quit" || input == "q" {
			break
		}
	}
}

func (inv *Inventory) AddItemFromShop(itemShop *ItemShop) {
	// Créez un nouvel élément d'inventaire basé sur l'article du magasin
	inventoryItem := &Item{
		Name:        itemShop.Name,
		Description: fmt.Sprintf("Armure: %d, Attaque: %d, Soins: %d", itemShop.Effectarmure, itemShop.Effectattack, itemShop.Effectheal),
		Quantity:    1, // L'article est ajouté avec une quantité de 1
		Stats:       0, // Vous pouvez initialiser les statistiques à zéro, car elles seront gérées différemment dans votre jeu
		Typestats:   0, // Vous pouvez initialiser le type de statistiques à zéro également
	}

	// Ajoutez l'article à l'inventaire
	inv.AddItem(inventoryItem)
}
