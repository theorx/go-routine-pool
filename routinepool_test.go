package routinepool

import (
	"testing"
	"time"
)

func TestNewPoolCreatesPoolWithCorrectSize(t *testing.T) {
	rp := NewPool(124)
	if rp.Size() != 124 {
		t.Errorf("Pool has incorrect number of workers. Want: 124, got: %d", rp.Size())
	}
}

func TestNewPoolWithNegativeWorkersCreatesWithOneWorker(t *testing.T) {
	rp := NewPool(-44)
	if rp.Size() != 1 {
		t.Errorf("Pool has incorrect number of workers. Want: 1, got: %d", rp.Size())
	}
}

func TestAllocateAndRelease(t *testing.T) {
	rp := NewPool(10)
	rp.Allocate()

	if rp.Used() != 1 {
		t.Errorf("Pool has incorrect number of allocations. Want 1, got: %d", rp.Used())
	}

	rp.Allocate()
	if rp.Used() != 2 {
		t.Errorf("Pool has incorrect number of allocations. Want 2, got: %d", rp.Used())
	}

	rp.Release()
	rp.Release()

	if rp.Used() != 0 {
		t.Errorf("Pool has incorrect number of allocations after releases. Want 0, got: %d", rp.Used())
	}
}

func TestAllocateBlocksWhenPoolIsExhausted(t *testing.T) {
	rp := NewPool(2)
	rp.Allocate()
	rp.Allocate()

	go func() {
		rp.Allocate()
		t.Errorf("Shouldn't be able to allocate here, expected to never reach when Pool is exhausted")
	}()
	time.Sleep(time.Millisecond * 10)
}

func TestAllocateAndReleaseResultsWithUsedOfZero(t *testing.T) {
	rp := NewPool(12)
	rp.Allocate()
	rp.Allocate()

	rp.Release()
	rp.Release()

	if rp.Used() != 0 {
		t.Errorf("Used should be zero, got: %d", rp.Used())
	}
}
