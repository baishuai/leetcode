
#ifndef LEETCODE_639_HPP
#define LEETCODE_639_HPP

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
A message containing letters from A-Z is being encoded to numbers
 using the following mapping way:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Beyond that, now the encoded string can also contain the character '*',
 which can be treated as one of the numbers from 1 to 9.

Given the encoded message containing digits and the character '*',
 return the total number of ways to decode it.

Also, since the answer may be very large, you should return the output mod 109 + 7.
 */

class Solution {
public:
    int numDecodings(string s) {
        int n = s.size(), p = 1000000007;
        long f1 = 1, f2 = helper(s.substr(0, 1));
        for (unsigned long i = 1; i < n; ++i) {
            long f3 = f2 * helper(s.substr(i, 1)) + f1 * helper(s.substr(i - 1, 2));
            f1 = f2;
            f2 = f3 % p;
        }
        return f2;
    }

private:
    int helper(string s) {
        if (s.size() == 1) {
            return s[0] == '*' ? 9 : s[0] == '0' ? 0 : 1;
        }
        if (s == "**") {
            return 15;
        } else if (s[1] == '*') {
            return s[0] == '1' ? 9 : s[0] == '2' ? 6 : 0;
        } else if (s[0] == '*') {
            return s[1] <= '6' ? 2 : 1;
        } else {
            return stoi(s) >= 10 && stoi(s) <= 26 ? 1 : 0;
        }
    }
};

#endif //LEETCODE_639_HPP
