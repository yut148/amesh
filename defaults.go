package amesh

import (
	"fmt"
	"log"
	"time"
)

// Default observer properties
const (
	DefaultIterationDuration    = 3 * time.Minute
	DefaultNotificationInterval = 20 * time.Minute
)

// DefaultOnRainHandleFunc ...
func DefaultOnRainHandleFunc(event Event) error {
	log.Println(fmt.Sprintf("IT'S RAINING NOW!! %v", event.Timestamp))
	return nil
}

// DefaultIsRainingFunc ...
// とりあえず全ピクセル舐めで
// ちょっとでも雨のピクセルが全体の30%を越えてたら雨ってことにする
func DefaultIsRainingFunc(ev Event) bool {
	max := ev.Img.Bounds().Max
	var hit, all float64 = 0, float64(max.X) * float64(max.Y)
	for y := 1; y < max.Y-1; y++ {
		for x := 1; x < max.X-1; x++ {
			r, g, b, a := ev.Img.At(x, y).RGBA()
			if r+g+b+a > 100 {
				hit++
			}
		}
	}
	var threshold float64 = 30
	if (hit*100)/all > threshold {
		return true
	}
	return false // 快晴だこれ
}
