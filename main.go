package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	niku Niku
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.niku.Neme)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	game := &Game{
		niku: Niku{
			Neme: "FFFFFFFF",
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Niku struct {
	Neme string
	sizu int
}
