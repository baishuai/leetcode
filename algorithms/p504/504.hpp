
#ifndef LEETCODE_504_HPP
#define LEETCODE_504_HPP

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
Given an integer, return its base 7 string representation.

Example 1:
Input: 100
Output: "202"
Example 2:
Input: -7
Output: "-10"
Note: The input will be in range of [-1e7, 1e7].
 */


class Solution {
public:
    string convertToBase7(int num) {
        bool negtive = num < 0;
        num = abs(num);
        int b7 = 0, pos = 1;
        while (num > 0) {
            b7 += pos * (num % 7);
            pos *= 10;
            num /= 7;
        }
        string s = negtive ? "-" : "";
        s.append(to_string(b7));
        return s;
    }
};

#endif //LEETCODE_504_HPP
