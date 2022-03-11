package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/riesawatani/montai/niku"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mPlus1pRegular_ttf font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	ft, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mPlus1pRegular_ttf = ft
}

type Game struct {
	Msg   string
	Count int
	niku  niku.Niku
}

func (g *Game) Update() error {
	g.Count = g.Count + 1
	if g.Count < 60 {
		return nil
	}
	g.Count = 0
	if len(g.Msg) > 0 {
		g.Msg = ""
	} else {
		g.Msg = g.niku.Taberu()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x := 90
	y := 120 + g.Count

	iro := color.RGBA{
		R: 255,
		G: 0,
		B: uint8(4 * g.Count),
		A: 100,
	}

	text.Draw(screen, g.Msg, mPlus1pRegular_ttf, x, y, iro)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	buta := niku.Niku{
		Neme: "buta",
	}
	log.Println(buta)

	ushi := niku.Niku{
		Neme: "ushi",
	}
	log.Println(ushi)

	game := &Game{
		niku: ushi,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
