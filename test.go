package main

import (
	"fmt"
	"time"

	"github.com/MichaelS11/go-hx711"
)

func main() {
	err := hx711.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	hx711, err := hx711.NewHx711("GPIO6", "GPIO5")
	if err != nil {
		fmt.Println("NewHx711 error:", err)
		return
	}

	defer hx711.Shutdown()

	err = hx711.Reset()
	if err != nil {
		fmt.Println("Reset error:", err)
		return
	}

	var data int
	for i := 0; i < 10000; i++ {
		time.Sleep(200 * time.Microsecond)

		data, err = hx711.ReadDataRaw()
		if err != nil {
			fmt.Println("ReadDataRaw error:", err)
			continue
		}

		fmt.Println(data)
	}

}
Calibrate the readings / get AdjustZero & AdjustScale values
To get the values needed to calibrate the scale's readings will need at least one weight of known value. Having two weights is preferred. In the below program change weight1 and weight2 to the known weight values. Make sure to set it in the unit of measurement that you prefer (pounds, ounces, kg, g, etc.). To start, make sure there are no weight on the scale. Run the program. When asked, put the first weight on the scale. Then when asked, put the second weight on the scale. It will print out the AdjustZero and AdjustScale values. Those are using in the next example.

Please note that temperature affects the readings. Also if you are planning on reading the weight often, maybe want to do that for about 20 minutes before calibration.

package main

import (
	"fmt"

	"github.com/MichaelS11/go-hx711"
)

func main() {
	err := hx711.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	hx711, err := hx711.NewHx711("GPIO6", "GPIO5")
	if err != nil {
		fmt.Println("NewHx711 error:", err)
		return
	}
  
	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	// hx711.SetGain(128)

	var weight1 float64
	var weight2 float64

	weight1 = 100
	weight2 = 200

	hx711.GetAdjustValues(weight1, weight2)
}
