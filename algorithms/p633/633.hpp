//
// Created by bai on 17-7-2.
//

#ifndef LEETCODE_633_HPP
#define LEETCODE_633_HPP

class Solution {

public:
    bool judgeSquareSum(int c) {
        if (isPerfectSquare(c)) {
            return true;
        }
        for (int i = 0; i * i <= c / 2; ++i) {
            if (isPerfectSquare(c - i * i)) {
                return true;
            }
        }
        return false;
    }

private:
    bool isPerfectSquare(int num) {
        if (num <= 1) {
            return true;
        }
        long r = num;
        while (r * r > num) {
            r = (r + num / r) / 2;
        }
        return r * r == num;
    }
};

#endif //LEETCODE_633_HPP
