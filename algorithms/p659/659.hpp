
#ifndef LEETCODE_659_HPP
#define LEETCODE_659_HPP

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
    bool isPossible(vector<int> &nums) {
        unordered_map<int, priority_queue<int, vector<int>, greater<int >>> backs;

        int need_more = 0;
        for (int num : nums) {
            if (backs[num - 1].empty()) {
                backs[num].push(1);
                need_more++;
            } else {
                int count = backs[num - 1].top();
                backs[num - 1].pop();
                backs[num].push(++count);
                if (count == 3)
                    need_more--;
            }
        }
        return need_more == 0;
    }
};


#endif //LEETCODE_659_HPP
