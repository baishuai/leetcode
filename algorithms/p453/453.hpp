//
// Created by bai on 17-6-27.
//

#ifndef LEETCODE_453_HPP
#define LEETCODE_453_HPP

#include <iostream>
#include <queue>
#include <algorithm>
#include <vector>
#include <unordered_map>
#include <numeric>

using namespace std;


/**
let's define sum as the sum of all the numbers, before any moves;
 minNum as the min number int the list; n is the length of the list;

After, say m moves, we get all the numbers as x , and we will get the following equation

 sum + m * (n - 1) = x * n
and actually,

  x = minNum + m
and finally, we will get

  sum - minNum * n = m
So, it is clear and easy now.
 */


class Solution {
public:
    int minMoves(vector<int> &nums) {
        int sum = accumulate(nums.begin(), nums.end(), 0);
        int min = accumulate(nums.begin(), nums.end(), numeric_limits<int>::max(), [](int acc, int x) -> int {
            return acc < x ? acc : x;
        });
        return (int) (sum - min * nums.size());
    }
};

#endif //LEETCODE_453_HPP
