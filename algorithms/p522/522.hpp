
#ifndef LEETCODE_522_HPP
#define LEETCODE_522_HPP

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

bool isSubsequence(const string &s, const string &t) {
    int si = 0;
    for (char c : t) {
        if (si < s.size() && s[si] == c)
            si++;
    }
    return si == s.size();
}

class Solution {
public:
    int findLUSlength(vector<string> &strs) {
        sort(strs.begin(), strs.end(), [](const string &lhs, const string &rhs) -> bool {
            return lhs.size() > rhs.size();
        });

        for (int i = 0; i < strs.size(); ++i) {
            auto &str = strs[i];
            bool all = true;
            for (int j = 0; j < strs.size(); ++j) {
                if (i == j)
                    continue;
                if (isSubsequence(str, strs[j])) {
                    all = false;
                    break;
                }
            }
            if (all)
                return static_cast<int>(str.size());
        }
        return -1;

    }
};

#endif //LEETCODE_522_HPP
