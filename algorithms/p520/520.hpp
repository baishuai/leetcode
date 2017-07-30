
#ifndef LEETCODE_520_HPP
#define LEETCODE_520_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include <stack>
#include <string>

using namespace std;

class Solution {
public:
    bool detectCapitalUse(string word) {
        if (word.size() <= 1)
            return true;
        if (word[0] > 'Z')
            return word[1] > 'Z' && allCapletters(word, false);
        else
            return allCapletters(word, word[1] <= 'Z');
    }

private:
    bool allCapletters(string &word, bool flag) {
        for (int i = 2; i < word.size(); ++i) {
            if ((word[i] <= 'Z') != flag)
                return false;
        }
        return true;
    }
};

#endif //LEETCODE_520_HPP
