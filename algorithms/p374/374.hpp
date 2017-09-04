
#ifndef LEETCODE_374_HPP
#define LEETCODE_374_HPP

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

// Forward declaration of guess API.
// @param num, your guess
// @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
int guess(int num);

class Solution {
public:
    int guessNumber(int n) {
        int l = 1, h = n + 1;
        int mean = 0;
        while (true) {
            mean = l + (h - l) / 2;
            int g = guess(mean);
            if (g == 0)
                break;
            else if (g > 0) {
                l = mean + 1;
            } else {
                h = mean;
            }
        }
        return mean;
    }
};


#endif //LEETCODE_374_HPP
