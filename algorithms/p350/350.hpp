
#ifndef LEETCODE_350_HPP
#define LEETCODE_350_HPP

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
Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].

Note:
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:
What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
 */

class Solution {
public:
    vector<int> intersect(vector<int> &nums1, vector<int> &nums2) {
        multiset<int> s1(nums1.begin(), nums1.end());
        multiset<int> s2(nums2.begin(), nums2.end());
        vector<int> interset;
        set_intersection(s1.begin(), s1.end(), s2.begin(), s2.end(), back_inserter(interset));
        return interset;
    }
};

#endif //LEETCODE_350_HPP
