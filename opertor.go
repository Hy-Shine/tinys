package go_utils

func If[T any](b bool, trueVal, falseVal T) T {
	if b {
		return trueVal
	}
	return falseVal
}
