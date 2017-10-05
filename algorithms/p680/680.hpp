
#ifndef LEETCODE_680_HPP
#define LEETCODE_680_HPP

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
    bool validPalindrome(string s) {
        int l = 0, r = static_cast<int>(s.length() - 1);
        while (l < r) {
            if (s[l] != s[r])
                return isPalindromic(s, l + 1, r) || isPalindromic(s, l, r - 1);
            l++;
            r--;
        }
        return true;
    }

private:
    bool isPalindromic(string &s, int l, int r) {
        while (l < r) {
            if (s[l] != s[r])
                return false;
            l++;
            r--;
        }
        return true;
    }
};

#endif //LEETCODE_680_HPP
