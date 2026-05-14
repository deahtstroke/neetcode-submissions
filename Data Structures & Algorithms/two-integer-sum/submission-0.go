func twoSum(nums []int, target int) []int {
   sums := make(map[int]int, len(nums))
   for i, num := range nums {
    if complimentIdx, exists := sums[num]; exists {
        return []int{complimentIdx, i}
    }
    sums[target - num] = i 
   } 
   return []int{}
}
