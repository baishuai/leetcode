
#ifndef LEETCODE_452_HPP
#define LEETCODE_452_HPP

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

/*
There are a number of spherical balloons spread in two-dimensional space.
 For each balloon, provided input is the start and end coordinates of the horizontal diameter.
 Since it's horizontal, y-coordinates don't matter and hence
 the x-coordinates of start and end of the diameter suffice.
 Start is always smaller than end. There will be at most 104 balloons.

An arrow can be shot up exactly vertically from different points along the x-axis.
 A balloon with xstart and xend bursts by an arrow shot at x if xstart ≤ x ≤ xend.
 There is no limit to the number of arrows that can be shot.
 An arrow once shot keeps travelling up infinitely.
 The problem is to find the minimum number of arrows that must be shot to burst all balloons.
 */

class Solution {
public:
    int findMinArrowShots(vector<pair<int, int>> &points) {
        sort(points.begin(), points.end(), [](const pair<int, int> &l, const pair<int, int> &r) -> bool {
            return l.second < r.second || (l.second == r.second && l.first < r.first);
        });
        int arrors = 0, cur = 0;
        while (cur < points.size()) {
            auto v = points[cur].second;
            arrors++;
            while (cur < points.size() && points[cur].first <= v)
                cur++;
        }
        return arrors;
    }
};

#endif //LEETCODE_452_HPP
