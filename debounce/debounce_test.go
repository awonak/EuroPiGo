package debounce_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/awonak/EuroPiGo/debounce"
)

func TestDebouncer(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		fn := func(any, time.Duration) {}
		if actual := debounce.NewDebouncer(fn); actual == nil {
			t.Fatalf("Debouncer[%p] NewDebouncer: expected[non-nil] actual[nil]", fn)
		}
	})

	t.Run("Debounce", func(t *testing.T) {
		t.Run("Delay0", func(t *testing.T) {
			delay := time.Duration(0)
			runs := 4
			runDebouncerTest(t, runs, delay, time.Duration(0), 4)
			runDebouncerTest(t, runs, delay, time.Millisecond*1, 4)
			runDebouncerTest(t, runs, delay, time.Second*1, 4)
		})

		t.Run("Delay100ms", func(t *testing.T) {
			delay := time.Millisecond * 100
			runs := 4
			runDebouncerTest(t, runs, delay, time.Duration(0), 0)
			runDebouncerTest(t, runs, delay, time.Millisecond*1, 0)
			runDebouncerTest(t, runs, delay, time.Second*1, 2)
		})

		t.Run("Delay10s", func(t *testing.T) {
			delay := time.Millisecond * 100
			runs := 4
			runDebouncerTest(t, runs, delay, time.Duration(0), 0)
			runDebouncerTest(t, runs, delay, time.Millisecond*1, 0)
			runDebouncerTest(t, runs, delay, time.Second*1, 0)
		})
	})

	t.Run("LastChange", func(t *testing.T) {
		fn := func(any, time.Duration) {}
		db := debounce.NewDebouncer(fn)
		dbFunc := db.Debounce(0)

		beforeExpected := time.Now()
		dbFunc(nil)
		if actual := db.LastChange(); !actual.After(beforeExpected) {
			t.Fatalf("Debouncer[%p] LastChange: expected(after)[%v] actual[%v]", fn, beforeExpected, actual)
		}
	})
}

func runDebouncerTest(t *testing.T, runs int, delay, interval time.Duration, minimumExpected int) {
	t.Helper()

	var testName string
	if interval == 0 {
		testName = "Immediate"
	} else {
		testName = fmt.Sprintf("Interval%v", interval)
	}

	var actual int
	fn := func(any, time.Duration) {
		actual++

	}

	db := debounce.NewDebouncer(fn).Debounce(delay)

	t.Run(testName, func(t *testing.T) {
		for i := 0; i < runs; i++ {
			db(nil)
			time.Sleep(interval)
		}
		// since these are timing-based, we have to be lenient
		if actual < minimumExpected {
			t.Fatalf("Debouncer[%p] Debounce(%v): expected[%v] actual[%v]", fn, delay, minimumExpected, actual)
		}
	})
}
