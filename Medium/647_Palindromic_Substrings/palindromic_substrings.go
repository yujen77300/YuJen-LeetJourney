package palindromicsubstrings

func countSubstrings(s string) int {
    answer := 0

    expand := func(l, r int) {
        for l >= 0 && r <= len(s) - 1 && s[l] == s[r] {
            answer ++
            l--
            r ++
        }
    }

    for i := 0 ; i <= len(s)-1 ; i++ {
        expand(i,i)
        expand(i,i+1)
    }

    return answer

}

