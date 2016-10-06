package roundrobin

import (
	"errors"
	"fmt"
	"sync"
)

type RoundRobin interface {
	Available() (string, error)
	NoLongerInUse(string) error
	Add(string) error
	Remove(string) error
}

type InMemory struct {
	sync.RWMutex
	available map[string]bool
}

func NewInMemory() *InMemory {
	return &InMemory{
		available: make(map[string]bool, 10),
	}
}

func (i *InMemory) Available() (string, error) {
	i.RLock()
	defer i.RUnlock()
	if len(i.available) < 1 {
		return "", errors.New("None added")
	}
	for k, v := range i.available {
		fmt.Println(k, v)
		if v {
			i.available[k] = false
			return k, nil
		}
	}
	return "", errors.New("None available")
}

func (i *InMemory) NoLongerInUse(noLongerInUse string) error {
	i.Lock()
	defer i.Unlock()
	_, ok := i.available[noLongerInUse]
	if !ok {
		return errors.New("Object removed, cannot be marked as available")
	}
	i.available[noLongerInUse] = true
	return nil
}

func (i *InMemory) Add(add string) error {
	i.Lock()
	defer i.Unlock()
	i.available[add] = true
	return nil
}

func (i *InMemory) Remove(remove string) error {
	i.Lock()
	defer i.Unlock()
	if len(i.available) < 1 {
		return errors.New("None to remove from")
	}
	delete(i.available, remove)
	return nil
}
