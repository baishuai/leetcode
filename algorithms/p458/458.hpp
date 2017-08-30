
#ifndef LEETCODE_458_HPP
#define LEETCODE_458_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <numeric>
#include <stack>
#include <cmath>
#include <string>

using namespace std;

class Solution {
public:
    int poorPigs(int buckets, int minutesToDie, int minutesToTest) {
        int pigs = 0;
        while (pow(minutesToTest / minutesToDie + 1, pigs) < buckets)
            pigs++;
        return pigs;
    }
};


#endif //LEETCODE_458_HPP
