
#ifndef LEETCODE_693_HPP
#define LEETCODE_693_HPP

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

class Solution {
public:
    bool hasAlternatingBits(int n) {
        bool f = (n & 1) == 1;
        n >>= 1;
        while (n > 0) {
            if ((n & 1) == f)
                return false;
            f = !f;
            n >>= 1;
        }
        return true;
    }
};

#endif //LEETCODE_693_HPP
