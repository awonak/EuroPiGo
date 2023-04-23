package event

import "sync"

type Bus interface {
	Subscribe(subject string, callback func(msg any))
	Post(subject string, msg any)
}

func Subscribe[TMsg any](bus Bus, subject string, callback func(msg TMsg)) {
	bus.Subscribe(subject, func(msg any) {
		if tmsg, ok := msg.(TMsg); ok {
			callback(tmsg)
		}
	})
}

func NewBus() Bus {
	b := &bus{}
	return b
}

type bus struct {
	chMap sync.Map
}

func (b *bus) Subscribe(subject string, callback func(msg any)) {
	b.chMap.Store(subject, callback)
}

func (b *bus) Post(subject string, msg any) {
	cb, found := b.chMap.Load(subject)
	if !found {
		return
	}

	if callback, ok := cb.(func(any)); ok {
		callback(msg)
	}
}
