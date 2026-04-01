class Solution {
    public boolean isAnagram(String s, String t) {
        if(s.length() != t.length()) {
            return false;
        }

        var frequency = new int[26];
        for(char ch : s.toCharArray()) {
            frequency[ch - 'a']++;
        }

        for(char ch : t.toCharArray()) {
            frequency[ch - 'a']--;
        }

        for(int pos : frequency) {
            if(pos != 0) return false;
        }
        return true;
    }
}
