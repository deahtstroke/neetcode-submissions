class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        largest = 0
        map = {}
        l = 0
        for i, c in enumerate(s):
            if c not in map:
                map[c] = i
            else:
                # increment left pointer until the last occurrence
                # of current char
                l = map[c]+1
                map = {k: v for k, v in map.items() if v >= l}

                map[c] = i
            largest = max(largest, len(map))

        return largest