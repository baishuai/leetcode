//
// Created by bai on 17-7-2.
//

#ifndef LEETCODE_634_HPP
#define LEETCODE_634_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <numeric>

using namespace std;


class Solution {
public:
    int findDerangement(int n) {
        long long int a = 0, b = 1;
        const long long int mod = 1000000000 + 7;
        for (int i = 2; i <= n; ++i) {
            auto c = i * (a + b) % mod;
            a = b;
            b = c;
        }
        return (int) a;
    }
};

#endif //LEETCODE_634_HPP
