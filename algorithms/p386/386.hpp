
#ifndef LEETCODE_386_HPP
#define LEETCODE_386_HPP

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
Given an integer n, return 1 - n in lexicographical order.

For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].

Please optimize your algorithm to use less time and space.
 The input size may be as large as 5,000,000.
 */

class Solution {
public:
    vector<int> lexicalOrder(int n) {
        int cur = 1, cnt = 0;
        vector<int> lst;
        while (++cnt <= n) {
            lst.emplace_back(cur);
            if (cur * 10 <= n)
                cur *= 10;
            else if (cur + 1 <= n && (cur % 10 != 9))
                cur += 1;
            else {
                while ((cur / 10) % 10 == 9)
                    cur /= 10;
                cur = cur / 10 + 1;
            }
        }
        return lst;
    }
};

#endif //LEETCODE_386_HPP
