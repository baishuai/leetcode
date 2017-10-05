
#ifndef LEETCODE_575_HPP
#define LEETCODE_575_HPP

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
    int distributeCandies(vector<int> &candies) {
        unordered_set<int> kinds;
        for (int c : candies)
            kinds.insert(c);
        auto cnt = kinds.size() > candies.size() / 2 ? candies.size() / 2 : kinds.size();
        return static_cast<int>(cnt);
    }
};

#endif //LEETCODE_575_HPP
