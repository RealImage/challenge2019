package models

import "sync"

type Errors struct {
	sync.RWMutex
	Container []error
}

func (e *Errors) Add(err error) {
	e.Lock()
	e.Container = append(e.Container, err)
	e.Unlock()
}
