package rate

import (
	"math"
	"sync"
	"time"

	"github.com/mohanson/daze/lib/doa"
)

// Limits represents a rate limiter that controls resource allocation over time.
type Limits struct {
	addition uint64
	capacity uint64
	last     time.Time
	mu       sync.Mutex
	size     uint64
	step     time.Duration
}

// Peek glances there are enough resources (n) available.
func (l *Limits) Peek(n uint64) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	cycles := uint64(time.Since(l.last) / l.step)
	if cycles > 0 {
		l.last = l.last.Add(l.step * time.Duration(cycles))
		doa.Doa(cycles <= math.MaxUint64/l.addition)
		doa.Doa(l.size <= math.MaxUint64-l.addition*cycles)
		l.size = l.size + l.addition*cycles
		l.size = min(l.size, l.capacity)
	}
	return l.size >= n
}

// Wait ensures there are enough resources (n) available, blocking if necessary.
func (l *Limits) Wait(n uint64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	cycles := uint64(time.Since(l.last) / l.step)
	if cycles > 0 {
		l.last = l.last.Add(l.step * time.Duration(cycles))
		doa.Doa(cycles <= math.MaxUint64/l.addition)
		doa.Doa(l.size <= math.MaxUint64-l.addition*cycles)
		l.size = l.size + l.addition*cycles
		l.size = min(l.size, l.capacity)
	}
	if l.size < n {
		cycles = (n - l.size + l.addition - 1) / l.addition
		time.Sleep(l.step * time.Duration(cycles))
		l.last = l.last.Add(l.step * time.Duration(cycles))
		l.size = l.size + l.addition*cycles
	}
	l.size -= n
	doa.Doa(l.size <= l.capacity)
}

// NewLimits creates a new rate limiter with rate r over period p.
func NewLimits(r uint64, p time.Duration) *Limits {
	doa.Doa(r > 0)
	doa.Doa(p > 0)
	gcd := func(a, b uint64) uint64 {
		t := uint64(0)
		for b != 0 {
			t = b
			b = a % b
			a = t
		}
		return a
	}(r, uint64(p))
	return &Limits{
		addition: r / gcd,
		capacity: r,
		last:     time.Now(),
		mu:       sync.Mutex{},
		size:     r,
		step:     p / time.Duration(gcd),
	}
}
