
#ifndef LEETCODE_462_HPP
#define LEETCODE_462_HPP

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
Given a non-empty integer array,
 find the minimum number of moves required to make all array elements equal,
 where a move is incrementing a selected element by 1 or decrementing a selected element by 1.

You may assume the array's length is at most 10,000.

Example:

Input:
[1,2,3]

Output:
2

Explanation:
Only two moves are needed (remember each move increments or decrements one element):

[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
 */

class Solution {
public:
    /*
     * 也可以使用quick select 算法找到中位数，然后累计差值
     */
    int minMoves2(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        int sum = 0, i = 0, j = static_cast<int>(nums.size()) - 1;
        while (i < j)
            sum += nums[j--] - nums[i++];
        return sum;
    }
};

#endif //LEETCODE_462_HPP
