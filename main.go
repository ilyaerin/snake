package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
)

const WIDTH = 30
const HEIGHT = 20

type Point struct {
	x int
	y int
}

var head = Point{rand.Intn(WIDTH), rand.Intn(HEIGHT)}

func main() {
	rand.Seed(time.Now().Unix())

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	draw()

	loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ || ev.Key == termbox.KeyCtrlD { break loop }
			if ev.Key == termbox.KeyArrowRight { goRight() }
			if ev.Key == termbox.KeyArrowLeft { goLeft() }
			if ev.Key == termbox.KeyArrowDown { goDown() }
			if ev.Key == termbox.KeyArrowUp { goUp() }

			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw()
			termbox.Flush()

		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw()
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func draw() {
	out := ""
	offsetX := ""
	offsetY := ""

	w, h := termbox.Size()
	for i := 0; i < (w - WIDTH) / 2; i +=1 { offsetX += " " }
	for i := 0; i < (h - HEIGHT) / 2 - 1; i +=1 { offsetY += "\n" }

	out += offsetY + offsetX + "+"
	for i := 0; i < WIDTH; i +=1 { out += "-" }
	out += "+\n"

	for j := 0; j < HEIGHT; j +=1 {
		out += offsetX + "|"
		for i := 0; i < WIDTH; i +=1 {
			s := " "
			if head.x == i && head.y == j {
				s = color.RedString("@")
			}
			out += s
		}
		out += "|\n"
	}

	out += offsetX + "+"
	for i := 0; i < WIDTH; i +=1 { out += "-" }
	out += "+" + offsetY

	fmt.Print(out)
}

func goDown() {
	if head.y < HEIGHT - 1 {
		head.y += 1
	}
}

func goUp() {
	if head.y >= 1 {
		head.y -= 1
	}
}

func goLeft() {
	if head.x >= 1 {
		head.x -= 1
	}
}

func goRight() {
	if head.x < WIDTH - 1 {
		head.x += 1
	}
}
