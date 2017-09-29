
#ifndef LEETCODE_500_HPP
#define LEETCODE_500_HPP

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
    vector<string> findWords(vector<string> &words) {
        unordered_map<char, int8_t> c2row = {
                {'q', 1},
                {'w', 1},
                {'e', 1},
                {'r', 1},
                {'t', 1},
                {'y', 1},
                {'u', 1},
                {'i', 1},
                {'o', 1},
                {'p', 1},
                {'a', 2},
                {'s', 2},
                {'d', 2},
                {'f', 2},
                {'g', 2},
                {'h', 2},
                {'j', 2},
                {'k', 2},
                {'l', 2},
                {'z', 3},
                {'x', 3},
                {'c', 3},
                {'v', 3},
                {'b', 3},
                {'n', 3},
                {'m', 3}
        };

        vector<string> res;
        for (auto &w :  words) {
            unordered_set<int8_t> chars;
            for (char c : w) {
                chars.insert(c2row[tolower(c)]);
            }
            if (chars.size() == 1)
                res.emplace_back(w);
        }
        return res;
    }
};

#endif //LEETCODE_500_HPP
