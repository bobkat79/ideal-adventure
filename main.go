package main

import (
	"image"
	"image/color"
	"log"

	camera "github.com/bobkat79/ideal-adventure/Camera"
	character "github.com/bobkat79/ideal-adventure/Character"
	mapping "github.com/bobkat79/ideal-adventure/Mapping"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player    character.Player
	Enemies   []character.Enemy
	Overworld *mapping.TileMap
	Camera    *camera.Camera
}

func (g *Game) Update() error {

	// react to keypresses
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Player.MoveRight()
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Player.MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Player.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Player.MoveDown()
	}
	g.Camera.FollowTarget(g.Player.X+8, g.Player.Y+8)
	g.Camera.Constrain(100.0*16.0, 80.0*16.0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}

	for _, layer := range g.Overworld.TMJ.Layers {
		// Loop over the layers
		for i, id := range layer.Data {
			x, y := g.Overworld.GetMapPos(i, layer.Width)

			// Set the draw image options
			opts.GeoM.Translate(float64(x), float64(y))

			opts.GeoM.Translate(g.Camera.X, g.Camera.Y)
			a, b, c, d := g.Overworld.TMImageTranslate(id)
			screen.DrawImage(
				g.Overworld.IMG.SubImage(image.Rect(a, b, c, d)).(*ebiten.Image),
				&opts,
			)
			opts.GeoM.Reset()
		}

	}

	opts.GeoM.Translate(g.Player.X, g.Player.Y)
	opts.GeoM.Translate(g.Camera.X, g.Camera.Y)
	// Draw our player
	screen.DrawImage(
		g.Player.Frame,
		&opts,
	)
	opts.GeoM.Reset()

	for _, sprite := range g.Enemies {
		opts.GeoM.Translate(sprite.X, sprite.Y)
		screen.DrawImage(
			sprite.Frame,
			&opts,
		)
		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	player := character.SetupNewPlayer()
	enemies, err := character.InitializeEnemies()
	if err != nil {
		log.Fatal(err)
	}
	ow, err := mapping.LoadOverworldMap()
	if err != nil {
		log.Fatal("Failed to load overworld map.")
	}

	c := camera.NewCamera(0, 0, 320, 240)
	game := Game{
		Player:    player,
		Enemies:   enemies,
		Overworld: ow,
		Camera:    c,
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
