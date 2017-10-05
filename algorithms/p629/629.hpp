#include <iostream>
#include <queue>
#include <algorithm>

using namespace std;

/**
 * Given two integers n and k, find how many different arrays consist of numbers
 * from 1 to n such that there are exactly k inverse pairs.

We define an inverse pair as following: For ith and jth element in the array,
 if i < j and a[i] > a[j] then it's an inverse pair; Otherwise, it's not.

Since the answer may very large, the answer should be modulo 109 + 7.
 */


//todo need review

class Solution {

    static const int mod = 1000000007;
public:
    int kInversePairs(int n, int k) {
        vector<long long> dp(n + k + 2);
        dp[0] = 1;

        for (int i = 1; i <= n; ++i) {
            for (int j = dp.size() - 1 - i; j >= 0; --j) {
                dp[j + i] -= dp[j];
                if (dp[j + i] < 0) {
                    dp[j + i] += mod;
                }
            }
        }

        for (int i = 1; i <= n; ++i) {
            for (int j = 0; j < dp.size() - 1; ++j) {
                dp[j + 1] += dp[j];
                dp[j + 1] %= mod;
            }
        }
        return int(dp[k]);
    }

    int kInversePairs2(int n, int k) {
        vector<vector<long long>> dp(n + 1, vector<long long>(k + 1));

        dp[1][0] = 1;

        fill_n(dp[1].begin(), k, 0);
        for (int i = 2; i <= n; ++i) {
            dp[i][0] = 1;
            for (int j = 1; j <= k; ++j) {
                dp[i][j] = (dp[i][j - 1] + dp[i - 1][j]);
                if (dp[i][j] > mod) {
                    dp[i][j] -= mod;
                }
                if (j >= i) {
                    dp[i][j] = (dp[i][j] - dp[i - 1][j - 1]) % mod;
                }
            }
        }
        return dp[n][k];

    }
};