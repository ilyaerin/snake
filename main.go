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
var speed = 250

type Point struct {
	x int
	y int
}

var direction = Point{0, -1}
var head = Point{WIDTH/2, HEIGHT/2}

func main() {
	rand.Seed(time.Now().Unix())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

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
	ticker := time.NewTicker(time.Millisecond * time.Duration(speed))
	for range ticker.C {
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
				if head.x == i && head.y == j {
					s = color.RedString("@")
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
		termbox.Flush()
	}
}

func Move() {
	newX := head.x + direction.x
	newY := head.y + direction.y

	if newX < 0 || newX >= WIDTH || newY < 0 || newY >= HEIGHT {
		color.Red("Game ower")
		os.Exit(0)
	} else {
		head.x = newX
		head.y = newY
	}
}

func goDown() {
	direction = Point{0, 1}
}

func goUp() {
	direction = Point{0, -1}
}

func goLeft() {
	direction = Point{-1, 0}
}

func goRight() {
	direction = Point{1, 0}
}
