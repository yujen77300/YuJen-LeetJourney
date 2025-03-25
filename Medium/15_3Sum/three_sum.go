package threesum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	//小到大排列
	// for i := 0; i < len(nums)-1; i++ {
	//     for j := i + 1; j < len(nums); j++ {
	//         if nums[i] > nums[j] {
	//             nums[i], nums[j] = nums[j], nums[i]
	//         }
	//     }
	// }
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		// 跳過重複的元素，確保不會有重複的 triplets
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

				// 跳過重複的 l 值，避免相同的 triplets
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
