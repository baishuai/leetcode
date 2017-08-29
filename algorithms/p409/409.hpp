
#ifndef LEETCODE_409_HPP
#define LEETCODE_409_HPP

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
Given a string which consists of lowercase or uppercase letters,
 find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.
 */

class Solution {
public:
    int longestPalindrome(string s) {
        unordered_map<char, int> chars;
        for (char c : s) {
            chars[c]++;
        }
        int len = 0;
        bool single = false;
        for (const auto &p : chars) {
            len += p.second - (p.second & 1);
            single = single || ((p.second & 1) == 1);
        }
        return len + (single ? 1 : 0);
    }
};

#endif //LEETCODE_409_HPP
