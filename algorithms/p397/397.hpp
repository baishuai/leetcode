
#ifndef LEETCODE_397_HPP
#define LEETCODE_397_HPP

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
Given a positive integer n and you can do operations as follow:

If n is even, replace n with n/2.
If n is odd, you can replace n with either n + 1 or n - 1.
What is the minimum number of replacements needed for n to become 1?

Example 1:

Input:
8

Output:
3

Explanation:
8 -> 4 -> 2 -> 1
Example 2:

Input:
7

Output:
4

Explanation:
7 -> 8 -> 4 -> 2 -> 1
or
7 -> 6 -> 3 -> 2 -> 1

 */

class Solution {
public:
    int integerReplacement(int n) {
        int cnt = 0;
        uint un = static_cast<uint>(n);
        while (un != 1) {
            if ((un & 1) == 0)
                un >>= 1;
            else if (un != 3 && ((un >> 1) & 1) == 1)
                un++;
            else
                un--;
            cnt++;
        }
        return cnt;
    }
};

#endif //LEETCODE_397_HPP

