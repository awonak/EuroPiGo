package europi

import (
	"log"
	"runtime"
	"time"
)

// Clamp returns a value that is no lower than "low" and no higher than "high".
func Clamp[V uint8 | uint16 | int | float32](value, low, high V) V {
	if value >= high {
		return high
	}
	if value <= low {
		return low
	}
	return value
}

func DebugMemoryUsage() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println("Heap Used: ", mem.HeapInuse, " Free: ", mem.HeapIdle, " Sys:", mem.HeapSys, "\r")
}

func DebugMemoryUsedPerSecond() {
	var (
		heapUsed uint64
		mem      runtime.MemStats
	)
	for {
		runtime.ReadMemStats(&mem)
		if mem.HeapInuse < heapUsed {
			log.Println("GC called\r")
			log.Println("Heap used per second: ", mem.HeapSys-mem.HeapInuse, " Total Heap Used: ", mem.HeapInuse, "\r")
		} else {
			log.Println("Heap used per second: ", mem.HeapInuse-heapUsed, " Total Heap Used: ", mem.HeapInuse, "\r")
		}
		heapUsed = mem.HeapInuse
		time.Sleep(time.Second)
	}
}
