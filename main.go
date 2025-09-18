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
		select {
		case ev := <-eventChan:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					return
				}
			}
		default:
			// continue
		}
		// reset to black for erasing pervious drop trail
		clearupScreen(width, height, screen)

		for i := range drops {
			d := &drops[i]

			for j := 0; j < d.length; j++ {
				y := d.y - j
				if y < 0 || y >= height {
					continue
				}
				char := rune(rand.Intn(94) + 33) // ASCII 33–126
				var style tcell.Style
				if j == 0 {
					// Head character with glow (bright green on black)
					style = tcell.StyleDefault.Foreground(tcell.ColorLime).Background(tcell.ColorBlack).Bold(true)
				} else {
					style = tcell.StyleDefault.Foreground(colors[min(j, len(colors)-1)]).Background(tcell.ColorBlack)
				}
				screen.SetContent(d.x, y, char, nil, style)
			}
			d.y += d.speed

			// reset position if reached bottom
			if d.y-d.length >= height {
				d.y = getDropHeight(height)
			}
		}
		screen.Show()
		time.Sleep(70 * time.Millisecond)
	}

}

// new screen object
func initScreen() tcell.Screen {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	return screen
}

// height drop starts falling from
func getDropHeight(height int) int {
	return rand.Intn(height / 2)
}

// speed of each drop trail
func getDropSpeed() int {
	return rand.Intn(3) + 1
}

// varying color shades for drop trail
func initColorShades() []tcell.Color {
	return []tcell.Color{
		tcell.ColorGreen,
		tcell.ColorDarkGreen,
		tcell.ColorDarkGreen,
		tcell.ColorOlive,
		tcell.ColorOlive,
		tcell.ColorForestGreen,
		tcell.ColorForestGreen,
		tcell.ColorBlack,
	}
}

// reset-clear screen
func clearupScreen(width, height int, screen tcell.Screen) {
	black := tcell.StyleDefault.Background(tcell.ColorBlack)
	for x := range width {
		for y := range height {
			screen.SetContent(x, y, ' ', nil, black)
		}
	}
}
