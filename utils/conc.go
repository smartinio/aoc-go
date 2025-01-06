package utils

import "sync"

func ConcurrentFunc[T any, E any](items []T, cb func(T) E) []E {
	wg := sync.WaitGroup{}
	results := make([]E, len(items))

	for i, item := range items {
		wg.Add(1)
		go func(i int, item T) {
			defer wg.Done()
			results[i] = cb(item)
		}(i, item)
	}

	wg.Wait()

	return results
}

type SafeCounter[E comparable] struct {
	mu  sync.Mutex
	Map map[E]int
}

func (c *SafeCounter[E]) Inc(key E, value int) {
	c.mu.Lock()
	c.Map[key] += value
	c.mu.Unlock()
}
