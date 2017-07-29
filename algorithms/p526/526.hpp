
#ifndef LEETCODE_526_HPP
#define LEETCODE_526_HPP

#include <iostream>
#include <algorithm>
#include <vector>
#include <numeric>

using namespace std;

/*
Suppose you have N integers from 1 to N. We define a beautiful arrangement as an array
 that is constructed by these N numbers successfully
 if one of the following is true for the ith position (1 <= i <= N) in this array:

The number at the ith position is divisible by i.
i is divisible by the number at the ith position.
Now given N, how many beautiful arrangements can you construct?

Example 1:
Input: 2
Output: 2
Explanation:

The first beautiful arrangement is [1, 2]:

Number at the 1st position (i=1) is 1, and 1 is divisible by i (i=1).

Number at the 2nd position (i=2) is 2, and 2 is divisible by i (i=2).

The second beautiful arrangement is [2, 1]:

Number at the 1st position (i=1) is 2, and 2 is divisible by i (i=1).

Number at the 2nd position (i=2) is 1, and i (i=2) is divisible by 1.
Note:
N is a positive integer and will not exceed 15.
 */


class Solution {
public:
    int countArrangement(int N) {
        vs.resize((unsigned long) (N + 1));
        iota(vs.begin(), vs.end(), 0);
        return counts(N);
    }

private:

    int counts(int n) {
        if (n <= 0)
            return 1;
        int ans = 0;
        for (int i = 1; i <= n; ++i) {
            if (vs[i] % n == 0 || n % vs[i] == 0) {
                swap(vs[i], vs[n]);
                ans += counts(n - 1);
                swap(vs[i], vs[n]);
            }
        }
        return ans;
    }

    vector<int> vs;

};

#endif //LEETCODE_526_HPP
