
#ifndef LEETCODE_646_HPP
#define LEETCODE_646_HPP

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
    int findLongestChain(vector<vector<int>> &pairs) {
        sort(pairs.begin(), pairs.end(), [](const vector<int> &l, const vector<int> &r) -> bool {
            return l[0] < r[0] || (l[0] == r[0] && l[1] < r[1]);
        });

        vector<vector<int>> chain;
        for (vector<int> &p : pairs) {
            auto iter = lower_bound(chain.begin(), chain.end(), p,
                                    [](const vector<int> &l, const vector<int> &r) -> bool {
                                        return l[1] <= r[1];
                                    });
            if (iter == chain.end()) {
                if (chain.size() == 0 || chain.back()[1] < p[0])
                    chain.push_back(p);
            } else if ((*iter)[0] < p[0]) {
                iter->swap(p);
            }
        }
        return chain.size();
    }
};

#endif //LEETCODE_646_HPP
