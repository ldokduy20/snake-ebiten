package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Direction = int

const (
	DIR_UP Direction = iota
	DIR_DOWN
	DIR_LEFT
	DIR_RIGHT
)

const CELL_SIZE int = 20

type Position struct {
	x int
	y int
}

type Snake struct {
	segments  []Position
	direction Direction
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for i := range s.segments {
		vector.DrawFilledRect(screen, float32(CELL_SIZE*s.segments[i].x), float32(CELL_SIZE*s.segments[i].y), float32(CELL_SIZE), float32(CELL_SIZE), color.White, false)
	}
}

func (s *Snake) Update() {
	// Shift every segment to its previous element
	for i := len(s.segments) - 1; i > 0; i-- {
		s.segments[i] = s.segments[i-1]
	}

	// Increment or decrement head position according to its direction
	switch s.direction {
	case DIR_RIGHT:
		s.segments[0].x++
	case DIR_LEFT:
		s.segments[0].x--
	case DIR_UP:
		s.segments[0].y--
	case DIR_DOWN:
		s.segments[0].y++
	}
}
