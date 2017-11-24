
#ifndef LEETCODE_383_HPP
#define LEETCODE_383_HPP

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
    bool canConstruct(string ransomNote, string magazine) {
        unordered_map<char, int> chars;
        for (char c: magazine)
            chars[c]++;
        for (char c: ransomNote)
            chars[c]--;
        return all_of(chars.begin(), chars.end(), [](const pair<char, int> &p) -> bool {
            return p.second >= 0;
        });
    }
};


#endif //LEETCODE_383_HPP
