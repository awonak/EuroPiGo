package europi

import (
	"log"
	"runtime"
)

// Clamp returns a value that is no lower than `low“ and no higher than `high“.
func Clamp[V int | float32](value, low, high V) V {
	if value > high {
		value = high
	}
	if value < low {
		value = low
	}
	return value
}

func DebugMemoryUsage() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println("Heap Used: ", mem.HeapInuse, " Free: ", mem.HeapIdle, " Sys:", mem.HeapSys, "\r")
}
