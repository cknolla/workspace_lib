package workspace_lib

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// AddNums adds numbers of any numeric type
func AddNums[T Number](a, b T) T {
	return a + b
}

// SubNums can only subtract integers for some reason
func SubNums(a, b int) int {
	return a - b
}
