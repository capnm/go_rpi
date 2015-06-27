// A fullscreen clock with some hardware data.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"

	g "openvg"
)

func main() {
	done := make(chan bool, 1)
	exitHandler(done)

	width, height := g.Init()
	dim := fmt.Sprintf("w=%d h=%d", width, height)

	w2 := float32(width / 2)
	h2 := float32(height / 2)
	w := float32(width)

	for {
		select {
		case <-done:
			fmt.Println("done")
			// jump outside the for loop
			goto end

		default:
			g.Start(width, height)
			g.BackgroundColor("black")

			g.FillRGB(44, 100, 232, 1) // blue 1/2 circle
			g.Circle(w2, 0, w)
			//g.Circle(w2, h2, h2*2)

			g.FillColor("white")
			l1, l2, l3 := clock()
			g.TextMid(w2, h2+h2/2, l1, "serif", width/15)
			g.TextMid(w2, h2, l2, "serif", width/7)
			g.TextMid(w2, h2-h2/2+100, l3, "serif", width/15)
			g.TextMid(w2, h2-h2/2-20, dim, "serif", width/30)

			l4 := hwinfo()
			g.TextMid(w2, h2-h2/2-100, l4, "serif", width/30)

			l5 := cpuTemperature()
			g.TextMid(w2, h2-h2/2-180, l5, "serif", width/30)

			g.End()
			time.Sleep(1000 * time.Millisecond)
		}
	}
end:
	g.Finish()
}

func clock() (string, string, string) {
	t := time.Now()
	return t.Local().Format("Mon 2006-01-02"),
		t.Local().Format("15:04:05"),
		t.Local().Format("-0700 MST")
}

// Read the CPU/GPU clock frequence.
func hwinfo() string {
	b1, err := exec.Command("vcgencmd", "measure_clock", "arm").Output()
	if err != nil {
		log.Println(err)
	}
	b2, err := exec.Command("vcgencmd", "measure_clock", "core").Output()
	if err != nil {
		log.Println(err)
	}
	f_arm := parseFreq(b1)
	f_core := parseFreq(b2)
	return fmt.Sprintf("arm=%sMhz core=%sMHz", f_arm, f_core)
}

// Read the CPU temperature.
func cpuTemperature() string {
	b, err := exec.Command("vcgencmd", "measure_temp").Output()
	if err != nil {
		log.Println(err)
	}
	return string(b)
}

// vcgencmd measure_clock arm/core
// 	frequency(45)=600_000000
// 	frequency(1)=250000000
func parseFreq(b []byte) string {
	i := 0
	for ; i < len(b); i++ {
		if b[i] == '=' {
			break
		}
	}
	if len(b[i:]) > 7 {
		return string(b[i+1 : len(b)-7])
	}
	return "?"
}

// Wait for the 'ctrl+c' or the 'enter' key.
// Wenn the key is pressed, send a 'true' over the channel 'done'.
func exitHandler(done chan<- bool) {
	c := make(chan os.Signal, 1)

	// ctrl+c interrupt
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			done <- true
		}
	}()

	go func() {
		// Wait for the return key.
		bufio.NewReader(os.Stdin).ReadByte()
		done <- true
	}()
}

// Just for testing, is insecure and slow.
func k_clk() string {
	b, err := exec.Command("sudo", "cat", "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_cur_freq").Output()
	if err != nil {
		log.Println(err)
	}
	return string(b)
}
