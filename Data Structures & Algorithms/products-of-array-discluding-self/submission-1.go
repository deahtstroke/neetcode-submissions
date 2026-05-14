func productExceptSelf(nums []int) []int {
	// Compute them on the go
	n := len(nums)
	prefix := make([]int, n)
	postfix := make([]int, n)
	res := make([]int, n)

	prefix[0] = 1
	for i := 1; i < n; i++ {
		// prefix[0] = 1 = 1
		// prefix[1] = 1 * 4 = 4
		// prefix[2] = 4 * 2 = 8
		// prefix[3] = 8 * 4 = 32
		// prefix = [1, 4, 8, 32]
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	postfix[n-1] = 1
	for i := n-2; i >= 0; i-- {
		// postfix[0] = 24 * 2 = 48
		// postfix[1] = 6 * 4 = 24
		// postfix[2] = 1 * 6 = 6
		// postfix[3] = 1 = 1
		// postfix = [48, 24, 6, 1]
		postfix[i] = postfix[i + 1] * nums[i + 1]
	}

	for i := 0; i < n; i++ {
		res[i] = prefix[i] * postfix[i]	
	}

	return res
}