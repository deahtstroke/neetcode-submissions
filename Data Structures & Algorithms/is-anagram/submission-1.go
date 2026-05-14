func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    m := make(map[string]int, len(s))
    for _, str := range s {
        if _, exists := m[string(str)]; exists {
           m[string(str)] += 1 
        } else {
            m[string(str)] = 1
        }
    }

    for _, str := range t {
        if _, exists := m[string(str)]; exists {
            m[string(str)] = m[string(str)] - 1
        }
    }

    for _, val := range m {
        if val > 0 {
            return false
        }
    }
    return true
}
