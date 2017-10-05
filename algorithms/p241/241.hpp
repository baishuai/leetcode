
#ifndef LEETCODE_241_HPP
#define LEETCODE_241_HPP

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
Given a string of numbers and operators,
 return all possible results from computing all the different possible ways to group numbers and operators.
 The valid operators are +, - and *.


Example 1
Input: "2-1-1".

((2-1)-1) = 0
(2-(1-1)) = 2
Output: [0, 2]


Example 2
Input: "2*3-4*5"

(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
Output: [-34, -14, -10, -10, 10]
 */

class Solution {
public:
    vector<int> diffWaysToCompute(string input) {
        unordered_map<string, vector<int>> memo;
        return diffWaysWithMemo(memo, input);
    }

private:
    vector<int> diffWaysWithMemo(unordered_map<string, vector<int>> &memo, string input) {
        if (memo.count(input))
            return memo[input];
        vector<int> result;
        for (int i = 0; i < input.size(); ++i) {
            char cur = input[i];
            if (cur == '+' || cur == '-' || cur == '*') {
                vector<int> v1 = diffWaysWithMemo(memo, input.substr(0, static_cast<unsigned long>(i)));
                vector<int> v2 = diffWaysWithMemo(memo, input.substr(static_cast<unsigned long>(i + 1)));
                for (int n1 : v1) {
                    for (int n2: v2) {
                        if (cur == '+')
                            result.push_back(n1 + n2);
                        else if (cur == '-')
                            result.push_back(n1 - n2);
                        else
                            result.push_back(n1 * n2);
                    }
                }
            }
        }
        if (result.empty())
            result.push_back(atoi(input.c_str()));
        memo[input] = result;
        return result;
    }
};

#endif //LEETCODE_241_HPP
