
#ifndef LEETCODE_593_HPP
#define LEETCODE_593_HPP

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
Given the coordinates of four points in 2D space, return whether the four points could construct a square.

The coordinate (x,y) of a point is represented by an integer array with two integers.

Example:
Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
Output: True
Note:

All the input integers are in the range [-10000, 10000].
A valid square has four equal sides with positive length and four equal angles (90-degree angles).
Input points have no order.

 */

class Solution {
public:
    bool validSquare(vector<int> &p1, vector<int> &p2, vector<int> &p3, vector<int> &p4) {
        unordered_set<int> s(
                {distance2(p1, p2), distance2(p1, p3), distance2(p1, p4), distance2(p2, p3), distance2(p2, p4),
                 distance2(p3, p4)});
        return s.count(0) == 0 && s.size() == 2;
    }

private:
    int distance2(vector<int> &p1, vector<int> &p2) {
        return (p1[0] - p2[0]) * (p1[0] - p2[0]) + (p1[1] - p2[1]) * (p1[1] - p2[1]);
    }
};

#endif //LEETCODE_593_HPP
