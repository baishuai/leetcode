
#ifndef LEETCODE_492_HPP
#define LEETCODE_492_HPP

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
#include <cmath>

using namespace std;

class Solution {
public:
    vector<int> constructRectangle(int area) {
        int w = static_cast<int>(sqrt(area));
        while (area % w != 0) {
            w--;
        }
        return {area / w, w};
    }
};

#endif //LEETCODE_492_HPP
