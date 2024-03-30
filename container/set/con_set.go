package set

import "sync"

type hashSet[K comparable] struct {
	lock sync.RWMutex
	m    map[K]struct{}
}

func New[K comparable](cap int) *hashSet[K] {
	if cap < 0 {
		cap = 0
	}
	return &hashSet[K]{
		lock: sync.RWMutex{},
		m:    make(map[K]struct{}),
	}
}

func (h *hashSet[K]) Add(l ...K) {
	h.lock.Lock()
	for _, v := range l {
		h.m[v] = struct{}{}
	}
	h.lock.Unlock()
}

func (h *hashSet[K]) Delete(v K) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.m, v)
}

func (h *hashSet[K]) Clear() {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.m = make(map[K]struct{})
}

func (h *hashSet[K]) Keys() []K {
	h.lock.Lock()
	keys := make([]K, 0, h.Len())
	for k := range h.m {
		keys = append(keys, k)
	}
	h.lock.Unlock()
	return keys
}

func (h *hashSet[K]) Len() int {
	h.lock.Lock()
	length := len(h.m)
	h.lock.Unlock()
	return length
}

func (h *hashSet[K]) IsExists(e K) bool {
	h.lock.RLock()
	_, ok := h.m[e]
	h.lock.Unlock()
	return ok
}
