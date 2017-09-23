
#ifndef LEETCODE_495_HPP
#define LEETCODE_495_HPP

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
    int findPoisonedDuration(vector<int> &timeSeries, int duration) {
        if (timeSeries.size() == 0)
            return 0;
        int total = 0;
        for (int i = 0; i < timeSeries.size() - 1; ++i) {
            total += (timeSeries[i] + duration <= timeSeries[i + 1]) ? duration : timeSeries[i + 1] - timeSeries[i];
        }
        return total + duration;
    }
};


#endif //LEETCODE_495_HPP
