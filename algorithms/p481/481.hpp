
#ifndef LEETCODE_481_HPP
#define LEETCODE_481_HPP

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

/*
A magical string S consists of only '1' and '2' and obeys the following rules:

The string S is magical because concatenating the number of
 contiguous occurrences of characters '1' and '2' generates the string S itself.

The first few elements of string S is the following: S = "1221121221221121122……"

If we group the consecutive '1's and '2's in S, it will be:

1 22 11 2 1 22 1 22 11 2 11 22 ......

and the occurrences of '1's or '2's in each group are:

1 2	2 1 1 2 1 2 2 1 2 2 ......

You can see that the occurrence sequence above is the S itself.

Given an integer N as input, return the number of '1's in the first N number in the magical string S.
 */

class Solution {
public:
    int magicalString(int n) {
        string s = "122";
        int i = 2;
        while (s.size() < n) {
            if (s[i] == '1')
                s.push_back('1' + '2' - s.back());
            else {
                char c = '1' + '2' - s.back();
                s.push_back(c);
                s.push_back(c);
            }
            i++;
        }
        int cnt = 0;
        for (int j = 0; j < n; ++j) {
            if (s[j] == '1')
                cnt++;
        }
        return cnt;
    }
};

#endif //LEETCODE_481_HPP
