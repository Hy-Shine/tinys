package slice

import (
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func FirstEle[T any](list []T) T {
	var first T
	if len(list) > 0 {
		first = list[0]
	}
	return first
}

func Distinct[T comparable](list []T) []T {
	if len(list) <= 1 {
		return list
	}

	distinct := make(map[T]struct{}, 3*len(list)/2)
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

func Merge[T constraints.Ordered](origin, target []T) []T {
	merged := make([]T, len(origin)+len(target))
	copy(merged, origin)
	copy(merged[len(origin):], target)

	return merged
}

func Contains[T comparable](l []T, target T) bool {
	for _, v := range l {
		if v == target {
			return true
		}
	}
	return false
}

func StrsContains[K ~string](l []K, target K) bool {
	for i := range l {
		if strings.Contains(string(l[i]), string(target)) {
			return true
		}
	}
	return false
}

func Split[T any](list []T, size int) [][]T {
	if size <= 0 {
		size = 1
	}

	result := make([][]T, 0, len(list)/size+1)
	for i := 0; i < len(list); i += size {
		end := i + size
		if end > len(list) {
			end = len(list)
		}
		result = append(result, list[i:end])
	}
	return result
}

func IntsToStrings[K constraints.Integer](l []K) []string {
	result := make([]string, len(l))
	for i, v := range l {
		result[i] = strconv.FormatInt(int64(v), 10)
	}
	return result
}

func StringsToInts[T constraints.Integer](list []string) []T {
	result := make([]T, 0, len(list))
	for _, v := range list {
		n, _ := strconv.ParseInt(v, 10, 64)
		result = append(result, T(n))
	}
	return result
}

func ToSet[T comparable](l []T) map[T]struct{} {
	m := make(map[T]struct{}, 2*len(l)/3)
	for i := range l {
		m[l[i]] = struct{}{}
	}
	return m
}

func ToSetFunc[T comparable, V any](l []any, f func(in any) (key T, value V)) map[T]V {
	m := make(map[T]V, len(l))
	for _, v := range l {
		key, value := f(v)
		m[key] = value
	}
	return m
}

func Reverse[T any](nums []T) {
	if len(nums) == 0 {
		return
	}

	left, right := 0, len(nums)-1
	for left <= right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
