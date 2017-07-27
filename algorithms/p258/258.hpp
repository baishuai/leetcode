
#ifndef LEETCODE_258_HPP
#define LEETCODE_258_HPP

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
Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

For example:

Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.

Follow up:
Could you do it without any loop/recursion in O(1) runtime?
 */

class Solution {
public:
    int addDigits(int num) {
        return num % 9 == 0 && num > 0 ? 9 : num % 9;
    }
};

#endif //LEETCODE_258_HPP
