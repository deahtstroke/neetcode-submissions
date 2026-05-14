func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return [][]string{}
    }
    res := make(map[string][]string)
    
    for _, str := range strs {
        // sort characters to create a unique key for anagrams
        s := strings.Split(str, "")
        sort.Strings(s)
        key := strings.Join(s, "")

        if _, exists := res[key]; exists {
            res[key] = append(res[key], str)
        } else {
            res[key] = []string{str}
        }
    }

    final := [][]string{}
    for _, v := range res {
        final = append(final, v) // Correctly appends the list of anagrams
    }
    return final
}