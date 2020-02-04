package locker

import (
	"time"
)

type Locker interface {
	Lock(resource string)
	Unlock(resource string)
	LockFor(resource string, duration time.Duration)
	IsLocked(resource string) bool
}

func CreateNewLock() Locker {
	v := make(map[string]bool)
	return simpleLock{
		v,
	}
}

type simpleLock struct {
	data map[string]bool
}

func (s simpleLock) Lock(resource string) {
	s.data[resource] = true

}

func (s simpleLock) Unlock(resource string) {
	delete(s.data, resource)
}

func (s simpleLock) LockFor(resource string, duration time.Duration) {
	s.Lock(resource)
	go s.remove_lock_after(resource, duration)
}

func (s simpleLock) IsLocked(resource string) bool {
	_, ok := s.data[resource]
	return ok
}

func (s simpleLock) remove_lock_after(resource string, duration time.Duration) {
	time.Sleep(duration)
	s.Unlock(resource)
}
