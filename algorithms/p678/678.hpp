
#ifndef LEETCODE_678_HPP
#define LEETCODE_678_HPP

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
    bool checkValidString(string s) {
        int lower = 0, upper = 0;
        for (char c : s) {
            lower += c == '(' ? 1 : -1;
            upper += c != ')' ? 1 : -1;
            if (upper < 0)
                break;
            lower = max(lower, 0);
        }
        return lower == 0;
    }
};


#endif //LEETCODE_678_HPP
