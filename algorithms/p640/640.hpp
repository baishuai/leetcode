
#ifndef LEETCODE_640_HPP
#define LEETCODE_640_HPP

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
    string solveEquation(string equation) {
        int coefs = 0, prod = 0, sign = 1, i = 0, flag = 1;
        int len = equation.length();
        for (int j = 0; j < len; ++j) {
            if (equation[j] == '+' || equation[j] == '-') {
                if (j > i)
                    prod += sign * -1 * flag * stoi(equation.substr(i, j - i));
                i = j + 1;
                flag = equation[j] == '+' ? 1 : -1;
            } else if (equation[j] == 'x') {
                if (j > i) {
                    coefs += sign * flag * stoi(equation.substr(i, j - i));
                } else {
                    coefs += sign * flag * 1;
                }
                i = j + 1;
            } else if (equation[j] == '=') {
                if (j > i) {
                    prod += sign * -1 * flag * stoi(equation.substr(i, j - i));
                }
                sign = -1;
                flag = 1;
                i = j + 1;
            }
        }
        if (i < len)
            prod += sign * -1 * flag * stoi(equation.substr(i));

        if (coefs == 0 && prod == 0) {
            return "Infinite solutions";
        } else if (coefs == 0 && prod != 0) {
            return "No solution";
        } else {
            return "x=" + to_string(prod / coefs);
        }
    }
};

#endif //LEETCODE_640_HPP
