package entity

// OldFibonacci calculates the nth fibonacci sequence number
func OldFibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return OldFibonacci(n-1) + OldFibonacci(n-2)
}

// fibonacci calculates the nth fibonacci sequence number
// I introduced a slice as a cache here to avoid unnecessary recursive calls,
// so when the index of N is calculated it fetches the calculated value in the slice.
func fibonacci(n int, cache []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if cache == nil {
		cache = make([]int, n+1)
	}
	if cache[n] == 0 {
		cache[n] = fibonacci(n-1, cache) + fibonacci(n-2, cache)
	}
	return cache[n]
}

func Fibonacci(n int) int {
	return fibonacci(n, nil)
}
