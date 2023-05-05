package common

import "time"

// ContextPi gives the EuroPi hardware the components necessary
// to perform rudimentary context operations
type ContextPi struct{}

func (ContextPi) Deadline() (deadline time.Time, ok bool) {
	return
}

func (ContextPi) Done() <-chan struct{} {
	return nil
}

func (ContextPi) Err() error {
	return nil
}

func (ContextPi) Value(key any) any {
	return nil
}
