//
// Created by bai on 17-6-30.
//

#ifndef LEETCODE_394_HPP
#define LEETCODE_394_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <numeric>
#include <stack>

using namespace std;

/**
Given an encoded string, return it's decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being
 repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces,
 square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are
 only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].

Examples:

s = "3[a]2[bc]", return "aaabcbc".
s = "3[a2[c]]", return "accaccacc".
s = "2[abc]3[cd]ef", return "abcabccdcdcdef".

 */

class Solution {

public:
    string decodeString(string s) {
        stack<int> cntStack;
        stack<string> strStack;

        int idx = 0;
        string res;
        while (idx < s.length()) {
            if (isdigit(s[idx])) {
                int cnt = 0;
                while (isdigit(s[idx])) {
                    cnt = cnt * 10 + s[idx++] - '0';
                }
                cntStack.emplace(cnt);
            } else if (s[idx] == '[') {
                strStack.emplace(res);
                res.clear();
                idx++;
            } else if (s[idx] == ']') {
                string &tmp = strStack.top();
                int repeat = cntStack.top();
                cntStack.pop();
                res.swap(tmp);
                while (repeat-- > 0) {
                    res.append(tmp);
                }
                strStack.pop();
                idx++;
            } else {
                res.push_back(s[idx++]);
            }
        }
        return res;
    }
};

#endif //LEETCODE_394_HPP
