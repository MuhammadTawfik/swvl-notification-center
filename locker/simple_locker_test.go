package locker

import (
	"testing"
	"time"
)

func TestLocking(t *testing.T) {
	lock := CreateNewLock()
	lock.Lock("object1")

	if !lock.IsLocked("object1") {
		// t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
		t.Errorf("it doesn't do simple locking scenario, locked(object1), but lock.IsLocked(object1) => false")
	}
}

func TestUnLocking(t *testing.T) {
	lock := CreateNewLock()
	lock.Lock("object1")
	lock.Unlock("object1")

	if lock.IsLocked("object1") {
		// t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
		t.Errorf("it doesn't do simple unlocking scenario, locked(object1), then unlock(object1) but lock.IsLocked(object1) => true")
	}
}

func TestLockingForPeriod(t *testing.T) {
	lock := CreateNewLock()
	duration := 2 * time.Second
	lock.LockFor("object1", duration)
	time.Sleep(duration)

	if lock.IsLocked("object1") {
		// t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
		t.Errorf("it doesn't do unlocking after period scenario, locked(object1) for 2 seconds, waited for 2 seconds but lock.IsLocked(object1) => true")
	}
}
