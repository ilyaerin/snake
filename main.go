package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/nsf/termbox-go"
)

const WIDTH = 30
const HEIGHT = 30

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
			if ev.Key == termbox.KeyCtrlQ || ev.Key == termbox.KeyCtrlD {
				break loop
			}
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			moveHead()
			draw()
			termbox.Flush()
		case termbox.EventResize:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			draw()
			fmt.Println(termbox.Size())
			termbox.Flush()
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func draw() {
	out := "\n\n\n\n\n\n\n\n\n\n"
	for i := 0; i < WIDTH; i +=1 {
		for j := 0; j < HEIGHT; j +=1 {
			s := "."
			if head.x == i && head.y == j { s = "*" }
			out += s
		}
		out += "\n"
	}
	fmt.Println(out)
}

func moveHead()  {
	head.x = rand.Intn(WIDTH)
	head.y = rand.Intn(HEIGHT)
}