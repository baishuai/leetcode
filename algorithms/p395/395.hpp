//
// Created by bai on 17-6-25.
//

#ifndef LEETCODE_395_HPP
#define LEETCODE_395_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <unordered_map>
#include <stack>

using namespace std;


class Solution {

public:

    int longestSubstring(string s, int k) {
        unordered_map<char, int> mp;
        for_each(s.begin(), s.end(), [&mp](char c) {
            mp[c]++;
        });
        auto f = find_if(mp.begin(), mp.end(), [k](pair<const char, int> &p) {
            return p.second < k;
        });
        if (f == mp.end()) {
            return (int) s.size();
        }
        auto c = (*f).first;

        int mlen = 0;
        unsigned long split = 0;
        auto p = s.find(c, split);
        while (p != string::npos) {
            mlen = max(mlen, longestSubstring(s.substr(split, p), k));
            split = p + 1;

            while (p <= split && p != string::npos) {
                split = p + 1;
                p = s.find(c, split);
            }
        }
        mlen = max(mlen, longestSubstring(s.substr(split), k));
        return mlen;
    }
};


#endif //LEETCODE_395_HPP
