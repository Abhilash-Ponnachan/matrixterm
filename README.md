# 🔳 Matrix-Style Terminal Drip (in `Go`)

A simple `Go` program to display an effect of green characters that drip endlessly down the screen as in the iconic _Matrix_ movie scene.
It is a purely hobby project and intended purely for the fun of it.
Don't forget to look out the _white rabbit_!
---

## ✨ Features
- Matrix-style falling characters in the terminal
- Helper functions to play around with varying speeds, style and characters
- Written in pure Go — no external dependencies (except [tcell](https://github.com/gdamore/tcell) for terminal rendering)

---

### 📋 Prerequisites
- [Go](https://go.dev/dl/),I developed & tested on `go1.24`
- A terminal that supports ANSI colors

---

### 💻 Code Description
- The main logic is all in the `main()` function
- It initializes a slice of 'drop' `struct`, each representing a 'drop trail' down the screen, and the collection spanning the width of the screen
- The drop trail has properties for teh height it starts from, the speed(number of steps) it drops by & the length of the trail
- Then the loops iterates through the drop trails across the width & trail-length to print/display runes/characters at the position (x,y) corresponding to the loop index
- It uses different color gradients for the style
- Occasionally a 'rabbit' character (🐇) is displayed, watch carefully
- A 'go routine' checks for key-press events and pushes to a channel, main loop reads from channel and if an `Esc` or `Ctl+c` key event is found it exits with small pause
- The actual technical detail of 'terminal display' is all handled by the `tcell` library

### 🔧 Build
- `git clone` or download the code repo
- There is a `Makefile`, so execute `make` and you should get a binary in the `./build` directory
- Or use `go build` command directly

### ⚙️ Output
- This is what it looks like on my terminal
![Demo](./assets/matrixterm-run.gif)


