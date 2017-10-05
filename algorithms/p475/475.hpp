
#ifndef LEETCODE_475_HPP
#define LEETCODE_475_HPP

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
    int findRadius(vector<int> &houses, vector<int> &heaters) {
        int radius = 0;
        sort(heaters.begin(), heaters.end());
        for (auto house : houses) {
            auto iter = lower_bound(heaters.begin(), heaters.end(), house);
            if (iter == heaters.end()) {
                radius = max(radius, house - heaters.back());
            } else if (iter == heaters.begin()) {
                radius = max(radius, heaters.front() - house);
            } else {
                int r = min(*iter - house, house - *(iter - 1));
                radius = max(r, radius);
            }
        }
        return radius;
    }
};

#endif //LEETCODE_475_HPP
