package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListMerge(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5, 6, 7}
	list2 := []int{4, 5, 6, 7}

	result := Merge(list1, list2)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 4, 5, 6, 7}, result)
}

func BenchmarkListMerge(b *testing.B) {
	list1 := make([]int, 0, 500)
	for i := 0; i < 1000; i++ {
		list1 = append(list1, i)
	}
	list2 := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for i := 0; i < b.N; i++ {
		_ = Merge(list1, list2)
	}
}

func TestStrsContains(t *testing.T) {
	cases := []struct {
		name   string
		l      []string
		target string
		want   bool
	}{
		{
			name:   "empty list",
			l:      []string{},
			target: "foo",
			want:   false,
		},
		{
			name:   "target found in first element",
			l:      []string{"foo", "bar", "baz"},
			target: "foo",
			want:   true,
		},
		{
			name:   "target found in last element",
			l:      []string{"foo", "bar", "baz"},
			target: "baz",
			want:   true,
		},
		{
			name:   "target found in middle element",
			l:      []string{"foo", "bar", "baz"},
			target: "bar",
			want:   true,
		},
		{
			name:   "target not found",
			l:      []string{"foo", "bar", "baz"},
			target: "qux",
			want:   false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := StrsContains(c.l, c.target)
			if got != c.want {
				t.Errorf("StrsContains(%q, %q) == %t, want %t", c.l, c.target, got, c.want)
			}
		})
	}
}
