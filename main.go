package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const WINDOW_WIDTH = 640
const WINDOW_HEIGHT = 480

func main() {
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Snake")

	game := NewGame()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
