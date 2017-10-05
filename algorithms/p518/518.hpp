
#ifndef LEETCODE_518_HPP
#define LEETCODE_518_HPP

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

/*
You are given coins of different denominations and a total amount of money. Write a function to compute the number of combinations that make up that amount. You may assume that you have infinite number of each kind of coin.

Note: You can assume that

0 <= amount <= 5000
1 <= coin <= 5000
the number of coins is less than 500
the answer is guaranteed to fit into signed 32-bit integer
 */

class Solution {
public:
    int change(int amount, vector<int> &coins) {
        sort(coins.begin(), coins.end());
        vector<int> dp(static_cast<unsigned long>(amount + 1));
        dp[0] = 1;
        for (int j : coins) {
            for (int i = j; i <= amount; ++i) {
                dp[i] += dp[i - j];
            }
        }
        return dp[amount];
    }
};

#endif //LEETCODE_518_HPP
