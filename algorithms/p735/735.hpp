
#ifndef LEETCODE_735_HPP
#define LEETCODE_735_HPP

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
    vector<int> asteroidCollision(vector<int> &asteroids) {
        vector<int> remains;
        for (int a : asteroids) {
            if (remains.empty() || a > 0 || remains.back() < 0) {
                remains.emplace_back(a);
            } else if (remains.back() + a == 0) {
                remains.pop_back();
            } else if (remains.back() < -a) {
                while (!remains.empty() && remains.back() > 0 && remains.back() < -a) {
                    remains.pop_back();
                }
                if (remains.empty() || remains.back() < 0) {
                    remains.emplace_back(a);
                } else if (!remains.empty() && remains.back() == -a) {
                    remains.pop_back();
                }
            }
        }
        return remains;
    }
};

#endif //LEETCODE_735_HPP
