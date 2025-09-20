/*
Conatins 'main' function with logic to iterate and print/display drops
*/
package main

import (
	"math/rand"
	"time"

	// use tcell for TUI
	"github.com/gdamore/tcell/v2"
)

// represents a drop character trail
type drop struct {
	x      int
	y      int
	speed  int
	length int // trail length
}

func main() {
	// init new tcell.Screen object
	screen := initScreen()
	defer screen.Fini()

	// initalise list of drops
	// drops should span the screen width
	width, height := screen.Size()
	drops := make([]drop, width)

	// initialise the drops
	for i := range drops {
		drops[i] = drop{
			x:      i,
			y:      getDropHeight(height),
			speed:  getDropSpeed(),
			length: rand.Intn(10) + 5,
		}
	}

	// color shades
	colors := initColorShades()

	// event channel for screen key press
	eventChan := make(chan tcell.Event, 10)
	// poll keypress event concurrently
	// go-routine
	go func() {
		for {
			ev := screen.PollEvent()
			eventChan <- ev
		}
	}()

	// make background black
	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
	screen.Clear()
	clearupScreen(width, height, screen)
	screen.Show()

	// loop indefinitely
	for {
		// check channel for key press event
		select {
		case ev := <-eventChan:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					//  pause a bit 10 X loop delay
					time.Sleep(10 * delay())
					// .. and exit main
					return
				}
			}
		default:
			// continue
		}

		// if not esc|ctrl+c key pressed - carry on..

		// reset to black for erasing pervious drop trail
		clearupScreen(width, height, screen)

		// iterate through the drop trails
		// ..across the screen width
		for i := range drops {
			d := &drops[i]

			// for each drop trail, print
			// .. the character at (x,y) position
			for j := 0; j < d.length; j++ {
				y := d.y - j
				if y < 0 || y >= height {
					continue
				}

				// rune/char to display
				char := getChar(i, j)

				// prepare font, color style
				// .. for displaying the character
				style := getStyle(i, j, colors)

				// print the drop character
				screen.SetContent(d.x, y, char, nil, style)
			}
			// move drop position down by its speed step
			d.y += d.speed

			// reset position if reached bottom
			if d.y-d.length >= height {
				d.y = getDropHeight(height)
			}
		}
		screen.Show()
		// delay between iterations
		time.Sleep(delay())
	}

}
