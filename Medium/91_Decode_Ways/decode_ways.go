package decodeways

func numDecodings(s string) int {
    dp := make(map[int]int)

    var dfs func(i int) int
    dfs = func(i int) int {
        if val, ok := dp[i]; ok {
            return val
        }

        if i == len(s) {
            return 1
        }

        if s[i] == '0' {
            return 0
        }

        res := dfs(i + 1)

        if i+1 < len(s) {
            if s[i] == '1' || (s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6') {
                res += dfs(i + 2)
            }
        }

        dp[i] = res
        return res
    }

    return dfs(0)
}

