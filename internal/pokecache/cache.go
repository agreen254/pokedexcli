package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]Entry
	mu       *sync.Mutex
	interval time.Duration
}

type Entry struct {
	createdAt time.Time
	Value     []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Entries:  map[string]Entry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}
	go c.reapLoop()

	return c
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		t := <-ticker.C
		// fmt.Println("ticker emitted:", t)
		for key, value := range c.Entries {
			if t.Sub(value.createdAt) > c.interval {
				c.delete(key)
			}
		}
	}
}

func (c Cache) delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// fmt.Println("DELETED entry", key)
	delete(c.Entries, key)
}

func (c Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entries[key] = Entry{
		createdAt: time.Now(),
		Value:     value,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.Entries[key]
	if !ok {
		return []byte{}, false
	} else {
		return entry.Value, true
	}
}
