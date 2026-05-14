func topKFrequent(nums []int, k int) []int {
    // split all numbers into buckets that are essentially just map entries
    // or even better, we don't need to keep track of individual numbers
    // just the frequency of each element found in the array
    buckets := make(map[int]int)

    for _, num := range nums {
        buckets[num]++
    }

    idx := make([][]int, len(nums) + 1)
    for num, count := range buckets {
       idx[count] = append(idx[count], num)
    }

    res := []int{}
    for i := len(idx) - 1; i > 0; i-- {
        for _, num := range idx[i] {
            res = append(res, num)
            if len(res) == k {
                return res
            }
        }
    }

    return res
}
