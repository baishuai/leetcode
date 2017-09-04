
#ifndef LEETCODE_372_HPP
#define LEETCODE_372_HPP

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
 * Your task is to calculate ab mod 1337 where a is a positive integer and
 * b is an extremely large positive integer given in the form of an array.

Example1:

a = 2
b = [3]

Result: 8
Example2:

a = 2
b = [1,0]

Result: 1024
 */

class Solution {

public:
    int superPow(int a, vector<int> &b) {
        if (b.empty())
            return 1;
        int last = b.back();
        b.pop_back();
        return powmod(superPow(a, b), 10) * powmod(a, last) % base;
    }

private:
    const int base = 1337;

    int powmod(int a, int k) { // a^k % 1337  0<= k <= 10
        a %= base;
        int result = 1;
        for (int i = 0; i < k; ++i) {
            result = (result * a) % base;
        }
        return result;
    }
};

#endif //LEETCODE_372_HPP
