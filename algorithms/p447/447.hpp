
#ifndef LEETCODE_447_HPP
#define LEETCODE_447_HPP

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
Given n points in the plane that are all pairwise distinct,
 a "boomerang" is a tuple of points (i, j, k)
 such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs.
 You may assume that n will be at most 500 and
 coordinates of points are all in the range [-10000, 10000] (inclusive).

Example:
Input:
[[0,0],[1,0],[2,0]]

Output:
2

Explanation:
The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
 */

class Solution {
public:
    int numberOfBoomerangs(vector<pair<int, int>> &points) {
        int cnt = 0;
        for (int i = 0; i < points.size(); i++) {
            unordered_map<int, int> group;
            for (int j = 0; j < points.size(); j++) {
                if (i == j)
                    continue;
                int d = dis(points[i], points[j]);
                group[d]++;
            }
            for (const auto &p : group) {
                if (p.second > 1) {
                    cnt += p.second * (p.second - 1);
                }
            }
        }
        return cnt;
    }

private:
    int dis(const pair<int, int> &p1, const pair<int, int> &p2) {
        int dx = p1.first - p2.first;
        int dy = p1.second - p2.second;
        return dx * dx + dy * dy;
    }
};

#endif //LEETCODE_447_HPP
