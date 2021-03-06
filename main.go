package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	CurrentFrame int
	Ticks        int
	Sprites      []*ebiten.Image
}

func (g *Game) Update() error {
	g.Ticks++
	if g.Ticks >= 5 {
		g.Ticks = 0
		g.CurrentFrame++
		if g.CurrentFrame >= 99 {
			g.CurrentFrame = 0
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x, y := ebiten.WindowPosition()
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			"Ticks: %d\nCurrentFrame: %d\nx: %v, y: %v",
			g.Ticks, g.CurrentFrame, x, y,
		),
	)
	screen.DrawImage(g.Sprites[g.CurrentFrame], nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (w, h int) {
	return outsideWidth, outsideHeight
}

func PanicIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var game Game

	// Should probably use go:embed somehow here
	for i := 1; i <= 100; i++ {
		img, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("sprites/f%03d.png", i))
		PanicIfErr(err)
		game.Sprites = append(game.Sprites, img)
	}

	ebiten.SetWindowSize(350, 450)
	ebiten.SetWindowTitle("Ahoy!")
	ebiten.SetWindowDecorated(false)
	ebiten.SetScreenTransparent(true)
	ebiten.SetWindowPosition(9999, 9999)
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
