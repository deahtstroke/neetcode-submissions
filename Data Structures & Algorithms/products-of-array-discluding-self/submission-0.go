func productExceptSelf(nums []int) []int {
	total := 1	
	res := []int{}
	for i, _ := range nums {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			total *= nums[j]
		}	
		res = append(res, total)
		total = 1
	}
	return res	
}