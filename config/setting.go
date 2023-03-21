package config

import (
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"
)

const (
	picoFlashStart = 0x10000000
)

type Setting[T any] struct {
	Value       T
	initialized bool
}

// Init initializes the setting with the expected default or loads the saved value from pico flash
func (s *Setting[T]) Init(value T) {
	valuePtr := unsafe.Slice(&s.Value, 1)
	vfp := unsafe.Add(unsafe.Pointer(&s.Value), picoFlashStart)
	valueFlashPtr := unsafe.Slice((*T)(vfp), 1)
	initializedFlashPtr := (*uint8)(unsafe.Add(unsafe.Pointer(&s.initialized), picoFlashStart))

	state := interrupt.Disable()
	if iv := volatile.LoadUint8(initializedFlashPtr); iv == 0 {
		s.Value = value
	} else {
		copy(valuePtr, valueFlashPtr)
		s.initialized = true
	}
	interrupt.Restore(state)
}

// Flush flushes the value out to pico flash storage along with the initialize flag
func (s *Setting[T]) Flush() {
	valuePtr := unsafe.Slice(&s.Value, 1)
	vfp := unsafe.Add(unsafe.Pointer(&s.Value), picoFlashStart)
	valueFlashPtr := unsafe.Slice((*T)(vfp), 1)
	initializedFlashPtr := (*uint8)(unsafe.Add(unsafe.Pointer(&s.initialized), picoFlashStart))

	state := interrupt.Disable()
	copy(valueFlashPtr, valuePtr)
	volatile.StoreUint8(initializedFlashPtr, 1)
	interrupt.Restore(state)
}
