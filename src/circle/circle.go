// Draw an OpenVG circle on the RPi famebuffer.
package main

import (
	"bufio"
	"os"

	"openvg"
)

func main() {
	width, height := openvg.Init()

	w2 := float64(width / 2)
	w := float64(width)

	openvg.Start(width, height)
	openvg.BackgroundColor("black")

	// draw a blue 1/2 circle
	openvg.FillRGB(44, 100, 232, 1)
	openvg.Circle(w2, 0, w)
	openvg.End()

	// wait for the return key
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	openvg.Finish()
}
