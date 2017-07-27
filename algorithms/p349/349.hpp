
#ifndef LEETCODE_349_HPP
#define LEETCODE_349_HPP

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
Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

Note:
Each element in the result must be unique.
The result can be in any order.
 */

class Solution {
public:
    vector<int> intersection(vector<int> &nums1, vector<int> &nums2) {
        set<int> s1(nums1.begin(), nums1.end());
        set<int> s2(nums2.begin(), nums2.end());
        vector<int> interset;
        set_intersection(s1.begin(), s1.end(), s2.begin(), s2.end(), back_inserter(interset));
        return interset;
    }
};

#endif //LEETCODE_349_HPP
