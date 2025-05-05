package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Food struct {
	Pos   Position
	Eaten bool
}

func (f *Food) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(CELL_SIZE*f.Pos.x), float32(CELL_SIZE*f.Pos.y), float32(CELL_SIZE), float32(CELL_SIZE), color.RGBA{255, 0, 0, 255}, false)
}

func GenerateRandomFood(w_max, h_max int) Food {
	x := rand.Intn(w_max)
	y := rand.Intn(h_max)
	return Food{Position{x, y}, false}
}
