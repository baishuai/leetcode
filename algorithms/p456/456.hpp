//
// Created by bai on 17-6-25.
//

#ifndef LEETCODE_456_HPP
#define LEETCODE_456_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <unordered_map>
#include <stack>

using namespace std;

struct Pair {
    int min;
    int max;
};

class Solution {
public:
    bool find132pattern(vector<int> &nums) {
        stack<Pair> sk;
        for (int n : nums) {
            if (sk.empty() || sk.top().min > n) {
                sk.emplace(Pair{n, n});
            } else if (n > sk.top().min) {
                // n > top.min
                auto l = sk.top();
                sk.pop();
                if (n < l.max) {
                    return true;
                } else {
                    l.max = n;
                    while (!sk.empty() && sk.top().max <= n) {
                        sk.pop();
                    }
                    // n < sk.top.max if sk is not empty
                    if (!sk.empty() && sk.top().min < n) {
                        return true;
                    }
                    sk.emplace(l);
                }

            }
        }
        return false;
    }
};


#endif //LEETCODE_456_HPP
