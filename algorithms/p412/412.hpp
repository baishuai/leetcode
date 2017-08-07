
#ifndef LEETCODE_412_HPP
#define LEETCODE_412_HPP

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
n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
 */

class Solution {
public:
    vector<string> fizzBuzz(int n) {
        vector<string> result(static_cast<unsigned long>(n));
        for (int i = 0; i < n; ++i) {
            if ((i + 1) % 15 == 0)
                result[i] = "FizzBuzz";
            else if ((i + 1) % 5 == 0)
                result[i] = "Buzz";
            else if ((i + 1) % 3 == 0)
                result[i] = "Fizz";
            else
                result[i] = to_string(i + 1);
        }
        return result;
    }
};

#endif //LEETCODE_412_HPP
