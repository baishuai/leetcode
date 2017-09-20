
#ifndef LEETCODE_665_HPP
#define LEETCODE_665_HPP

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
    bool checkPossibility(vector<int> &nums) {
        int cnt = 0;
        for (int i = 1; i < nums.size() && cnt <= 1; ++i) {
            if (nums[i - 1] > nums[i]) {
                cnt++;
                if (i < 2 || nums[i - 2] <= nums[i])
                    nums[i - 1] = nums[i];
                else
                    nums[i] = nums[i - 1];
            }
        }
        return cnt <= 1;
    }
};


#endif //LEETCODE_665_HPP
