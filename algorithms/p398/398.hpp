
#ifndef LEETCODE_398_HPP
#define LEETCODE_398_HPP

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
Given an array of integers with possible duplicates, randomly output the index of a given target number. You can assume that the given target number must exist in the array.

Note:
The array size can be very large. Solution that uses too much extra space will not pass the judge.

Example:

int[] nums = new int[] {1,2,3,3,3};
Solution solution = new Solution(nums);

// pick(3) should return either index 2, 3, or 4 randomly. Each index should have equal probability of returning.
solution.pick(3);

// pick(1) should return 0. Since in the array only nums[0] is equal to 1.
solution.pick(1);
 */

class Solution {
public:
    Solution(vector<int> nums) : nums(nums) {}

    int pick(int target) {
        int result = -1;
        int count = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] != target)
                continue;
            count++;
            if (rand() % count == 0)
                result = i;
        }
        return result;
    }

private:
    const vector<int> nums;
};


#endif //LEETCODE_398_HPP
