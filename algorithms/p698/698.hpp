
#ifndef LEETCODE_698_HPP
#define LEETCODE_698_HPP

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
    bool canPartitionKSubsets(vector<int> &nums, int k) {
        sort(nums.begin(), nums.end(), [](int a, int b) -> bool { return a > b; });
        int sum = accumulate(nums.begin(), nums.end(), 0);
        if (sum % k != 0)
            return false;
        int pk = sum / k;
        vector<int> buckets(k, 0);
        dfs(nums, pk, 0, buckets);
        return find;
    }

private:
    bool find = false;

    void dfs(vector<int> &nums, int limit, int pos, vector<int> &buckets) {
        if (find)
            return;
        if (pos >= nums.size()) {
            find = true;
            return;
        }
        bool flag = false;
        for (int i = 0; i < buckets.size(); ++i) {
            if (flag && buckets[i] == 0)
                continue;
            if (buckets[i] == 0)
                flag = true;
            if (buckets[i] + nums[pos] > limit)
                continue;
            buckets[i] += nums[pos];
            dfs(nums, limit, pos + 1, buckets);
            buckets[i] -= nums[pos];
        }
    }
};


#endif //LEETCODE_698_HPP
