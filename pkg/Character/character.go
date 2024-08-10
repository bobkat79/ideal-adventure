package character

import (
	"errors"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img   *ebiten.Image
	X, Y  float64
	Frame *ebiten.Image
}

type Player struct {
	*Sprite
	Health int
	Speed  float64
}

type Enemy struct {
	*Sprite
	Health        int
	FollowsPlayer bool
}

func (p *Player) MoveLeft() {
	p.X -= p.Speed
}

func (p *Player) MoveRight() {
	p.X += p.Speed
}

func (p *Player) MoveUp() {
	p.Y -= p.Speed
}

func (p *Player) MoveDown() {
	p.Y += p.Speed
}

func SetupNewPlayer() Player {
	// Default loads the mage
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/NinjaMage.png")
	frame := playerImg.SubImage(
		image.Rect(0, 0, 16, 16),
	).(*ebiten.Image)
	if err != nil {
		log.Fatal("can't load player image")
	}
	plsprite := Sprite{
		Img:   playerImg,
		X:     20,
		Y:     20,
		Frame: frame,
	}
	player := Player{
		Sprite: &plsprite,
		Health: 100,
		Speed:  2,
	}
	return player

}

func InitializeEnemies() ([]Enemy, error) {
	// Setup two robots
	robotimg, _, err := ebitenutil.NewImageFromFile("assets/images/RobotG.png")
	frame := robotimg.SubImage(
		image.Rect(0, 0, 16, 16),
	).(*ebiten.Image)
	if err != nil {
		return nil, errors.New("failed to load robot image")
	}
	e := []Enemy{}
	r1s := Sprite{
		Img:   robotimg,
		X:     10,
		Y:     40,
		Frame: frame,
	}
	r2s := Sprite{
		Img:   robotimg,
		X:     30,
		Y:     40,
		Frame: frame,
	}
	r1 := Enemy{
		Sprite:        &r1s,
		Health:        5,
		FollowsPlayer: true,
	}
	e = append(e, r1)
	r2 := Enemy{
		Sprite:        &r2s,
		Health:        5,
		FollowsPlayer: false,
	}
	e = append(e, r2)
	return e, nil
}
