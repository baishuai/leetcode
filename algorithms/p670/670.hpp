
#ifndef LEETCODE_670_HPP
#define LEETCODE_670_HPP

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
    int maximumSwap(int num) {
        string numString = to_string(num);
        vector<int> maxPos(numString.size());
        int n = static_cast<int>(numString.size());
        int curMaxPos = n - 1;
        for (int i = n - 1; i >= 0; --i) {
            if (numString[i] > numString[curMaxPos])
                curMaxPos = i;
            maxPos[i] = curMaxPos;
        }

        for (int j = 0; j < n; ++j) {
            if (numString[maxPos[j]] != numString[j]) {
                swap(numString[j], numString[maxPos[j]]);
                break;
            }
        }
        return stoi(numString);

    }
};

#endif //LEETCODE_670_HPP
