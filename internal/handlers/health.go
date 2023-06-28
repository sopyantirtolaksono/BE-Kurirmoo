package handlers

import (
	"runtime"	
	"fmt"
	"math"
)

func (h *handler) Health() (string) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	emil := fmt.Sprint(math.Round(float64(memStats.Alloc)/1000), " kb")

	return emil
}
