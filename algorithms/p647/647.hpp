
#ifndef LEETCODE_647_HPP
#define LEETCODE_647_HPP

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
    int countSubstrings(string s) {
        int cnt = 0;
        for (int i = 0; i < s.size(); ++i) {
            cnt += countExpand(s, i, i);
            cnt += countExpand(s, i, i + 1);
        }
        return cnt;
    }

private:
    int countExpand(string &s, int i, int j) {
        int cnt = 0;
        while (i >= 0 && j < s.size() && s[i] == s[j]) {
            cnt++;
            i--;
            j++;
        }
        return cnt;
    }
};

#endif //LEETCODE_647_HPP
