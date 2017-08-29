
#ifndef LEETCODE_482_HPP
#define LEETCODE_482_HPP

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
    string licenseKeyFormatting(string S, int K) {
        string rev;
        int k = 0;
        for (auto it = S.rbegin(); it != S.rend(); it++) {
            if (*it == '-')
                continue;
            if (k == K) {
                rev.push_back('-');
                k = 0;
            }
            k++;
            rev.push_back(static_cast<char>(toupper(*it)));
        }
        return string(rev.rbegin(), rev.rend());
    }
};

#endif //LEETCODE_482_HPP
