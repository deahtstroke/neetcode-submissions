class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        largest = 0
        map = {}
        l = 0
        for i, c in enumerate(s):
            if c in map and map[c] >= l:
                l = map[c] + 1
            # increment left pointer until the last occurrence
            # of current char
            map[c] = i
            largest = max(largest, i - l + 1)

        return largest