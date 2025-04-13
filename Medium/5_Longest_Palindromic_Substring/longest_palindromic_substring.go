package longestpalindromicsubstring


func longestPalindrome(s string) string {
    result := ""
    maxLength := 0

    for i, _ := range s {
        // odd length
        l := i
        r := i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if (r-l+1) > maxLength{
                result = s[l:r+1]
                maxLength = r-l+1
            }

            l -= 1
            r += 1
        }

        // even length
        l = i
        r = i + 1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if (r-l+1) > maxLength{
                result = s[l:r+1]
                maxLength = r-l+1
            }

            l -= 1
            r += 1
        }
    } 

    return result
}