package go_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListMerge(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5, 6, 7}
	list2 := []int{4, 5, 6, 7}

	result := ListMerge(list1, list2)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 4, 5, 6, 7}, result)
}

func BenchmarkListMerge(b *testing.B) {
	list1 := make([]int, 0, 500)
	for i := 0; i < 500; i++ {
		list1 = append(list1, i)
	}
	list2 := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for i := 0; i < b.N; i++ {
		_ = ListMerge(list1, list2)
	}
}
