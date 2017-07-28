
#ifndef LEETCODE_345_HPP
#define LEETCODE_345_HPP

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
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:
Given s = "hello", return "holle".

Example 2:
Given s = "leetcode", return "leotcede".

Note:
The vowels does not include the letter "y".
 */


class Solution {
public:
    string reverseVowels(string s) {
        unordered_set<char> vowels = {'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'};
        int i = 0, j = static_cast<int>(s.size());
        while (i < j) {
            while (i < j && vowels.count(s[i]) == 0) {
                i++;
            }
            while (i < j && vowels.count(s[j]) == 0) {
                j--;
            }
            swap(s[i], s[j]);
            i++;
            j--;
        }
        return s;
    }
};

#endif //LEETCODE_345_HPP
