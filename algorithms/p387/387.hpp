
#ifndef LEETCODE_387_HPP
#define LEETCODE_387_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <unordered_set>
#include <set>
#include <map>
#include <numeric>
#include <stack>
#include <string>

using namespace std;

/*

Given a string, find the first non-repeating character in it and return it's index.
 If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.

 */

class Solution {
public:
    int firstUniqChar(string s) {
        map<char, int> cnts;
        for (int i = 0; i < s.size(); ++i) {
            cnts[s[i]] += 1;
        }
        int idx = -1;
        for (int i = 0; i < s.size(); ++i) {
            if (cnts[s[i]] == 1) {
                idx = i;
                break;
            }
        }
        return idx;
    }
};


#endif //LEETCODE_387_HPP
