package main

import (
	"encoding/json"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/riesawatani/montai/niku"
	"github.com/riesawatani/montai/qstrage"
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

type MyColor struct {
	R uint8
	G uint8
	B uint8
}
type QAP struct {
	Question  string
	Answer    string
	Color     MyColor
	Question2 string
	Hinto     string
	Question3 string
	Question4 string
}

type Game struct {
	Msg            string
	Count          int
	niku           niku.Niku
	keys           []ebiten.Key
	Questionlist   []QAP
	Questionnunvar uint
	seikaisita     bool
	IsMachi        bool
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.Count = g.Count + 1
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	haikei := color.RGBA{
		R: 250,
		G: 150,
		B: 100,
		A: 255,
	}
	iro := color.RGBA{
		R: 0,
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
	kotaenoiro := color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	k := len(g.Questionlist)
	if g.Questionnunvar == uint(k) {
		screen.Fill(haikei)
		text.Draw(screen, "おめでとうございます！！", mPlus1pRegular_ttf, 20, 120, seikainoiro)

		return
	}
	t := g.Questionlist[g.Questionnunvar]
	q := t.Question
	a := t.Answer
	q2 := t.Question2
	q3 := t.Question3
	q4 := t.Question4
	screen.Fill(color.RGBA{
		R: t.Color.R,
		G: t.Color.G,
		B: t.Color.B,
		A: 255,
	})
	text.Draw(screen, q, mPlus1pRegular_ttf, 0, 20, iro)
	text.Draw(screen, q2, mPlus1pRegular_ttf, 0, 40, iro)
	text.Draw(screen, q3, mPlus1pRegular_ttf, 0, 60, iro)
	text.Draw(screen, q4, mPlus1pRegular_ttf, 0, 80, iro)
	if len(g.keys) > 0 {
		st := strings.TrimPrefix(g.keys[0].String(), "Digit")
		text.Draw(screen, st, mPlus1pRegular_ttf, 0, 100, kotaenoiro)
		if st == a {
			g.Questionnunvar = g.Questionnunvar + 1
			g.seikaisita = true
		} else {
			text.Draw(screen, "不正解", mPlus1pRegular_ttf, 0, 130, fuseikainoiro)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	data, err := qstrage.ReadJson("cmd/upload/rie.json")
	if err != nil {
		log.Fatal(err)
	}
	var list []QAP
	if err := json.Unmarshal([]byte(data), &list); err != nil {
		log.Fatal(err)
	}

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
		niku:         ushi,
		Questionlist: list,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
