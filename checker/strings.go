package checker

func NotEmpty(args ...string) int {
	for i, arg := range args {
		if len(arg) < 1 {
			return i
		}
	}
	return 0
}
