package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
	"os"
)

const WIDTH = 30
const HEIGHT = 20
const DIFFICULT_UP_PERSENT = 5
const MINUS_PERCENT = 2
const PLUS_PERCENT = 10
var speed = 250
var score = 0

type Point struct {
	x int
	y int
}

var vector = Point{0, -1}
var plus = Point{-1, -1}
var minus = Point{-1, -1}
var snake = []*Point{}

func main() {
	rand.Seed(time.Now().Unix())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	snake = append(snake, &Point{WIDTH/2, HEIGHT/2})

	go draw()

	loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ || ev.Key == termbox.KeyCtrlD { break loop }
			if ev.Key == termbox.KeyArrowRight { goRight() }
			if ev.Key == termbox.KeyArrowLeft { goLeft() }
			if ev.Key == termbox.KeyArrowDown { goDown() }
			if ev.Key == termbox.KeyArrowUp { goUp() }
		//case termbox.EventResize:
		//	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		//	draw()
		//	termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func draw() {
	for {
		gameOver := Move()
		addPlus()
		addMinus()

		offsetX := ""
		offsetY := ""

		w, h := termbox.Size()
		for i := 0; i < (w-WIDTH)/2; i += 1 {
			offsetX += " "
		}
		for i := 0; i < (h-HEIGHT)/2; i += 1 {
			offsetY += "\n"
		}


		out := fmt.Sprintf("%s%s Score: %d", offsetY, offsetX, score)

		if gameOver { out += color.RedString("    Game over!") }

		out += "\n" + offsetX + "+"
		for i := 0; i < WIDTH; i += 1 {
			out += "-"
		}
		out += "+\n"

		for j := 0; j < HEIGHT; j += 1 {
			out += offsetX + "|"
			for i := 0; i < WIDTH; i += 1 {
				s := " "
				for _, part := range snake {
					if part.x == i && part.y == j {
						s = color.YellowString("@")
					}
				}
				if plus.x == i && plus.y == j { s = color.GreenString("@") }
				if minus.x == i && minus.y == j { s = color.RedString("X") }
				out += s
			}
			out += "|\n"
		}

		out += offsetX + "+"
		for i := 0; i < WIDTH; i += 1 {
			out += "-"
		}
		out += "+" + offsetY

		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		fmt.Print(out)
		termbox.Flush()
		if gameOver { os.Exit(0) }

		time.Sleep(time.Millisecond * time.Duration(speed))
	}
}

func Move() bool {
	newX := snake[0].x + vector.x
	newY := snake[0].y + vector.y

	if newX < 0 || newX >= WIDTH || newY < 0 || newY >= HEIGHT {
		return true
	} else {
		if len(snake) >= 2 {
			for i := len(snake) - 2; i >= 0; i -= 1 {
				snake[i + 1].x = snake[i].x
				snake[i + 1].y = snake[i].y
			}
		}
		snake[0].x = newX
		snake[0].y = newY

		if newX == plus.x && newY == plus.y { addPart() }
		if newX == minus.x && newY == minus.y { removePart() }
		return false
	}
}

func addPlus() {
	if rand.Intn(100) <= PLUS_PERCENT && plus.x == -1 {
		plus = Point{rand.Intn(WIDTH), rand.Intn(HEIGHT)}
	}
}

func addMinus() {
	if rand.Intn(100) <= MINUS_PERCENT && minus.x == -1 {
		minus = Point{rand.Intn(WIDTH + 1), rand.Intn(HEIGHT + 1)}
	}
}

func addPart() {
	last := snake[len(snake) - 1]
	plus = Point{-1, -1}
	snake = append(snake, &Point{last.x, last.y})
	speed -= int(float64(speed) / 100 * DIFFICULT_UP_PERSENT)
	score += 1
}

func removePart() {
	snake = snake[:len(snake) - 1]
	minus = Point{-1, -1}
	score -= 1
}

func goDown() {
	if vector.y == 0 { vector = Point{0, 1} }
}

func goUp() {
	if vector.y == 0 { vector = Point{0, -1} }
}

func goLeft() {
	if vector.x == 0 { vector = Point{-1, 0} }
}

func goRight() {
	if vector.x == 0 { vector = Point{1, 0} }
}
