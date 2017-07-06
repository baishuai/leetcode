//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_583_HPP
#define LEETCODE_583_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <iterator>
#include <unordered_map>

using namespace std;

/**
Given two words word1 and word2,
 find the minimum number of steps required to make word1 and word2 the same,
 where in each step you can delete one character in either string.

Example 1:
Input: "sea", "eat"
Output: 2
Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
Note:
The length of given words won't exceed 500.
Characters in given words can only be lower-case letters.
 */

//Edit Distance

class Solution {
public:
    int minDistance(string word1, string word2) {
        vector<int> dp(word2.size() + 1);
        vector<int> preDp(word2.size() + 1);

        for (int i = 0; i <= word2.size(); ++i) {
            dp[i] = i;
        }
        for (int j = 0; j < word1.size(); ++j) {
            dp.swap(preDp);
            dp[0] = j + 1;
            for (int k = 0; k < word2.size(); ++k) {
                if (word1[j] != word2[k]) {
                    preDp[k] += 2;
                }
                dp[k + 1] = min(min(preDp[k], dp[k] + 1), preDp[k + 1] + 1);
            }
        }
        return dp[word2.size()];
    }
};

#endif //LEETCODE_583_HPP
