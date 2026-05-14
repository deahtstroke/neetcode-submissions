func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))	
	for _, num := range nums {
		m[num] = true
	}

	max := 0;
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		length := 1
		for  {
			_, ok := m[curr - 1]
			if !ok {
				break
			}
			length++
			curr -= 1
		}

		max = int(math.Max(float64(max), float64(length)))
	}

	return max
}
