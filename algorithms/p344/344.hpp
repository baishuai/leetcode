//
// Created by bai on 17-6-26.
//

#ifndef LEETCODE_344_HPP
#define LEETCODE_344_HPP

#include <iostream>
#include <algorithm>
#include <sstream>

using namespace std;


class Solution {
public:
    string reverseString(string s) {
        return string(s.rbegin(), s.rend());
    }
};


#endif //LEETCODE_344_HPP
