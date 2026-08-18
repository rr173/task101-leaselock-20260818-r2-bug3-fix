package store

import (
	"testing"
	"time"

	"task101-leaselock/internal/lease"
)

func TestListWaitersReturnsOldestFirst(t *testing.T) {
	s, _ := newTestStore(t)
	defer s.Close()
	now := time.Unix(300, 0)
	_, _ = s.Acquire("X", "owner", now, time.Minute)
	first, _ := s.EnqueueWaiter("X", "first", time.Minute, now.Add(time.Second))
	second, _ := s.EnqueueWaiter("X", "second", time.Minute, now.Add(2*time.Second))
	got, err := s.ListWaiters("X", lease.WaiterPending)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].ID != first.ID || got[1].ID != second.ID {
		t.Fatalf("waiters=%+v want [%s %s]", got, first.ID, second.ID)
	}
}
