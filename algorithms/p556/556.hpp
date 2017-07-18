
#ifndef LEETCODE_556_HPP
#define LEETCODE_556_HPP

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
Given a positive 32-bit integer n,
 you need to find the smallest 32-bit integer which has exactly the same digits existing
 in the integer n and is greater in value than n.
 If no such positive 32-bit integer exists, you need to return -1.

Example 1:
Input: 12
Output: 21
Example 2:
Input: 21
Output: -1

 */

class Solution {
public:
    int nextGreaterElement(int n) {
        auto digits = to_string(n);
        next_permutation(digits.begin(), digits.end());
        auto res = stoll(digits);
        return (res > numeric_limits<int>::max() || res <= n) ? -1 : static_cast<int>(res);
    }
};

#endif //LEETCODE_556_HPP
