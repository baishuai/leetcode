
#ifndef LEETCODE_242_HPP
#define LEETCODE_242_HPP

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
    bool isAnagram(string s, string t) {
        if (s.size() != t.size())
            return false;
        unordered_map<char, int> cmaps;
        for (char c : s) {
            cmaps[c]++;
        }
        for (char c: t) {
            cmaps[c]--;
        }
        return all_of(cmaps.begin(), cmaps.end(), [](const pair<char, int> &p) {
            return p.second == 0;
        });
    }
};

#endif //LEETCODE_242_HPP
