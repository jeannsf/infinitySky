package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player *Player
}

func NewGame() *Game {
	player := NewPlayer()
	return &Game{
		Player: player,
	}
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
