
#ifndef LEETCODE_532_HPP
#define LEETCODE_532_HPP

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
    int findPairs(vector<int> &nums, int k) {
        if (k < 0)
            return 0;
        unordered_map<int, int> values;
        for (int n : nums) {
            values[n]++;
        }
        int pairs = 0;
        for (auto p : values) {
            if ((values.count(p.first + k) && k != 0) || (k == 0 && p.second > 1)) {
                pairs++;
            }
        }
        return pairs;
    }
};


#endif //LEETCODE_532_HPP
