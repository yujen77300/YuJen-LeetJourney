package numberof1bits

func hammingWeight(n int) int {
    count := 0
    for n != 0 {
        if n & 1 == 1 {
            count += 1
        }

        n = n >> 1
    }

    return count
}