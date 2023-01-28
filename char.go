package go_utils

func IsAlphabet(s byte) bool {
	return s >= 'A' && s <= 'Z' || s >= 'a' && s <= 'z'
}

func IsNumber(s byte) bool {
	return s >= '0' && s <= '9'
}
