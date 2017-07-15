
#ifndef LEETCODE_334_HPP
#define LEETCODE_334_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>

using namespace std;

/**
Given an unsorted array return whether an increasing subsequence of length 3 exists or not in the array.

Formally the function should:
Return true if there exists i, j, k
such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
Your algorithm should run in O(n) time complexity and O(1) space complexity.

Examples:
Given [1, 2, 3, 4, 5],
return true.

Given [5, 4, 3, 2, 1],
return false.
 */

class Solution {
public:
    bool increasingTriplet(vector<int> &nums) {
        vector<int> inc;
        for (auto n: nums) {
            auto f = lower_bound(inc.begin(), inc.end(), n);
            if (f == inc.end()) {
                inc.emplace_back(n);
                if (inc.size() == 3) {
                    break;
                }
            } else {
                *f = n;
            }
        }
        return inc.size() == 3;
    }
};

#endif //LEETCODE_334_HPP
