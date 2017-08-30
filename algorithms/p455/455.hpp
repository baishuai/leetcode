
#ifndef LEETCODE_455_HPP
#define LEETCODE_455_HPP

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
    int findContentChildren(vector<int> &g, vector<int> &s) {
        sort(g.begin(), g.end());
        sort(s.begin(), s.end());
        int gi = 0, si = 0, cnt = 0;
        while (gi < g.size() && si < s.size()) {
            while (si < s.size() && s[si] < g[gi])
                si++;
            if (si < s.size()) {
                gi++;
                si++;
                cnt++;
            }
        }
        return cnt;
    }
};


#endif //LEETCODE_455_HPP
