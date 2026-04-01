func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var freq [26]int

    for _, ch := range s {
        freq[ch - 'a']++;
    }

    for _, ch := range t {
        freq[ch - 'a']--;
    }

    for _, count := range freq {
        if count != 0 {
            return false
        }
    }

    return true
}
