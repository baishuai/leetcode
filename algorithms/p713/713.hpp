
#ifndef LEETCODE_713_HPP
#define LEETCODE_713_HPP

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
    int numSubarrayProductLessThanK(vector<int> &nums, int k) {
        if (k <= 1)
            return 0;
        int i = 0, j = 0, pProd = 1, cnt = 0;
        while (j < nums.size()) {
            pProd *= nums[j];
            while (pProd >= k && i <= j) {
                pProd /= nums[i++];
            }
            cnt += j - i + 1;
            j++;
        }
        return cnt;
    }
};


#endif //LEETCODE_713_HPP
