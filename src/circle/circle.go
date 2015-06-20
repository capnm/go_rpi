// Draw a blue marble on the RPi screen.
package main

import (
	"bufio"
	"os"

	g "openvg"
)

func main() {
	width, height := g.Init()

	w2 := float32(width / 2)
	w := float32(width)

	g.Start(width, height)
	g.BackgroundColor("black")

	// Draw a blue 1/2 circle.
	g.FillRGB(44, 100, 232, 1)
	g.Circle(w2, 0, w)
	g.End()

	// Wait for the return key.
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	g.Finish()
}
