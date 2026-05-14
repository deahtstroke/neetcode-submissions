func twoSum(numbers []int, target int) []int {
    // Slow pointer initialized in loop-signature
    // Fast pointer initialized inside the loop
    for slow := 0; slow < len(numbers)-1; slow++ {
        fast := slow + 1
        for fast < len(numbers) {
            if numbers[slow]+numbers[fast] == target {
                return []int{slow + 1, fast + 1}
            }
            fast++
        }
    }
    return []int{}
}