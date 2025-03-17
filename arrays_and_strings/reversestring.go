package arrays_and_strings

func ReverseString(s []byte) []byte {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}
