package functions

func Apply(x int, f func(int) int) int {
	return f(x)
}
