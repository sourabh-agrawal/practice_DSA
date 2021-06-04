package algorithms

// FROM: https://www.hackerrank.com/challenges/drawing-book/problem
func pageCount(n, p int32) int32 {
	var distFromBack int32
	distFromFront := p / 2

	if n%2 == 0 {
		distFromBack = (n - p + 1) / 2
	} else {
		distFromBack = (n - p) / 2
	}

	return getMin(distFromFront, distFromBack)
}

func getMin(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}
