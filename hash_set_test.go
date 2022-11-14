package go_utils

import (
	"fmt"
	"sync"
	"testing"
)

func TestHashSet_Add(t *testing.T) {
	hash := NewHashSet()
	wg := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		n := i
		go func() {
			hash.Add(n, n+1)
			wg.Done()
		}()
	}
	wg.Wait()
	for i := range hash.Range() {
		fmt.Println(i)
	}
}

func BenchmarkHashSet_Add(b *testing.B) {
	hash := NewHashSet()
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		n := i
		go func() {
			hash.Add(n)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(hash.Len())
}
