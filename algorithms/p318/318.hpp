
#ifndef LEETCODE_318_HPP
#define LEETCODE_318_HPP

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
#include <iomanip>

using namespace std;

class Solution {
public:
    //bit manipulation
    int maxProduct(vector<string> &words) {
        vector<int> masks(words.size());
        int result = 0;
        for (int i = 0; i < words.size(); ++i) {
            for (char c: words[i]) {
                masks[i] |= 1 << (c - 'a');
            }
            for (int j = 0; j < i; ++j) {
                if ((masks[i] & masks[j]) == 0) {
                    result = max(result, static_cast<int> (words[i].size() * words[j].size()));
                }
            }
        }
        return result;
    }
};

#endif //LEETCODE_318_HPP
