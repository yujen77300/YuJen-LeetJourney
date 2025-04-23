from collections import defaultdict

class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        count_map = defaultdict(int)
        answer = 0
        max_freq = 0
        right = 0

        for left in range(len(s)):
            while right < len(s) and (right - left - max_freq <= k) :
                count_map[s[right]] += 1
                max_freq = max(max_freq, count_map[s[right]])
                right += 1


            if right - left - max_freq > k:
                answer = max(answer, right - left - 1)
            else:
                answer = max(answer, right - left)

            count_map[s[left]] -= 1

        return answer
