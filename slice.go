package go_utils

import (
	"strings"

	"golang.org/x/exp/constraints"
)

func ListFirst[T constraints.Ordered](list []T) T {
	var first T
	if len(list) > 0 {
		first = list[0]
	}

	return first
}

func ListDistinct[T constraints.Ordered](list []T) []T {
	if len(list) <= 1 {
		return list
	}

	distinct := make(map[T]struct{}, len(list))
	result := make([]T, 0, len(list))
	for i := range list {
		_, ok := distinct[list[i]]
		if !ok {
			distinct[list[i]] = struct{}{}
			result = append(result, list[i])
		}
	}
	return result
}

func ListMerge[T constraints.Ordered](origin, target []T) []T {
	merged := make([]T, len(origin)+len(target))
	copy(merged, origin)
	copy(merged[len(origin):], target)

	return merged
}

func SliceToStr(slice []string) string {
	l := make([]string, 0)

	for _, v := range slice {
		if v != "" {
			l = append(l, v)
		}
	}

	return strings.Join(l, ",")
}

func findInt(l []int, target int) int {
	const notFound = -1
	if l[0] > target || l[len(l)-1] < target {
		return notFound
	}

	low, high := 0, len(l)-1
	for low <= high {
		middle := (low + high) / 2
		if l[middle] == target {
			return middle
		}
		if l[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return notFound
}

// 1,2,3,4,5 10 --> -1
func findFirstGreatThan(l []int, target int) int {
	if l[len(l)-1] < target {
		return -1
	}
	if l[0] > target {
		return 0
	}
	low, high := 0, len(l)-1
	middle := 0
	for low <= high {
		middle = (low + high) / 2
		if l[middle] > target && l[middle-1] <= target {
			break
		}
		if l[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return middle
}
