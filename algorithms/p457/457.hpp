
#ifndef LEETCODE_457_HPP
#define LEETCODE_457_HPP

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
    bool circularArrayLoop(vector<int> &nums) {
        bool loop{false};
        for (int i = 0; i < nums.size(); ++i) {
            int n = nums[i];
            if (n == 0)
                continue;
            int next = i, j;
            while (true) {
                j = next;
                if (nums[j] < 0)
                    break;
                next = static_cast<int>((j + nums.size() + nums[j]) % nums.size());
                nums[j] = 0;
                if (next == j)
                    break;
                else if (nums[next] == 0) {
                    loop = true;
                    break;
                }
            }
        }
        return loop;
    }
};

#endif //LEETCODE_457_HPP
