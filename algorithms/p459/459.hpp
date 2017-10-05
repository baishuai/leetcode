
#ifndef LEETCODE_459_HPP
#define LEETCODE_459_HPP

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

/*
Given a non-empty string check if it can be constructed by taking a substring of it
 and appending multiple copies of the substring together.
 You may assume the given string consists of lowercase English letters
 only and its length will not exceed 10000.
 */

class Solution {
public:
    bool repeatedSubstringPattern(string s) {
        auto ss = s + s;
        return ss.substr(1, (s.size() << 1) - 2).find(s) != string::npos;
    }
};

#endif //LEETCODE_459_HPP
