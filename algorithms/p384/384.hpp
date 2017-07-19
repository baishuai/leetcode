
#ifndef LEETCODE_384_HPP
#define LEETCODE_384_HPP

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

//Shuffle a set of numbers without duplicates.

class Solution {
public:
    Solution(vector<int> nums) : nums(nums) {}

    /** Resets the array to its original configuration and return it. */
    vector<int> reset() {
        return nums;
    }

    /** Returns a random shuffling of the array. */
    vector<int> shuffle() {
        vector<int> shuf(nums);
        for (int i = static_cast<int>(shuf.size() - 1); i > 0; --i) {
            std::swap(shuf[i], shuf[rand() % (i + 1)]);
        }
        return shuf;
    }

private:
    vector<int> nums;
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution obj = new Solution(nums);
 * vector<int> param_1 = obj.reset();
 * vector<int> param_2 = obj.shuffle();
 */

#endif //LEETCODE_384_HPP
