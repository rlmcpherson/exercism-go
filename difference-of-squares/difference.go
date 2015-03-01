package diffsquares

func SquareOfSums(n int) int {
	sumn := sum(n)
	return sumn * sumn
}

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n-1)
}

func SumOfSquares(n int) int {

	if n == 0 {
		return 0
	}
	return n*n + SumOfSquares(n-1)
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
