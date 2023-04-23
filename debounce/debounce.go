package debounce

import "time"

type Debouncer[T func(U, time.Duration), U any] interface {
	Debounce(delay time.Duration) func(U)
	LastChange() time.Time
}

type debouncer[T func(U, time.Duration), U any] struct {
	last time.Time
	fn   T
}

func NewDebouncer[T func(U, time.Duration), U any](fn T) Debouncer[T, U] {
	return &debouncer[T, U]{
		last: time.Now(),
		fn:   fn,
	}
}

func (d *debouncer[T, U]) Debounce(delay time.Duration) func(U) {
	return func(u U) {
		now := time.Now()
		if deltaTime := now.Sub(d.last); deltaTime >= delay {
			d.last = now
			d.fn(u, deltaTime)
		}
	}
}

func (d debouncer[T, U]) LastChange() time.Time {
	return d.last
}
