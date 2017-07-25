
#ifndef LEETCODE_554_HPP
#define LEETCODE_554_HPP

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
There is a brick wall in front of you. The wall is rectangular and has several rows of bricks.
 The bricks have the same height but different width.
 You want to draw a vertical line from the top to the bottom and cross the least bricks.

The brick wall is represented by a list of rows.
 Each row is a list of integers representing the width of each brick in this row from left to right.

If your line go through the edge of a brick,
 then the brick is not considered as crossed.
 You need to find out how to draw the line to cross the least bricks and return the number of crossed bricks.

You cannot draw a line just along one of the two vertical edges of the wall,
 in which case the line will obviously cross no bricks.

 */

class Solution {
public:
    int leastBricks(vector<vector<int>> &wall) {
        unordered_map<int, int> gaps;
        for (const auto &row : wall) {
            int length = 0;
            for (int i=0; i < row.size()-1;i++) {
                length += row[i];
                gaps[length]++;
            }
        }
        int maxGaps = 0;
        for (const auto &p : gaps) {
            maxGaps = max(maxGaps, p.second);
        }
        return wall.size() - maxGaps;
    }
};


#endif //LEETCODE_554_HPP
