func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := [26]int{}
    for i := 0; i < len(s); i++ {
        count[s[i] - 'a']++
        count[t[i] - 'a']--
    }

    for _, num := range count {
        if num != 0 {
            return false
        }
    }

    return true
}
