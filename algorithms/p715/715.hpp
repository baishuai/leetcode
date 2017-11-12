
#ifndef LEETCODE_715_HPP
#define LEETCODE_715_HPP

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
#include <map>

using namespace std;

class RangeModule {
public:

    void addRange(int left, int right) {
        auto l = intervals.upper_bound(left), r = intervals.upper_bound(right);
        if (l != intervals.begin()) {
            l--;
            if (l->second < left)
                l++;
        }
        if (l != r) {
            left = min(left, l->first);
            right = max(right, (--r)->second);
            intervals.erase(l, ++r);
        }
        intervals[left] = right;
    }

    bool queryRange(int left, int right) {
        auto ite = intervals.upper_bound(left);
        return (ite != intervals.begin() && (--ite)->second >= right);
    }

    void removeRange(int left, int right) {
        auto l = intervals.upper_bound(left), r = intervals.upper_bound(right);
        if (l != intervals.begin()) {
            l--;
            if (l->second < left)
                l++;
        }
        if (l == r)
            return;
        int l1 = min(left, l->first), r1 = max(right, (--r)->second);
        intervals.erase(l, ++r);
        if (l1 < left)
            intervals[l1] = left;
        if (r1 > right)
            intervals[right] = r1;
    }

private:

    map<int, int> intervals;
};


#endif //LEETCODE_715_HPP
