class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
       # division solution: get the product of the entire array
       # [1, 2, 4, 6] 
       # total = 1 * 2 * 4 * 6 = 48
       # [48, 24, 12, 8]
       # for i in nums:
       #    position at i = total / nums[i]
       # BUT! What happens if num[i] = 0?
       # We cannot divide by 0
       # V2: If we have only one zero, we're good!
       # else if we have more than one zero, the entire array
       # will be zero
       #
       # Runtime of getting total = O(n) + O(n) = 2O(n) = O(n)
        zeroes = []
        for i, num in enumerate(nums):
            if len(zeroes) > 1:
                return [0] * len(nums)

            if num == 0:
                zeroes.append(i)

        res = []
        total = 1
        if len(zeroes) == 0 :
            for i, num in enumerate(nums):
                total *= num
    
            for i, num in enumerate(nums):
                if num == 0:
                    res.append(0)
                else:
                    res.append(int(total/num))
        else:
            for i, num in enumerate(nums):
                if i == zeroes[0]:
                    continue
                total *= num
            
            for i, num in enumerate(nums):
                if i == zeroes[0]:
                    res.append(total)
                else:
                    res.append(0)
        return res