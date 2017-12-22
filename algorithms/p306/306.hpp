
#ifndef LEETCODE_306_HPP
#define LEETCODE_306_HPP

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
    bool isAdditiveNumber(string num) {
        if (num.empty() || num.length() < 3 )
            return false;
        auto n = static_cast<int >(num.length());

        for (int i = 1; i < n; ++i) {
            if (i > 1 && num[0] == '0')
                break;
            for (int j = i + 1; j < n; ++j) {
                int first = 0, second = i, third = j;
                if (num[second] == '0' && third > second + 1)
                    break;
                while (third < n) {
                    long result = stol(num.substr(first, second - first))
                                  + stol(num.substr(second, third - second));
                    if (to_string(result) == num.substr(third, to_string(result).length())) {
                        first = second;
                        second = third;
                        third += to_string(result).length();
                    } else {
                        break;
                    }
                }
                if (third == n)
                    return true;
            }
        }
        return false;
    }
};

#endif //LEETCODE_306_HPP
