package go_utils

import "golang.org/x/exp/constraints"

func If[T constraints.Ordered](b bool, trueVal, falseVal T) T {
	if b {
		return trueVal
	}
	return falseVal
}
