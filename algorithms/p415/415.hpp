
#ifndef LEETCODE_415_HPP
#define LEETCODE_415_HPP

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
    string addStrings(string num1, string num2) {
        int i = static_cast<int>(num1.size() - 1), j = static_cast<int>(num2.size() - 1);
        string res;
        int carry = 0;
        while (i >= 0 || j >= 0 || carry > 0) {
            carry += i >= 0 ? (num1[i] - '0') : 0;
            carry += j >= 0 ? (num2[j] - '0') : 0;
            res.push_back(static_cast<char>('0' + (carry % 10)));
            carry = carry / 10;
            i--;
            j--;
        }
        return string(res.rbegin(), res.rend());
    }
};


#endif //LEETCODE_415_HPP
