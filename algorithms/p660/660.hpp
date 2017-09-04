
#ifndef LEETCODE_660_HPP
#define LEETCODE_660_HPP

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
Start from integer 1, remove any integer that contains 9 such as 9, 19, 29...

So now, you will have a new integer sequence: 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, ...

Given a positive integer n, you need to return the n-th integer after removing. Note that 1 will be the first integer.

Example 1:
Input: 9
Output: 10
Hint: n will not exceed 9 x 10^8.
 */

class Solution {
public:
    int newInteger(int n) {
        int result = 0, pos = 1;
        while (n >= 9) {
            result += n % 9 * pos;
            n /= 9;
            pos *= 10;
        }
        result += n % 9 * pos;
        return result;
    }
};

#endif //LEETCODE_660_HPP
