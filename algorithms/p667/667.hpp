
#ifndef LEETCODE_667_HPP
#define LEETCODE_667_HPP

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
    vector<int> constructArray(int n, int k) {
        vector<int> arr(n);
        int l = 1, h = n;
        int flag = 1, cur = 0;
        while (--k > 0) {
            if (flag > 0) {
                arr[cur] = l++;
            } else {
                arr[cur] = h--;
            }
            flag *= -1;
            cur++;
        }
        int r = flag > 0 ? l : h;
        while (cur < n) {
            arr[cur++] = r;
            r += flag;
        }
        return arr;
    }
};

#endif //LEETCODE_667_HPP
