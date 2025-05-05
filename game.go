package main

import (
	"image/color"
	"log"
	"os"

	"fmt"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Game struct {
	snake        Snake
	tickCount    int
	moveInterval int
	food         Food
	drawText     bool
	font_face    font.Face
	score        int
}

func NewGame() Game {
	drawText := true
	data, err := os.ReadFile("ARIAL.TTF")
	if err != nil {
		log.Printf("Could not find ARIAL.TTF font\n")
		drawText = false
	}

	font, err := truetype.Parse(data)
	if err != nil {
		log.Printf("Could not parse font content\n")
		drawText = false
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 19})
	return Game{
		snake: Snake{
			segments:  []Position{{3, 3}},
			direction: DIR_RIGHT,
		},
		moveInterval: 10,
		food:         GenerateRandomFood(int(WINDOW_WIDTH/CELL_SIZE), int(WINDOW_HEIGHT/CELL_SIZE)),
		drawText:     drawText,
		font_face:    face,
		score:        0,
	}
}

func (g *Game) Update() error {
	g.tickCount++
	if g.tickCount >= g.moveInterval {
		if g.food.Pos == g.snake.segments[0] {
			g.food = GenerateRandomFood(int(WINDOW_WIDTH/CELL_SIZE), int(WINDOW_HEIGHT/CELL_SIZE))
			last_segment := g.snake.segments[len(g.snake.segments)-1]
			second_last_segment := Position{0, 0}
			if len(g.snake.segments) > 1 {
				second_last_segment = g.snake.segments[len(g.snake.segments)-2]
			} else {
				second_last_segment = last_segment
			}

			dx := last_segment.x - second_last_segment.x
			dy := last_segment.y - second_last_segment.y
			new_segment := Position{
				x: last_segment.x + dx,
				y: last_segment.y + dy,
			}
			g.snake.segments = append(g.snake.segments, new_segment)
			g.score++
		}
		g.snake.Update()
		g.tickCount = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.food.Draw(screen)
	g.snake.Draw(screen)

	if g.drawText {
		text.Draw(screen, "Score: "+fmt.Sprint(g.score), g.font_face, 30, 30, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
