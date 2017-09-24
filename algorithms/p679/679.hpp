
#ifndef LEETCODE_679_HPP
#define LEETCODE_679_HPP

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
    bool judgePoint24(vector<int> &nums) {
        sort(begin(nums), end(nums));
        do {
            if (valid(nums))
                return true;
        } while (next_permutation(nums.begin(), nums.end()));
        return false;
    }

private:
    bool valid(vector<int> &nums) {
        double a = nums[0], b = nums[1], c = nums[2], d = nums[3];
        return ((valid(a + b, c, d) || valid(a - b, c, d) || valid(a * b, c, d) || valid(a / b, c, d)) ||
                (valid(a, b + c, d) || valid(a, b - c, d) || valid(a, b * c, d) || valid(a, b / c, d)) ||
                (valid(a, b, c + d) || valid(a, b, c - d) || valid(a, b, c * d) || valid(a, b, c / d)));
    }

    bool valid(double a, double b, double c) {
        return ((valid(a + b, c) || valid(a - b, c) || valid(a * b, c) || b && valid(a / b, c)) ||
                (valid(a, b + c) || valid(a, b - c) || valid(a, b * c) || c && valid(a, b / c)));

    }

    bool valid(double a, double b) {
        return abs(a + b - 24.0) < 0.0001 || abs(a - b - 24.0) < 0.0001 || abs(a * b - 24.0) < 0.0001 ||
               b && abs(a / b - 24.0) < 0.0001;
    }
};


#endif //LEETCODE_679_HPP
