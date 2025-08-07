package kyu

func GetSum(a, b int) int {
	if a > b {
		a, b = b, a
	}
	n := b - a + 1
	return n * (a + b) / 2
}
