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

const (
	MAX_W_CELL = int(WINDOW_WIDTH / CELL_SIZE)
	MAX_H_CELL = int(WINDOW_HEIGHT / CELL_SIZE)
)

type Game struct {
	snake        Snake
	tickCount    int
	moveInterval int
	food         Food
	drawText     bool
	font_face    font.Face
	score        int
	gameOver     bool
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
		moveInterval: 7,
		food:         GenerateRandomFood(int(MAX_W_CELL), int(MAX_H_CELL)),
		drawText:     drawText,
		font_face:    face,
		score:        0,
		gameOver:     false,
	}
}

func (g *Game) Update() error {
	if !g.gameOver {
		g.tickCount++
		k_up := ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)
		k_down := ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown)
		k_left := ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
		k_right := ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight)
		has_only_head := len(g.snake.segments) == 1
		// Snake should not move in its opposite direction.
		// But a snake with only head could!
		if k_up && (g.snake.direction != DIR_DOWN || has_only_head) {
			g.snake.direction = DIR_UP
		}
		if k_down && (g.snake.direction != DIR_UP || has_only_head) {
			g.snake.direction = DIR_DOWN
		}
		if k_left && (g.snake.direction != DIR_RIGHT || has_only_head) {
			g.snake.direction = DIR_LEFT
		}
		if k_right && (g.snake.direction != DIR_LEFT || has_only_head) {
			g.snake.direction = DIR_RIGHT
		}
		if g.tickCount >= g.moveInterval {
			g.snake.Update()
			if len(g.snake.segments) > 1 {
				for _, seg := range g.snake.segments[1:] {
					if g.snake.segments[0] == seg {
						g.gameOver = true
					}
				}
			}
			head := g.snake.segments[0]
			if head.x < 0 || head.x >= MAX_W_CELL || head.y < 0 || head.y >= MAX_H_CELL {
				g.gameOver = true
			}

			//Eat food
			if g.food.Pos == g.snake.segments[0] {
				g.food = GenerateRandomFood(int(MAX_W_CELL), int(MAX_H_CELL))
				// last_segment := g.snake.segments[len(g.snake.segments)-1]
				// second_last_segment := Position{0, 0}
				// if len(g.snake.segments) > 1 {
				// 	second_last_segment = g.snake.segments[len(g.snake.segments)-2]
				// } else {
				// 	second_last_segment = last_segment
				// }
				// dx := last_segment.x - second_last_segment.x
				// dy := last_segment.y - second_last_segment.y
				// new_segment := Position{
				// 	x: last_segment.x + dx,
				// 	y: last_segment.y + dy,
				// }
				// g.snake.segments = append(g.snake.segments, new_segment)
				tail := g.snake.segments[len(g.snake.segments)-1]
				g.snake.segments = append(g.snake.segments, tail)
				g.score++
			}
			g.tickCount = 0
		}
	} else {

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	if !g.gameOver {
		g.food.Draw(screen)
		g.snake.Draw(screen)

		if g.drawText {
			text.Draw(screen, fmt.Sprintf("Score: %d", g.score), g.font_face, 30, 30, color.White)
		}
	} else {
		if g.drawText {
			width := font.MeasureString(g.font_face, "Game Over!").Round()
			text.Draw(screen, "Game Over!", g.font_face, int(WINDOW_WIDTH/2-width/2), int(WINDOW_HEIGHT/2), color.White)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
