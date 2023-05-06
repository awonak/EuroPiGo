package europi

import (
	"context"
	"log"
	"runtime"
	"time"
)

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

// used for non-pico testing of europi apps
var (
	activateNonPicoWebSocket func(ctx context.Context, e Hardware) NonPicoWSActivation
)

type NonPicoWSActivation interface {
	Shutdown() error
}

func ActivateNonPicoWS(ctx context.Context, e Hardware) NonPicoWSActivation {
	if activateNonPicoWebSocket == nil {
		return nil
	}
	return activateNonPicoWebSocket(ctx, e)
}
