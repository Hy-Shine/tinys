package go_utils

import "sync"

type hashSet struct {
	lock sync.Mutex
	m    map[interface{}]struct{}
}

func NewHashSet() *hashSet {
	return &hashSet{
		lock: sync.Mutex{},
		m:    make(map[interface{}]struct{}),
	}
}

func NewHashWithCap(cap int) *hashSet {
	return &hashSet{
		lock: sync.Mutex{},
		m:    make(map[interface{}]struct{}, cap),
	}
}

func (h *hashSet) Add(s ...interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()
	for _, v := range s {
		h.m[v] = struct{}{}
	}
}

func (h *hashSet) Delete(v interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.m, v)
}

func (h *hashSet) Clear() {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.m = make(map[interface{}]struct{})
}

func (h *hashSet) Range() map[interface{}]struct{} {
	h.lock.Lock()
	defer h.lock.Unlock()
	return h.m
}

func (h *hashSet) Len() int {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(h.m)
}

func (h *hashSet) IsExists(e interface{}) bool {
	h.lock.Lock()
	defer h.lock.Unlock()
	_, ok := h.m[e]
	return ok
}
