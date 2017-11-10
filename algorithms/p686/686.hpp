
#ifndef LEETCODE_686_HPP
#define LEETCODE_686_HPP

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
    int repeatedStringMatch(string A, string B) {
        int mtimes = -1;
        for (int i = 0; i < A.length(); ++i) {
            int j = 0;
            while (j < B.size() ) {
                if (A[(i + j) % A.length()] != B[j]){
                    if (i + j >= A.length())
                        return -1;
                    break;
                }
                j++;
            }
            if (j == B.size())
                return static_cast<int>((i + j + A.length() - 1) / A.length());
        }
        return -1;
    }
};

#endif //LEETCODE_686_HPP
