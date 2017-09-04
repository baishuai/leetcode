
#ifndef LEETCODE_216_HPP
#define LEETCODE_216_HPP

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
Find all possible combinations of k numbers that add up to a number n,
 given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.


Example 1:

Input: k = 3, n = 7

Output:

[[1,2,4]]

Example 2:

Input: k = 3, n = 9

Output:

[[1,2,6], [1,3,5], [2,3,4]]
 */

class Solution {
public:
    vector<vector<int>> combinationSum3(int k, int n) {
        vector<vector<int>> res;
        vector<int> one_case;
        sumTail(res, one_case, k, 1, n);
        return res;
    }

private:
    void sumTail(vector<vector<int>> &res, vector<int> &one_case, int k, int cur, int tail) {
        if (tail <= 0) {
            if (tail == 0 && one_case.size() == k)
                res.emplace_back(vector<int>(one_case.begin(), one_case.end()));
            return;
        }
        if (one_case.size() >= k)
            return;
        for (int i = cur; i <= 9; ++i) {
            one_case.push_back(i);
            sumTail(res, one_case, k, i + 1, tail - i);
            one_case.pop_back();
        }
    }

};

#endif //LEETCODE_216_HPP
