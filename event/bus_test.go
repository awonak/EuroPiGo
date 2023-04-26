package event_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/awonak/EuroPiGo/event"
)

func TestBus(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		if actual := event.NewBus(); actual == nil {
			t.Fatal("Bus NewBus: expected[non-nil] actual[nil]")
		}
	})

	t.Run("Post", func(t *testing.T) {
		bus := event.NewBus()

		t.Run("Unsubscribed", func(t *testing.T) {
			// test to see if we block on a post to an unsubscribed subject
			subject := fmt.Sprintf("testing_%v", time.Now().UnixNano())

			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				bus.Post(subject, "hello world")
			}()

			// if we somehow block here, then the test deadline will eventually be reached
			// and then it will be failed automatically
			wg.Wait()
		})
	})

	t.Run("Subscribe", func(t *testing.T) {
		bus := event.NewBus()
		t.Run("Untyped", func(t *testing.T) {
			var actual any
			subject := "untyped"
			bus.Subscribe(subject, func(msg any) {
				actual = msg
			})

			expected := 5
			bus.Post(subject, expected)
			if actual != expected {
				t.Fatalf("Bus Subscribe(%v): expected[%v] actual[%v]", subject, expected, actual)
			}
		})

		t.Run("Typed", func(t *testing.T) {
			var actual string
			subject := "typed"
			event.Subscribe(bus, subject, func(msg string) {
				actual = msg
			})

			expected := "hello world"
			bus.Post(subject, expected)
			if actual != expected {
				t.Fatalf("Bus Subscribe(%v): expected[%v] actual[%v]", subject, expected, actual)
			}
		})
	})

	t.Run("Unsubscribe", func(t *testing.T) {
		bus := event.NewBus()

		t.Run("Subscribed", func(t *testing.T) {
			subject := "unsub"

			var actual any
			bus.Subscribe(subject, func(msg any) {
				actual = msg
			})

			bus.Unsubscribe(subject)

			var expected any
			bus.Post(subject, "hello world")
			if actual != expected {
				t.Fatalf("Bus Unsubscribe(%v): expected[%v] actual[%v]", subject, expected, actual)
			}
		})
	})
}
