
#ifndef LEETCODE_260_HPP
#define LEETCODE_260_HPP

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
Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly twice.
 Find the two elements that appear only once.

For example:

Given nums = [1, 2, 1, 3, 2, 5], return [3, 5].

Note:
The order of the result is not important. So in the above example, [5, 3] is also correct.
Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

 */

class Solution {
public:
    vector<int> singleNumber(vector<int> &nums) {
        int aXorb = accumulate(nums.begin(), nums.end(), 0, bit_xor<int>());
        int diff = aXorb & (~(aXorb - 1));
        int na = 0, nb = 0;
        for (int n : nums) {
            if (n & diff)
                na ^= n;
            else
                nb ^= n;
        }
        return {na, nb};
    }
};

#endif //LEETCODE_260_HPP
