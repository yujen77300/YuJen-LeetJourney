# [Two Sum](https://leetcode.com/problems/two-sum/)

## Solution 1: Brute Force with Nested Loops
- **Time Complexity**: O(n²)
- **Space Complexity**: O(1)
- **Approach**: Use two nested loops to check every possible pair of numbers in the array and verify if their sum equals the target value.



## Solution 2: Hash Map (Dictionary)
- **Time Complexity**:: O(n)
- **Space Complexity**: O(n)
- **Approach**: Use a hash map to store visited numbers and their indices. For each number, check if the complement (target - current number) exists in the hash map.