package utils

func Sum(a int, b int) int {
	return a + b
}

func MultipleSum(a int, b int, c int, d int) (firstSumParams int, secondSumParams int) {
	firstSum := a + b
	secondSum := c + d
	return firstSum, secondSum
}

func Introduce(name string, age int) (string, int) {
	return name, age
}
