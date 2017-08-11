
#ifndef LEETCODE_524_HPP
#define LEETCODE_524_HPP

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
    string findLongestWord(string s, vector<string> &d) {
        sort(d.begin(), d.end(), [](const string &l, const string &r) -> bool {
            return l.size() > r.size() || (l.size() == r.size() && l < r);
        });

        for (const string &word : d) {
            int i = 0, j = 0;
            for (; i < s.size() && j < word.size(); ++i) {
                if (word[j] == s[i])
                    j++;
            }
            if (j == word.size()) {
                return word;
            }
        }
        return "";
    }
};

#endif //LEETCODE_524_HPP
