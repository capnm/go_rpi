// A fullscreen clock with some hardware data.
package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"openvg"
)

func main() {
	width, height := openvg.Init()
	dim := fmt.Sprintf("w=%d h=%d", width, height)

	w2 := float32(width / 2)
	h2 := float32(height / 2)
	w := float32(width)

	for {
		openvg.Start(width, height)
		openvg.BackgroundColor("black")

		openvg.FillRGB(44, 100, 232, 1) // blue 1/2 circle
		openvg.Circle(w2, 0, w)
		//openvg.Circle(w2, h2, h2*2)

		openvg.FillColor("white")
		l1, l2, l3 := clock()
		openvg.TextMid(w2, h2+h2/2, l1, "serif", width/15)
		openvg.TextMid(w2, h2, l2, "serif", width/7)
		openvg.TextMid(w2, h2-h2/2+100, l3, "serif", width/15)
		openvg.TextMid(w2, h2-h2/2-20, dim, "serif", width/30)

		l4 := hwinfo()
		openvg.TextMid(w2, h2-h2/2-100, l4, "serif", width/30)

		l5 := cpuTemperature()
		openvg.TextMid(w2, h2-h2/2-180, l5, "serif", width/30)

		openvg.End()
		time.Sleep(1000 * time.Millisecond)
	}
	//TODO catch keys or signals
	openvg.Finish()
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
	if len(b1) < 17 || len(b2) < 17 {
		return "error: parse hw data"
	}
	return fmt.Sprintf("arm=%sMhz core=%sMHz", b1[14:17], b2[13:16])
}

// Read the CPU temperature.
func cpuTemperature() string {
	b, err := exec.Command("vcgencmd", "measure_temp").Output()
	if err != nil {
		log.Println(err)
	}
	return string(b)
}

// Just for testing, is insecure and slow.
func k_clk() string {
	b, err := exec.Command("sudo", "cat", "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_cur_freq").Output()
	if err != nil {
		log.Println(err)
	}
	return string(b)
}
