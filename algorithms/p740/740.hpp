
#ifndef LEETCODE_740_HPP
#define LEETCODE_740_HPP

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
    int deleteAndEarn(vector<int> &nums) {
        unordered_map<int, int> values;
        for (int num: nums) {
            values[num] += num;
        }

        int minus2 = 0, minus1 = 0;
        for (int i = 0; i < 10001; ++i) {
            int m = max(minus1, minus2+values[i]);
            minus2 = minus1;
            minus1 = m;
        }
        return max(minus1, minus2);
    }
};


#endif //LEETCODE_740_HPP
