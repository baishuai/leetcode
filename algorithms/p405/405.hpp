
#ifndef LEETCODE_405_HPP
#define LEETCODE_405_HPP

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
Given an integer, write an algorithm to convert it to hexadecimal.
 For negative integer, two’s complement method is used.

Note:

All letters in hexadecimal (a-f) must be in lowercase.
The hexadecimal string must not contain extra leading 0s.
 If the number is zero, it is represented by a single zero character '0';
 otherwise, the first character in the hexadecimal string will not be the zero character.
The given number is guaranteed to fit within the range of a 32-bit signed integer.
You must not use any method provided by the library which converts/formats the number to hex directly.

 */

class Solution {
public:
    string toHex(int num) {
        string hex;
        unsigned n = static_cast<unsigned int>(num);
        int mask = 0xF;
        while (n != 0) {
            int h = n & mask;
            n = n >> 4;
            hex.push_back(static_cast<char>(h > 9 ? 'a' + h - 10 : '0' + h));
        }
        if (hex.size() == 0)
            hex.push_back('0');
        return string(hex.rbegin(), hex.rend());
    }
};


#endif //LEETCODE_405_HPP
