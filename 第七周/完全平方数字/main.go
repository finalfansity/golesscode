func numSquares(n int) int {
	count := int(math.Floor(math.Sqrt(float64(n))))
	f := make([]int, n+1)
	for i := range f {
		f[i] = i
	}
	for i := 0; i <= count; i++ {
		for j := i * i; j <= n; j++ {
			f[j] = Min(f[j], f[j-i*i]+1)
		}
	}
	return f[n]
}
func Min(args ...int) int {
	min := args[0]
	for _, i := range args {
		if i < min {
			min = i
		}
	}
	return min
}
