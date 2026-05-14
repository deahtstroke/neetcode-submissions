# Two pointer:
# Have one pointer start at s[0] and the other at s[s.length-1]
# move them until left > right and check to see if letters match
class Solution:
    def isPalindrome(self, s: str) -> bool:
        # Declare both pointers
        left, right = 0, len(s) - 1
        while left < right:
            # If current char isn't alphanumeric keep iterating
            while left < right and not s[left].isalnum():
                left += 1

            # Same here
            while right > left and not s[right].isalnum():
                right -= 1

            if s[left].lower() != s[right].lower():
                return False

            left += 1
            right -= 1

        return True
