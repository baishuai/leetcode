
#ifndef LEETCODE_506_HPP
#define LEETCODE_506_HPP

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
    vector<string> findRelativeRanks(vector<int> &nums) {
        vector<int> index(nums.size());
        iota(index.begin(), index.end(), 0);
        sort(index.begin(), index.end(), [&nums](const int &l, const int &r) -> bool {
            return nums[l] > nums[r];
        });

        vector<string> result(nums.size());
        for (int i = 0; i < nums.size(); i++) {
            if (i == 0) {
                result[index[i]] = "Gold Medal";
            } else if (i == 1) {
                result[index[i]] = "Silver Medal";
            } else if (i == 2) {
                result[index[i]] = "Bronze Medal";
            } else {
                result[index[i]] = to_string(i + 1);
            }
        }

        return result;
    }
};

#endif //LEETCODE_506_HPP
