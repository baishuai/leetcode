
#ifndef LEETCODE_331_HPP
#define LEETCODE_331_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <iterator>
#include <unordered_set>
#include <set>
#include <numeric>
#include <stack>
#include <sstream>
#include <string>

using namespace std;

class Solution {
public:
    bool isValidSerialization(string preorder) {
        stringstream ss(preorder);
        string item;
        vector<string> tokens;
        while (getline(ss, item, ',')) {
            tokens.push_back(item);
        }

        int degree = 1;
        for (auto &t : tokens) {
            if (--degree < 0) {
                return false;
            }
            if (t != "#") {
                degree += 2;
            }
        }
        return degree == 0;
    }
};


#endif //LEETCODE_331_HPP
