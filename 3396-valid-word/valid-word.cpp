class Solution {
public:
    bool isValid(string word) {
        if (word.length()<3){
            return false;
        }

        bool has_vowel = false;
        bool has_consonant = false;

        for (char c : word) {

            if(!isalnum(c)){
                return false;
            }

            if (isalpha(c)) {
                char lower_c = tolower(c);
                if (lower_c == 'a' || lower_c == 'e' || lower_c == 'i' || lower_c == 'o' || lower_c == 'u') {
                    has_vowel = true;
                } else {
                    has_consonant = true;
                }
                
            }
        }

        return has_consonant && has_vowel;        
    }
};