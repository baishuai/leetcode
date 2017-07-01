//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_367_HPP
#define LEETCODE_367_HPP


#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>

using namespace std;

/**
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

Input: 16
Returns: True
Example 2:

Input: 14
Returns: False
 */

class Solution {
public:
    bool isPerfectSquare(int num) {
        if (num < 1) {
            return false;
        }
        long r = num;
        while (r * r > num) {
            r = (r + num / r) / 2;
        }
        return r * r == num;
    }
};


#endif //LEETCODE_367_HPP
