
#ifndef LEETCODE_645_HPP
#define LEETCODE_645_HPP

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

class Solution {
public:
    vector<int> findErrorNums(vector<int> &nums) {
        vector<int> res;
        for (int n : nums) {
            if (nums[abs(n) - 1] < 0) {
                res.push_back(abs(n));
            } else {
                nums[abs(n) - 1] *= -1;
            }
        }
        for (int i = 1; i <= nums.size(); ++i) {
            if (nums[i - 1] > 0) {
                res.push_back(i);
                break;
            }
        }
        return res;
    }
};


#endif //LEETCODE_645_HPP
