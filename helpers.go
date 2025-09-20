/*
Conatins helper functions called from the main function
*/
package main

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

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
	// trail and error to tune
	// .. height for best effect
	return rand.Intn(height / 3)
}

// speed of each drop trail
func getDropSpeed() int {
	return rand.Intn(3) + 1
}

// varying color shades for drop trail
func initColorShades() []tcell.Color {
	return []tcell.Color{
		tcell.ColorGreen,
		tcell.ColorGreen,
		tcell.ColorDarkGreen,
		tcell.ColorDarkGreen,
		tcell.ColorDarkGreen,
		tcell.ColorOlive,
		tcell.ColorOlive,
		tcell.ColorForestGreen,
		tcell.ColorForestGreen,
		tcell.ColorForestGreen,
		tcell.ColorSeaGreen,
		tcell.ColorSeaGreen,
		tcell.ColorSeaGreen,
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

// to print white rabbit
const rabbit = '\U0001F407'

// character/rune that needs to be printed
func getChar(i, j int) rune {
	// for now just return a random
	//.. we can modify for more advanced later
	val := rand.Intn(94) + 33
	if (val > 33) && (val < 36) {
		// simple logic for occassional
		// .. rabbit rune
		val = rabbit
	}
	return rune(val) // ASCII 33–126
}

// decide the colour font style
// .. to dispaly the character
func getStyle(i, j int, colors []tcell.Color) tcell.Style {
	var style tcell.Style
	if j == 0 {
		// Head character with glow (bright green on black)
		style = tcell.StyleDefault.Foreground(tcell.ColorLime).Background(tcell.ColorBlack).Bold(true)
	} else {
		style = tcell.StyleDefault.Foreground(colors[min(j, len(colors)-1)]).Background(tcell.ColorBlack)
	}
	return style
}

// time delay between iteration
// .. in milliseconds
func delay() time.Duration {
	return 80 * time.Millisecond
}
