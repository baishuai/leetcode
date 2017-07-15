//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_287_HPP
#define LEETCODE_287_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <numeric>


using namespace std;


/**
Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive),
 prove that at least one duplicate number must exist.
 Assume that there is only one duplicate number, find the duplicate one.

Note:
You must not modify the array (assume the array is read only).
You must use only constant, O(1) extra space.
Your runtime complexity should be less than O(n2).
There is only one duplicate number in the array, but it could be repeated more than once.
  */

class Solution {
public:
    int findDuplicate(vector<int> &nums) {
        // 双指针，类似判断 linklist 是否有圆环, 太隐晦了
        // 重复值就出现在圆环开始的位置
        int fast = nums[nums[0]], slow = nums[0];
        while (fast != slow) {
            slow = nums[slow];
            fast = nums[nums[fast]];

        };
        slow = 0;
        while (fast != slow) {
            slow = nums[slow];
            fast = nums[fast];
        }
        return slow;
    }
};


#endif //LEETCODE_287_HPP
