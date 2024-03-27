package set

type gset[K comparable] struct {
	c int
	m map[K]struct{}
}

func NewSet[K comparable](cap int) *gset[K] {
	if cap < 0 {
		cap = 0
	}
	return &gset[K]{
		c: cap,
		m: make(map[K]struct{}, cap),
	}
}

func (s *gset[K]) Add(v K) {
	s.m[v] = struct{}{}
}

func (s *gset[K]) Delete(v K) {
	delete(s.m, v)
}

func (s *gset[K]) IsExists(v K) bool {
	_, ok := s.m[v]
	return ok
}

func (s *gset[K]) Len() int {
	return len(s.m)
}

func (s *gset[K]) Keys() []K {
	keys := make([]K, 0, s.Len())
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s *gset[K]) Clear() {
	s.m = make(map[K]struct{}, s.c)
}
