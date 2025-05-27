package climbingstairs

func climbStairs(n int) int {
    first:= 1
    second := 1

    for i:=0 ; i < n-1 ; i++ {
        temp := first
        first = first + second
        second = temp
    }

    return first
}

