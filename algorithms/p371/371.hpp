//
// Created by bai on 17-6-28.
//

#ifndef LEETCODE_371_HPP
#define LEETCODE_371_HPP

#include <cstdlib>
#include <iostream>

class Solution {
public:
    int getSum(int a, int b) {
        return b == 0 ? a : getSum(a ^ b, (a & b) << 1);
    }
};

#endif //LEETCODE_371_HPP
