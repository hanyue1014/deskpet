package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	// "github.com/veandco/go-sdl2/img"
)

const (
	windowWidth  = 200
	windowHeight = 200
)

func main() {
	fmt.Println("Hello from go")

	sdl.Init(sdl.INIT_EVERYTHING)
	// quit sdl to prevent memory leak
	defer sdl.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(windowWidth, windowHeight, sdl.WINDOW_ALWAYS_ON_TOP)
	// close window and renderer properly
	defer w.Destroy()
	defer r.Destroy()

	if err != nil {
		panic(fmt.Sprintf("cannot initialise window or renderer %v", err))
	}

	time.Sleep(5 * time.Second)

	// img.Init(img.INIT_PNG)
}
