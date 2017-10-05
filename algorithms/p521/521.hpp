
#ifndef LEETCODE_521_HPP
#define LEETCODE_521_HPP

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
    int findLUSlength(string a, string b) {
        return a == b ? -1 : max(a.length(), b.length());
    }
};

#endif //LEETCODE_521_HPP
