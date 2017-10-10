
#ifndef LEETCODE_682_HPP
#define LEETCODE_682_HPP

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
    int calPoints(vector<string> &ops) {
        vector<int> points;
        for (string &s : ops) {
            if (s == "C") {
                if (!points.empty())
                    points.pop_back();
            } else if (s == "D") {
                if (!points.empty())
                    points.emplace_back(points.back() * 2);
            } else if (s == "+") {
                if (points.size() >= 2) {
                    points.emplace_back(*(points.end() - 1) + *(points.end() - 2));
                }
            } else {
                points.emplace_back(stoi(s));
            }
        }
        return accumulate(points.begin(), points.end(), 0);
    }
};

#endif //LEETCODE_682_HPP
