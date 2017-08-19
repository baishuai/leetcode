
#ifndef LEETCODE_541_HPP
#define LEETCODE_541_HPP

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
Given a string and an integer k, you need to reverse the first k characters for every 2k characters
 counting from the start of the string.
 If there are less than k characters left, reverse all of them.
 If there are less than 2k but greater than or equal to k characters,
 then reverse the first k characters and left the other as original.

Example:
Input: s = "abcdefg", k = 2
Output: "bacdfeg"
Restrictions:
The string consists of lower English letters only.
Length of the given string and k will in the range [1, 10000]
 */

class Solution {
public:
    string reverseStr(string s, int k) {
        for (int i = 0; i < s.size(); i += 2 * k) {
            int l = i, h = min(static_cast<int >(s.size()), i + k);
            reverse(s.begin() + l, s.begin() + h);
        }
        return s;
    }
};

#endif //LEETCODE_541_HPP
