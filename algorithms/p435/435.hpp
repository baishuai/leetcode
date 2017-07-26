
#ifndef LEETCODE_435_HPP
#define LEETCODE_435_HPP

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
#include "../common/leetcode.hpp"

using namespace std;

/*
Given a collection of intervals,
 find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note:
You may assume the interval's end point is always bigger than its start point.
Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.
 */

class Solution {
public:
    int eraseOverlapIntervals(vector<Interval> &intervals) {
        if (intervals.size() == 0)
            return 0;
        sort(intervals.begin(), intervals.end(), [](const Interval &lhs, const Interval &rhs) -> bool {
            return lhs.end < rhs.end || ((lhs.end == rhs.end) && (lhs.start > rhs.start));
        });
        Interval &tail = intervals.front();
        int cnt = 0;
        for (int i = 1; i < intervals.size(); ++i) {
            if (intervals[i].start >= tail.end)
                tail = intervals[i];
            else
                cnt++;
        }
        return cnt;

    }
};


#endif //LEETCODE_435_HPP
