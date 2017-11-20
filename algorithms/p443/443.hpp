
#ifndef LEETCODE_443_HPP
#define LEETCODE_443_HPP

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
    int compress(vector<char> &chars) {
        int cnt = 1, cur = 0;
        char pre = 0;
        for (int i = 0; i < chars.size(); ++i) {
            if (chars[i] == pre) {
                cnt++;
            } else {
                cur = writeOneChar(chars, pre, cnt, cur);
                pre = chars[i];
                cnt = 1;
            }
        }
        cur = writeOneChar(chars, pre, cnt, cur);
        return cur;

    }

    int writeOneChar(vector<char> &chars, char c, int cnt, int cur) {
        if (c == 0)
            return 0;
        chars[cur++] = c;
        if (cnt > 1) {
            string count = to_string(cnt);
            for (char c: count) {
                chars[cur++] = c;
            }
        }
        return cur;
    }
};

#endif //LEETCODE_443_HPP
