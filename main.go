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
const DIFFICULT_UP_PERSENT = 10
const DIFFICULT_UP_AFTER_TURNS = 20
var speed = 250

type Point struct {
	x int
	y int
}

var vector = Point{0, -1}
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
	t := 0
	for {
		t += 1
		if t % DIFFICULT_UP_AFTER_TURNS == 0 {
			speed -= int(float64(speed) / 100 * DIFFICULT_UP_PERSENT)
			addPart()
		}

		Move()

		out := ""
		offsetX := ""
		offsetY := ""

		w, h := termbox.Size()
		for i := 0; i < (w-WIDTH)/2; i += 1 {
			offsetX += " "
		}
		for i := 0; i < (h-HEIGHT)/2; i += 1 {
			offsetY += "\n"
		}

		out += offsetY + offsetX + "+"
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
						s = color.RedString("@")
					}
				}
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
		fmt.Print(speed)
		termbox.Flush()

		time.Sleep(time.Millisecond * time.Duration(speed))
	}
}

func Move() {
	newX := snake[0].x + vector.x
	newY := snake[0].y + vector.y

	if newX < 0 || newX >= WIDTH || newY < 0 || newY >= HEIGHT {
		color.Red("Game ower")
		os.Exit(0)
	} else {
		if len(snake) >= 2 {
			for i := len(snake) - 2; i >= 0; i -= 1 {
				snake[i + 1].x = snake[i].x
				snake[i + 1].y = snake[i].y
			}
		}
		snake[0].x = newX
		snake[0].y = newY
	}
}

func addPart() {
	last := snake[len(snake) - 1]
	snake = append(snake, &Point{last.x, last.y})
}

func goDown() {
	vector = Point{0, 1}
}

func goUp() {
	vector = Point{0, -1}
}

func goLeft() {
	vector = Point{-1, 0}
}

func goRight() {
	vector = Point{1, 0}
}
