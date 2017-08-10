
#ifndef LEETCODE_537_HPP
#define LEETCODE_537_HPP

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
#include <sstream>

using namespace std;

/*
Given two strings representing two complex numbers.

You need to return a string representing their multiplication. Note i2 = -1 according to the definition.

Example 1:
Input: "1+1i", "1+1i"
Output: "0+2i"
Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
Example 2:
Input: "1+-1i", "1+-1i"
Output: "0+-2i"
Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
Note:

The input strings will not have extra blank.
The input strings will be given in the form of a+bi,
 where the integer a and b will both belong to the range of [-100, 100].
 And the output should be also in this form.
 */

class Solution {
public:
    string complexNumberMultiply(string a, string b) {
        int ra, ia, rb, ib;
        char buf;
        stringstream aa(a), bb(b), ans;
        aa >> ra >> buf >> ia >> buf;
        bb >> rb >> buf >> ib >> buf;
        ans << ra * rb - ia * ib << "+" << ra * ib + rb * ia << "i";
        return ans.str();
    }
};


#endif //LEETCODE_537_HPP
