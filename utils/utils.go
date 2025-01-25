package utils

func Switch[value comparable](condition bool, ifTrue, ifFalse value) value {
	if condition {
		return ifTrue
	}
	return ifFalse
}
