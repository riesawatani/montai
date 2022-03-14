package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	keys  []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
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
	haikei := color.RGBA{
		R: 0,
		G: 150,
		B: 150,
		A: 255,
	}
	screen.Fill(haikei)
	iro := color.RGBA{
		R: 1,
		G: 0,
		B: 0,
		A: 255,
	}
	seikainoiro := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	fuseikainoiro := color.RGBA{
		R: 0,
		G: 0,
		B: 225,
		A: 255,
	}
	text.Draw(screen, "3 + 5 = ", mPlus1pRegular_ttf, 0, 20, iro)
	if len(g.keys) > 0 {
		answer := g.keys[0].String()
		text.Draw(screen, answer, mPlus1pRegular_ttf, 90, 20, iro)
		if answer == "Digit8" {
			text.Draw(screen, "正解", mPlus1pRegular_ttf, 0, 20*2, seikainoiro)

		} else {
			text.Draw(screen, "不正解", mPlus1pRegular_ttf, 0, 20*2, fuseikainoiro)

		}
	}

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
