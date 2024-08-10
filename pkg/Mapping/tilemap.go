package mapping

import (
	"encoding/json"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MapLayerJSON struct {
	Data   []int `js:"data"`
	Width  int   `js:"width"`
	Height int   `js:"height"`
}

type TileMapJSON struct {
	Layers []MapLayerJSON `js:"layers"`
}

type TileMap struct {
	TMJ        TileMapJSON
	IMG        *ebiten.Image
	TileSize   int
	IMGMapSize int // amount of tiles per row in the image
}

func (tm *TileMap) GetMapPos(location int, mapwidth int) (mapx int, mapy int) {
	// calculates a map position based on Tile size and map width

	mapx = location % mapwidth
	mapy = location / mapwidth

	// Convert tile position to pixel position
	mapx *= tm.TileSize
	mapy *= tm.TileSize
	return mapx, mapy
}

func (tm *TileMap) TMImageTranslate(id int) (int, int, int, int) {
	// gets the coordinates on the TileSet image to pass to the rect function
	// Get the position on the image where the tile id is
	srcX := (id - 1) % tm.IMGMapSize
	srcY := (id - 1) / tm.IMGMapSize

	// Convert source position to pixel position
	srcX *= tm.TileSize
	srcY *= tm.TileSize
	dstX := srcX + tm.TileSize
	dstY := srcY + tm.TileSize
	return srcX, srcY, dstX, dstY
}

func LoadOverworldMap() (*TileMap, error) {
	fileh, err := os.ReadFile("assets/maps/overworld-floor-a.json")
	if err != nil {
		return nil, err
	}
	var overworldJSON TileMapJSON
	err = json.Unmarshal(fileh, &overworldJSON)
	if err != nil {
		return nil, err
	}
	OWImg, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		return nil, err
	}
	ow := TileMap{
		TMJ:        overworldJSON,
		IMG:        OWImg,
		TileSize:   16,
		IMGMapSize: 22,
	}
	return &ow, nil
}
