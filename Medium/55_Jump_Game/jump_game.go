package jumpgame

func canJump(nums []int) bool {
    goal := len(nums) - 1

    for i := goal - 1 ; i >= 0 ; i -- {
        maxJump := nums[i]
        if i + maxJump >= goal {
            goal = i
        }
    }

    return goal == 0

}


