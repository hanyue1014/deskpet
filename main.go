package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 400
	windowHeight = 400
)

func main() {
	fmt.Println("Hello from go")

	sdl.Init(sdl.INIT_EVERYTHING)
	// quit sdl to prevent memory leak
	defer sdl.Quit()

	// slight note: can use sdl.WINDOW_ALWAYS_ON_TOP | sdl.WINDOW_BORDERLESS as flag

	w, r, err := sdl.CreateWindowAndRenderer(windowWidth, windowHeight, sdl.WINDOW_ALWAYS_ON_TOP|sdl.WINDOW_BORDERLESS)
	// close window and renderer properly
	defer w.Destroy()
	defer r.Destroy()

	if err != nil {
		panic(fmt.Sprintf("cannot initialise window or renderer %v", err))
	}

	// setting window to 0 opacity just makes the whole window not visible
	// if err = w.SetWindowOpacity(0); err != nil {
	// 	panic(fmt.Sprintf("Cannot set window to 0 opacity: %v", err))
	// }

	// img.Init(img.INIT_PNG)
	t, err := img.LoadTexture(r, "./res/paimon.png")
	defer t.Destroy()

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load image: %v", err)
	}

	if err = r.SetDrawColor(255, 255, 255, 0); err != nil {
		fmt.Printf("Cannot set renderer color: %v", err)
	}

	r.Clear()
	if err = r.Copy(t, nil, nil); err != nil {
		fmt.Printf("Error copying texture: %v", err)
	}
	r.Present()
	time.Sleep(5 * time.Second)
}
