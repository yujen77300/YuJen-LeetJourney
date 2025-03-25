# [3sum](https://leetcode.com/problems/3sum/)

## Solution 1: Three Loops
- **Time Complexity**: O(n³) - where n is the length of the array
- **Space Complexity**: O(n) - for the result array
- **Approach**:
    1. Use three nested loops to consider every possible triplet
    2. Check if each triplet sums to zero
    3. De-duplicate results to avoid repeating the same triplet

```go
func threeSum(nums []int) [][]int {
    var result [][]int
    n := len(nums)

    sort.Ints(nums)
    for i := 0; i < n-2; i++ {
        // Skip duplicates
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        for j := i+1; j < n-1; j++ {
            // Skip duplicates
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            for k := j+1; k < n; k++ {
                // Skip duplicates
                if k > j+1 && nums[k] == nums[k-1] {
                    continue
                }
                if nums[i] + nums[j] + nums[k] == 0 {
                    result = append(result, []int{nums[i], nums[j], nums[k]})
                }
            }
        }
    }

    return result
}
```

## Solution 2: Improve solution 1 with Hash Map
- **Time Complexity**: O(n²) - where n is the length of the array
- **Space Complexity**: O(1) - excluding the output array
- **Approach**:
    1. Sort the array to help with deduplication
    2. For each element, use a hash map to find pairs that sum to its negative
    3. Keep track of used values to avoid duplicate triplets
    4. Reduces one loop compared to the brute force approach

```go
func threeSum(nums []int) [][]int {
    var result [][]int
    n := len(nums)
    sort.Ints(nums)
    
    for i := 0; i < n-2; i++ {
        // Skip duplicates
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        target := -nums[i]
        seen := make(map[int]bool)
        
        for j := i+1; j < n; j++ {
            complement := target - nums[j]
            
            if seen[complement] {
                result = append(result, []int{nums[i], complement, nums[j]})
                
                // Skip duplicates
                for j+1 < n && nums[j] == nums[j+1] {
                    j++
                }
            }
            seen[nums[j]] = true
        }
    }
    return result
}
```


## Solution 3: Two Pointers
- **Time Complexity**: O(n²) - O(n log n) for sorting + O(n²) for the algorithm
- **Space Complexity**: O(n) - for the result array (excluding the space for sorting)
- **Approach**:
  1. Sort the array to enable the two-pointer technique
  2. For each element, use two pointers to find pairs that sum to its negative
  3. Skip duplicate elements to avoid redundant triplets
  4. Move pointers based on sum comparison to efficiently search the array


```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := nums[i] + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
```