
#ifndef LEETCODE_658_HPP
#define LEETCODE_658_HPP

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
    vector<int> findClosestElements(vector<int> &arr, int k, int x) {
        auto ite = lower_bound(arr.begin(), arr.end(), x);
        auto lite = ite, rite = ite + 1;
        k--;
        if (*ite != x) {
            if (ite != arr.begin() && *ite - x >= x - *(ite - 1)) {
                lite--;
                rite--;
            }
        }
        while (k > 0) {
            if (lite == arr.begin() || rite == arr.end())
                break;
            if (x - *(lite - 1) <= *rite - x) {
                lite--;
            } else {
                rite++;
            }
            k--;
        }
        if (lite == arr.begin()) {
            while (k-- > 0) {
                rite++;
            }
        }
        if (rite == arr.end()) {
            while (k-- > 0) {
                lite--;
            }
        }
        return vector<int>(lite, rite);
    }
};

#endif //LEETCODE_658_HPP
