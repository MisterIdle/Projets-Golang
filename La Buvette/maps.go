package rpg

type Teleporter struct {
	X, Y           int
	DestinationX   int
	DestinationY   int
	DestinationMap *Map
	Appearance     rune
}

type Shop struct {
	X, Y       int
	Appearance rune
	num        int
}

type Map struct {
	Name          string // Ajoutez un champ pour le nom de la carte
	Width, Height int
	Tiles         [][]rune
	Teleporters   []Teleporter
	Shop          []Shop
	Mobs          []Mob
}

func NewMap(name string, width, height int) *Map {
	mapInstance := &Map{
		Name:   name,
		Width:  width,
		Height: height,
		Tiles:  InitializeMap(width, height),
	}
	return mapInstance
}

func InitializeMap(width, height int) [][]rune {
	tiles := make([][]rune, height)
	for y := 0; y < height; y++ {
		tiles[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			if y == 0 || y == height-1 || x == 0 || x == width-1 {
				tiles[y][x] = '#' // Contour
			} else {
				tiles[y][x] = ' ' // Espace à l'intérieur du contour
			}
		}
	}
	return tiles
}

func MakeAndPlaceTeleporter(currentMap *Map, TeleporterX, TeleporterY int, TeleportX, TeleportY int, DestinationMap *Map, Appearance rune) {
	teleporter := Teleporter{
		X:              TeleporterX,
		Y:              TeleporterY,
		DestinationX:   TeleportX,
		DestinationY:   TeleportY,
		DestinationMap: DestinationMap,
		Appearance:     Appearance,
	}
	currentMap.Teleporters = append(currentMap.Teleporters, teleporter)
}

func MakeAndPlaceShop(currentMap *Map, ShopX, ShopY int, Appearance rune, num int) {
	shop := Shop{
		X:          ShopX,
		Y:          ShopY,
		Appearance: Appearance,
		num:        num,
	}
	currentMap.Shop = append(currentMap.Shop, shop)
}
